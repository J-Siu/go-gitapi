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
type GitApiIn struct {
	Data       string       `json:"Data"`       // Json marshaled Info
	Entrypoint string       `json:"Entrypoint"` // Api base url
	Endpoint   string       `json:"Endpoint"`   // Api endpoint
	Header     *http.Header `json:"Header"`     // Http request header
	Method     string       `json:"Method"`     // Http request method
	Token      string       `json:"Token"`      // Api auth token
}

// GitApi http output structure
type GitApiOut struct {
	Body    *[]byte      `json:"Body"`
	Err     string       `json:"Err"`
	Header  *http.Header `json:"Header"`  // Http response header
	Url     string       `json:"Url"`     // In.Uri + In.Endpoint
	Output  *string      `json:"Output"`  // Api response body in string
	Status  string       `json:"Status"`  // Http response status
	Success bool         `json:"Success"` // Set to true if status code 2xx
}

// GitApi
type GitApi[T GitApiInfo] struct {
	In     GitApiIn  `json:"In"`     // Api http input
	Out    GitApiOut `json:"Out"`    // Api http output
	Name   string    `json:"Name"`   // Name of connection
	User   string    `json:"User"`   // Api username
	Vendor string    `json:"Vendor"` // github/gitea
	Repo   string    `json:"Repo"`   // Repository name
	Info   T         `json:"Info"`   // Pointer to structure. Use NilType.Nil() for nil pointer
}

// Setup a *GitApi
func GitApiNew[T GitApiInfo](
	name string,
	token string,
	entrypoint string,
	user string,
	vendor string,
	repo string,
	info T) *GitApi[T] {
	var self GitApi[T]
	self.Name = name
	self.User = user
	self.Vendor = vendor
	self.Repo = repo
	self.Info = info
	self.In.Entrypoint = entrypoint
	self.In.Token = token
	return &self
}

// Initialize endpoint /user/repos
func (self *GitApi[GitApiInfo]) EndpointUserRepos() {
	self.In.Endpoint = "/user/repos"
}

// Initialize endpoint /repos/OWNER/REPO
//
// Use current directory if GitApi.Repo is empty
func (self *GitApi[GitApiInfo]) EndpointRepos() {
	self.In.Endpoint = path.Join("repos", self.User, self.Repo)
}

// Initialize endpoint /repos/OWNER/REPO/topics
func (self *GitApi[GitApiInfo]) EndpointReposTopics() {
	self.EndpointRepos()
	self.In.Endpoint = path.Join(self.In.Endpoint, "topics")
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets
func (self *GitApi[GitApiInfo]) EndpointReposSecrets() {
	self.EndpointRepos()
	self.In.Endpoint = path.Join(self.In.Endpoint, "actions", "secrets")
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets/public-key
func (self *GitApi[GitApiInfo]) EndpointReposSecretsPubkey() {
	self.EndpointReposSecrets()
	self.In.Endpoint = path.Join(self.In.Endpoint, "public-key")
}

// Set github/gitea header
func (self *GitApi[GitApiInfo]) HeaderGithub() {
	header := make(http.Header)
	self.In.Header = &header
	self.In.Header.Add("Accept", "application/vnd.github.v3+json")
	self.In.Header.Add("Content-Type", "application/json")
	if len(self.In.Token) > 0 {
		self.In.Header.Add("Authorization", "token "+self.In.Token)
	}
}

//	Execute http request using info in GitApi.In. Then put response info in GitApi.Out.
//
//	GitApi.In.Token, if empty, authorizeation header will not be set.
//
//	GitApi.Info, if not nil, will be
//			- auto marshaled for send other than "GET"
//			- auto unmarshaled from http response body
func (self *GitApi[GitApiInfo]) Do() bool {
	// Prepare Api Data
	if self.In.Method != http.MethodGet {
		j, _ := json.Marshal(&self.Info)
		self.In.Data = string(j)
	}
	// Prepare url
	url, _ := url.Parse(self.In.Entrypoint)
	url.Path = path.Join(url.Path, self.In.Endpoint)
	self.Out.Url = url.String()
	// Prepare request
	dataBufferP := bytes.NewBufferString(self.In.Data)
	req, err := http.NewRequest(
		self.In.Method,
		self.Out.Url,
		dataBufferP)
	if err != nil {
		self.Out.Err = err.Error()
	}
	// Set request headers
	req.Header = *self.In.Header
	// Request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		self.Out.Err = err.Error()
	}
	// Response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		self.Out.Err = err.Error()
	}
	res.Body.Close()
	// Fill in self.Out
	self.Out.Body = &body
	self.Out.Header = &res.Header
	self.Out.Status = res.Status
	// if http status code == 2XX
	self.Out.Success = (self.Out.Status[0] == '2')

	// Unmarshal
	self.ProcessOutput()

	helper.ReportDebug(&self, "api", false)

	return self.Out.Success
}

// GitApi Get action wrapper
func (self *GitApi[GitApiInfo]) Get() bool {
	self.In.Method = http.MethodGet
	return self.Do()
}

// GitApi Del action wrapper
func (self *GitApi[GitApiInfo]) Del() bool {
	self.In.Method = http.MethodDelete
	return self.Do()
}

// GitApi Patch action wrapper
func (self *GitApi[GitApiInfo]) Patch() bool {
	self.In.Method = http.MethodPatch
	return self.Do()
}

// GitApi Post action wrapper
func (self *GitApi[GitApiInfo]) Post() bool {
	self.In.Method = http.MethodPost
	return self.Do()
}

// GitApi Put action wrapper
func (self *GitApi[GitApiInfo]) Put() bool {
	self.In.Method = http.MethodPut
	return self.Do()
}

// Print both Body and Err into string pointer
func (self *GitApi[GitApiInfo]) ProcessOutput() {
	// Unmarshal
	err := json.Unmarshal(*self.Out.Body, self.Info)
	if self.Out.Success && err == nil && self.Info != nil {
		// Use Info string func
		self.Out.Output = self.Info.StringP()
	} else {
		var output string
		strP := helper.ReportSp(self.Out.Body, "", true)
		if strP != nil {
			output += *strP
		}
		strP = helper.ReportSp(self.Out.Err, "", true)
		if strP != nil {
			output += *strP
		}
		self.Out.Output = &output
	}
}
