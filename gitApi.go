/*
Copyright © 2022 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/J-Siu/go-helper"
)

// GitApi http input structure
type GitApiReq struct {
	Data       string       `json:"Data"`       // Json marshaled Info
	Entrypoint string       `json:"Entrypoint"` // Api base url
	Endpoint   string       `json:"Endpoint"`   // Api endpoint
	Header     *http.Header `json:"Header"`     // Http request header
	Method     string       `json:"Method"`     // Http request method
	Token      string       `json:"Token"`      // Api auth token
	UrlVal     *url.Values  `json:"UrlVal"`     // Api url values
}

// GitApi http output structure
type GitApiRes struct {
	Body   *[]byte      `json:"Body"`
	Err    string       `json:"Err"`
	Header *http.Header `json:"Header"` // Http response header
	Url    *url.URL     `json:"Url"`    // In.Uri + In.Endpoint
	Output *string      `json:"Output"` // Api response body in string
	Status string       `json:"Status"` // Http response status
}

// GitApi
type GitApi struct {
	Req    GitApiReq  `json:"In"`     // Api http input
	Res    GitApiRes  `json:"Out"`    // Api http output
	Name   string     `json:"Name"`   // Name of connection
	User   string     `json:"User"`   // Api username
	Vendor string     `json:"Vendor"` // github/gitea
	Repo   string     `json:"Repo"`   // Repository name
	Info   GitApiInfo `json:"Info"`   // Pointer to structure. Use NilType.Nil() for nil pointer
}

// Setup a *GitApi
func GitApiNew(
	name string,
	token string,
	entrypoint string,
	user string,
	vendor string,
	repo string,
	info GitApiInfo) *GitApi {
	var self GitApi
	self.Name = name
	self.User = user
	self.Vendor = vendor
	self.Repo = repo
	self.Info = info
	self.Req.Entrypoint = entrypoint
	self.Req.Token = token
	return &self
}

// Initialize endpoint /user/repos
func (self *GitApi) EndpointUserRepos() *GitApi {
	self.Req.Endpoint = "/user/repos"
	return self
}

// Initialize endpoint /repos/OWNER/REPO
//
// Use current directory if GitApi.Repo is empty
func (self *GitApi) EndpointRepos() *GitApi {
	self.Req.Endpoint = path.Join("repos", self.User, self.Repo)
	return self
}

// Initialize endpoint /repos/OWNER/REPO/topics
func (self *GitApi) EndpointReposTopics() *GitApi {
	self.EndpointRepos()
	self.Req.Endpoint = path.Join(self.Req.Endpoint, "topics")
	return self
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets
func (self *GitApi) EndpointReposSecrets() *GitApi {
	self.EndpointRepos()
	self.Req.Endpoint = path.Join(self.Req.Endpoint, "actions", "secrets")
	return self
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets/public-key
func (self *GitApi) EndpointReposSecretsPubkey() *GitApi {
	self.EndpointReposSecrets()
	self.Req.Endpoint = path.Join(self.Req.Endpoint, "public-key")
	return self
}

// Set github/gitea header
//
// GitApi.Req.Token, if empty, authorizeation header will not be set.
func (self *GitApi) HeaderGithub() *GitApi {
	header := make(http.Header)
	self.Req.Header = &header
	self.Req.Header.Add("Accept", "application/vnd.github.v3+json")
	self.Req.Header.Add("Content-Type", "application/json")
	if self.Req.Token != "" {
		self.Req.Header.Add("Authorization", "token "+self.Req.Token)
	}
	return self
}

// Setup empty API header
func (self *GitApi) HeaderInit() *GitApi {
	header := make(http.Header)
	self.Req.Header = &header
	return self
}

// Execute http request using info in GitApi.Req. Then put response info in GitApi.Res.
//
//	GitApi.Info, if not nil, will be
//			- auto marshaled for send other than "GET"
//			- auto unmarshaled from http response body
func (self *GitApi) Do() *GitApi {
	// Prepare Api Data
	if self.Req.Method != http.MethodGet {
		j, _ := json.Marshal(&self.Info)
		self.Req.Data = string(j)
	}
	// Prepare url
	self.Res.Url, _ = url.Parse(self.Req.Entrypoint)
	self.Res.Url.Path = path.Join(self.Res.Url.Path, self.Req.Endpoint)
	if self.Req.UrlVal != nil {
		self.Res.Url.RawQuery = self.Req.UrlVal.Encode()
	}
	// Prepare request
	dataBufferP := bytes.NewBufferString(self.Req.Data)
	req, err := http.NewRequest(
		self.Req.Method,
		self.Res.Url.String(),
		dataBufferP)
	if err != nil {
		self.Res.Err = err.Error()
	}
	// Set request headers
	req.Header = *self.Req.Header
	// Request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		self.Res.Err = err.Error()
	}
	// Response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		self.Res.Err = err.Error()
	}
	res.Body.Close()
	// Fill in self.Out
	self.Res.Body = &body
	self.Res.Header = &res.Header
	self.Res.Status = res.Status

	// Unmarshal
	self.ProcessOutput()

	helper.ReportDebug(&self, "api", false, false)

	return self
}

// GitApi Get action wrapper
func (self *GitApi) Get() *GitApi {
	self.Req.Method = http.MethodGet
	return self.Do()
}

// GitApi Del action wrapper
func (self *GitApi) Del() *GitApi {
	self.Req.Method = http.MethodDelete
	return self.Do()
}

// GitApi Patch action wrapper
func (self *GitApi) Patch() *GitApi {
	self.Req.Method = http.MethodPatch
	return self.Do()
}

// GitApi Post action wrapper
func (self *GitApi) Post() *GitApi {
	self.Req.Method = http.MethodPost
	return self.Do()
}

// GitApi Put action wrapper
func (self *GitApi) Put() *GitApi {
	self.Req.Method = http.MethodPut
	return self.Do()
}

// Print both Body and Err into string pointer
func (self *GitApi) ProcessOutput() *GitApi {
	// Unmarshal
	err := json.Unmarshal(*self.Res.Body, self.Info)
	if self.Res.Ok() && err == nil && self.Info != nil {
		// Use Info string func
		self.Res.Output = self.Info.StringP()
	} else {
		var output string
		strP := helper.ReportSp(self.Res.Body, "", true, false)
		if strP != nil {
			output += *strP
		}
		strP = helper.ReportSp(self.Res.Err, "", true, false)
		if strP != nil {
			output += *strP
		}
		self.Res.Output = &output
	}
	return self
}

// Setup empty API url values
func (self *GitApiReq) UrlValInit() {
	urlVal := make(url.Values)
	self.UrlVal = &urlVal
}

// Check response status == 2xx
func (self *GitApiRes) Ok() bool {
	return self.Status != "" && self.Status[0] == '2'
}
