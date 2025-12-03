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
	"net/http"
	"path"

	"github.com/J-Siu/go-restapi"
)

// GitApi
type GitApi struct {
	*restapi.Api `json:"rest_api,omitempty"`
	*Property
}

// Setup a *GitApi
func (t *GitApi) New(property *Property) *GitApi {
	apiProperty := restapi.Property{
		Debug:      property.Debug,
		EntryPoint: property.EntryPoint,
		Info:       property.Info,
		SkipVerify: property.SkipVerify,
	}
	t.Api = new(restapi.Api).New(&apiProperty)
	t.Property = property
	return t
}

// Setup a *GitApi
func New(property *Property) *GitApi {
	return new(GitApi).New(property)
}

// Initialize endpoint /user/repos
func (t *GitApi) EndpointUserRepos() *GitApi {
	t.Req.Endpoint = "/user/repos"
	return t
}

// Initialize endpoint /repos/OWNER/REPO
//
// Use current directory if GitApi.Repo is empty
func (t *GitApi) EndpointRepos() *GitApi {
	t.Req.Endpoint = path.Join("repos", t.User, t.Repo)
	return t
}

// Initialize endpoint /repos/OWNER/REPO/topics
func (t *GitApi) EndpointReposTopics() *GitApi {
	t.Req.Endpoint = path.Join(t.EndpointRepos().Req.Endpoint, "topics")
	return t
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets
func (t *GitApi) EndpointReposSecrets() *GitApi {
	t.Req.Endpoint = path.Join(t.EndpointRepos().Req.Endpoint, "actions", "secrets")
	return t
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets/public-key
func (t *GitApi) EndpointReposSecretsPubkey() *GitApi {
	t.Req.Endpoint = path.Join(t.EndpointReposSecrets().Req.Endpoint, "public-key")
	return t
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets/public-key
func (t *GitApi) EndpointReposActionsGithub() *GitApi {
	t.Req.Endpoint = path.Join(t.EndpointRepos().Req.Endpoint, "actions", "permissions")
	return t
}

// Set github/gitea header
//
// GitApi.Req.Token, if empty, authorization header will not be set.
func (t *GitApi) HeaderGithub() *GitApi {
	header := make(http.Header)
	t.Req.Header = &header
	t.Req.Header.Add("Accept", "application/vnd.github.v3+json")
	t.Req.Header.Add("Content-Type", "application/json")
	if t.Token != "" {
		t.Req.Header.Add("Authorization", "token "+t.Token)
	}
	return t
}

// Setup empty API header
func (t *GitApi) HeaderInit() *GitApi {
	header := make(http.Header)
	t.Req.Header = &header
	return t
}
