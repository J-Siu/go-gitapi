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

	"github.com/J-Siu/go-helper"
)

// GitApi input structure
type GitApiIn struct {
	Data       string       `json:"Data"` // Json marshaled Info
	Endpoint   string       `json:"Endpoint"`
	Entrypoint string       `json:"Entrypoint"` // Api base url
	Header     *http.Header `json:"Header"`     // Http request header
	Method     string       `json:"Method"`
	Token      string       `json:"Token"`
}

// GitApi output structure
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
	In     GitApiIn  `json:"In"`
	Out    GitApiOut `json:"Out"`
	Name   string    `json:"Name"`   // Name of connection
	User   string    `json:"User"`   // Api username
	Vendor string    `json:"Vendor"` // github/gitea
	Info   T         `json:"Info"`   // Pointer to structure
}

func GitApiNew[T GitApiInfo](name string, token string, entrypoint string, user string, vendor string, info T) *GitApi[T] {
	var self GitApi[T]
	self.Name = name
	self.User = user
	self.Vendor = vendor
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
func (self *GitApi[GitApiInfo]) EndpointRepos() {
	self.In.Endpoint = "/repos/" + self.User + "/" + helper.CurrentDirBase()
}

// Initialize endpoint /repos/OWNER/REPO/topics
func (self *GitApi[GitApiInfo]) EndpointReposTopics() {
	self.EndpointRepos()
	self.In.Endpoint += "/topics"
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets
func (self *GitApi[GitApiInfo]) EndpointReposSecrets() {
	self.EndpointRepos()
	self.In.Endpoint += "/actions/secrets"
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets/public-key
func (self *GitApi[GitApiInfo]) EndpointReposSecretsPubkey() {
	self.EndpointRepos()
	self.In.Endpoint += "/actions/secrets/public-key"
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
	// Prepare request
	self.Out.Url = self.In.Entrypoint + self.In.Endpoint
	dataBufferP := bytes.NewBufferString(self.In.Data)
	req, err := http.NewRequest(
		self.In.Method,
		self.Out.Url,
		dataBufferP)
	if err != nil {
		self.Out.Err = err.Error()
	}
	// Set headers
	self.In.Header = &req.Header
	self.In.Header.Add("Accept", "application/vnd.github.v3+json")
	self.In.Header.Add("Content-Type", "application/json")
	if len(self.In.Token) > 0 {
		self.In.Header.Add("Authorization", "token "+self.In.Token)
	}
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
	// check status code == 2XX
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
