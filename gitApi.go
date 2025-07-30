/*
The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package gitapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/J-Siu/go-helper"
)

// GitApi http input structure
type GitApiReq struct {
	Entrypoint string       `json:"Entrypoint"` // Api base url
	Endpoint   string       `json:"Endpoint"`   // Api endpoint
	Token      string       `json:"Token"`      // Api auth token
	Header     *http.Header `json:"Header"`     // Http request header
	UrlVal     *url.Values  `json:"UrlVal"`     // Api url values
	Data       string       `json:"Data"`       // Json marshaled Info
}

// Setup empty API url values
func (gaReq *GitApiReq) UrlValInit() {
	urlVal := make(url.Values)
	gaReq.UrlVal = &urlVal
}

// GitApi http output structure
type GitApiRes struct {
	Url    *url.URL     `json:"Url"`    // In.Uri + In.Endpoint
	Header *http.Header `json:"Header"` // Http response header
	Status string       `json:"Status"` // Http response status
	Body   *[]byte      `json:"Body"`
	Output *string      `json:"Output"` // Api response body in string
	Err    string       `json:"Err"`
}

// Check response status == 2xx
func (gaRes *GitApiRes) Ok() bool {
	return gaRes.Status != "" && gaRes.Status[0] == '2' && gaRes.Err == ""
}

// GitApi
type GitApi struct {
	Method     string     `json:"Method"`     // Http request method
	SkipVerify bool       `json:"skipverify"` // Api request skip cert verify (allow self-signed cert)
	Vendor     string     `json:"Vendor"`     // github/gitea
	Name       string     `json:"Name"`       // Name of connection
	Repo       string     `json:"Repo"`       // Repository name
	User       string     `json:"User"`       // Api username
	Req        GitApiReq  `json:"In"`         // Api http input
	Res        GitApiRes  `json:"Out"`        // Api http output
	Info       GitApiInfo `json:"Info"`       // Pointer to structure. Use NilType.Nil() for nil pointer
}

// Setup a *GitApi
func GitApiNew(
	name string,
	token string,
	entrypoint string,
	user string,
	vendor string,
	skipverify bool,
	repo string,
	info GitApiInfo) *GitApi {
	var self GitApi
	self.Name = name
	self.User = user
	self.Vendor = vendor
	self.SkipVerify = skipverify
	self.Repo = repo
	self.Info = info
	self.Req.Entrypoint = entrypoint
	self.Req.Token = token
	return &self
}

// Return Res.Body
func (ga *GitApi) Body() *string {
	return helper.ReportSp(ga.Res.Body, "", true, true)
}

// Return Res.Err
func (ga *GitApi) Err() *string {
	return &ga.Res.Err
}

// Return Res.Ok()
func (ga *GitApi) Ok() bool {
	return ga.Res.Ok()
}

// Return Res.Output
func (ga *GitApi) Output() *string {
	return ga.Res.Output
}

// Initialize endpoint /user/repos
func (ga *GitApi) EndpointUserRepos() *GitApi {
	ga.Req.Endpoint = "/user/repos"
	return ga
}

// Initialize endpoint /repos/OWNER/REPO
//
// Use current directory if GitApi.Repo is empty
func (ga *GitApi) EndpointRepos() *GitApi {
	ga.Req.Endpoint = path.Join("repos", ga.User, ga.Repo)
	return ga
}

// Initialize endpoint /repos/OWNER/REPO/topics
func (ga *GitApi) EndpointReposTopics() *GitApi {
	ga.EndpointRepos()
	ga.Req.Endpoint = path.Join(ga.Req.Endpoint, "topics")
	return ga
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets
func (ga *GitApi) EndpointReposSecrets() *GitApi {
	ga.EndpointRepos()
	ga.Req.Endpoint = path.Join(ga.Req.Endpoint, "actions", "secrets")
	return ga
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets/public-key
func (ga *GitApi) EndpointReposSecretsPubkey() *GitApi {
	ga.EndpointReposSecrets()
	ga.Req.Endpoint = path.Join(ga.Req.Endpoint, "public-key")
	return ga
}

// Set github/gitea header
//
// GitApi.Req.Token, if empty, authorization header will not be set.
func (ga *GitApi) HeaderGithub() *GitApi {
	header := make(http.Header)
	ga.Req.Header = &header
	ga.Req.Header.Add("Accept", "application/vnd.github.v3+json")
	ga.Req.Header.Add("Content-Type", "application/json")
	if ga.Req.Token != "" {
		ga.Req.Header.Add("Authorization", "token "+ga.Req.Token)
	}
	return ga
}

// Setup empty API header
func (ga *GitApi) HeaderInit() *GitApi {
	header := make(http.Header)
	ga.Req.Header = &header
	return ga
}

// Execute http request using info in GitApi.Req. Then put response info in GitApi.Res.
//
//	GitApi.Info, if not nil, will be
//			- auto marshal for send other than "GET"
//			- auto unmarshal from http response body
func (ga *GitApi) Do() *GitApi {
	// Prepare Api Data
	if ga.Method != http.MethodGet && ga.Info != nil {
		j, _ := json.Marshal(&ga.Info)
		ga.Req.Data = string(j)
	}
	// Prepare url
	ga.Res.Url, _ = url.Parse(ga.Req.Entrypoint)
	ga.Res.Url.Path = path.Join(ga.Res.Url.Path, ga.Req.Endpoint)
	if ga.Req.UrlVal != nil {
		ga.Res.Url.RawQuery = ga.Req.UrlVal.Encode()
	}
	// Prepare request
	dataBufferP := bytes.NewBufferString(ga.Req.Data)
	// ga.Req.Method = ga.Method
	req, err := http.NewRequest(
		ga.Method,
		ga.Res.Url.String(),
		dataBufferP)
	if err != nil {
		ga.Res.Err = err.Error()
	}
	// Set request headers
	req.Header = *ga.Req.Header
	// Request
	// - Configure transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: ga.SkipVerify},
	}

	client := &http.Client{Transport: transport}
	res, err := client.Do(req)
	if err != nil {
		ga.Res.Err = err.Error()
	} else {
		// Response
		body, err := io.ReadAll(res.Body)
		if err != nil {
			ga.Res.Err = err.Error()
		}
		res.Body.Close()
		// Fill in self.Out
		ga.Res.Body = &body
		ga.Res.Header = &res.Header
		ga.Res.Status = res.Status
	}

	// Unmarshal
	if ga.Res.Err == "" {
		ga.ProcessOutput()
	} else {
		ga.ProcessError()
	}

	helper.ReportDebug(&ga, "api", false, false)
	helper.ReportDebug(ga.Res.Body, "api.Out.Body (decoded)", false, false)

	return ga
}

// GitApi Get action wrapper
func (ga *GitApi) Get() *GitApi {
	ga.SetGet()
	return ga.Do()
}

// GitApi Del action wrapper
func (ga *GitApi) Del() *GitApi {
	ga.SetDel()
	return ga.Do()
}

// GitApi Patch action wrapper
func (ga *GitApi) Patch() *GitApi {
	ga.SetPatch()
	return ga.Do()
}

// GitApi Post action wrapper
func (ga *GitApi) Post() *GitApi {
	ga.SetPost()
	return ga.Do()
}

// GitApi Put action wrapper
func (ga *GitApi) Put() *GitApi {
	ga.Method = http.MethodPut
	return ga.Do()
}

func (ga *GitApi) SetGet() *GitApi {
	ga.Method = http.MethodGet
	return ga
}

// GitApi set http Del
func (ga *GitApi) SetDel() *GitApi {
	ga.Method = http.MethodDelete
	return ga
}

// GitApi set http Patch
func (ga *GitApi) SetPatch() *GitApi {
	ga.Method = http.MethodPatch
	return ga
}

// GitApi set http Post
func (ga *GitApi) SetPost() *GitApi {
	ga.Method = http.MethodPost
	return ga
}

// GitApi set http Put
func (ga *GitApi) SetPut() *GitApi {
	ga.Method = http.MethodPut
	return ga
}

// Print HTTP Body into string pointer
func (ga *GitApi) ProcessOutput() *GitApi {
	// Check API error
	ga.ProcessOutputError()
	if ga.Res.Ok() {
		// Unmarshal
		if ga.Info == Nil() {
			// Return the whole JSON
			tmpStr := string(*ga.Res.Body)
			ga.Res.Output = &tmpStr
		} else {
			err := json.Unmarshal(*ga.Res.Body, ga.Info)
			if ga.Res.Ok() && err == nil && ga.Info != nil {
				// Use Info string func
				ga.Res.Output = ga.Info.StringP()
			}
		}
	}
	return ga
}

// Print HTTP Body into string pointer
func (ga *GitApi) ProcessOutputError() *GitApi {
	var e RepoError
	err := json.Unmarshal(*ga.Res.Body, &e)
	if err == nil {
		// Use Info string func
		if e.Message != "" {
			ga.Res.Err = e.String()
		}
	}
	return ga
}

// Print HTTP Body Err into string pointer
func (ga *GitApi) ProcessError() *GitApi {
	helper.ReportDebug(*ga.Res.Body, "api.Out.Body", false, false)
	// Unmarshal
	var output string
	strP := helper.ReportSp(ga.Res.Body, "", true, false)
	if strP != nil {
		output += *strP
	}
	strP = helper.ReportSp(ga.Res.Err, "", true, false)
	if strP != nil {
		output += *strP
	}
	ga.Res.Output = &output
	return ga
}
