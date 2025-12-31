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

package base

import (
	"net/http"
	"path"

	"github.com/J-Siu/go-restapi"
)

// Base
type Base struct {
	*Property
	*restapi.Api
}

// Setup a *GitApi
func (t *Base) New(property *Property) *Base {
	t.Property = property
	apiProperty := restapi.Property{
		Debug:      t.Debug,
		EntryPoint: t.EntryPoint,
		Info:       t.Info,
		SkipVerify: t.SkipVerify,
	}
	t.Api = new(restapi.Api).New(&apiProperty)
	t.HeaderGithub()
	return t
}

func (t *Base) Do() *Base {
	t.Api.Do()
	return t
}

// Initialize endpoint /user/repos
func (t *Base) EndpointUserRepos() *Base {
	t.Req.Endpoint = "/user/repos"
	return t
}

// Initialize endpoint /repos/OWNER/REPO
//
// Use current directory if GitApi.Repo is empty
func (t *Base) EndpointRepos() *Base {
	t.Req.Endpoint = path.Join("repos", t.User, *t.Repo())
	return t
}

// Initialize endpoint /repos/OWNER/REPO/topics
func (t *Base) EndpointReposTopics() *Base {
	t.Req.Endpoint = path.Join(t.EndpointRepos().Req.Endpoint, "topics")
	return t
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets
func (t *Base) EndpointReposSecrets() *Base {
	t.Req.Endpoint = path.Join(t.EndpointRepos().Req.Endpoint, "actions", "secrets")
	return t
}

// Initialize endpoint /repos/OWNER/REPO/actions/secrets/public-key
func (t *Base) EndpointReposSecretsPubkey() *Base {
	t.Req.Endpoint = path.Join(t.EndpointReposSecrets().Req.Endpoint, "public-key")
	return t
}

// Initialize endpoint /repos/OWNER/REPO/actions/permission
func (t *Base) EndpointReposActionsGithub() *Base {
	t.Req.Endpoint = path.Join(t.EndpointRepos().Req.Endpoint, "actions", "permissions")
	return t
}

// Set github/gitea header
//
// GitApi.Req.Token, if empty, authorization header will not be set.
func (t *Base) HeaderGithub() *Base {
	header := make(http.Header)
	t.Req.Header = &header
	t.Req.Header.Add("Accept", "application/vnd.github+json")
	t.Req.Header.Add("Content-Type", "application/json")
	t.Req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	if t.Token != "" {
		t.Req.Header.Add("Authorization", "token "+t.Token)
	}
	return t
}

// Setup empty API header
func (t *Base) HeaderInit() *Base {
	header := make(http.Header)
	t.Req.Header = &header
	return t
}

func (t *Base) SetGet() *Base {
	t.Api.SetGet()
	return t
}

func (t *Base) SetDel() *Base {
	t.Api.SetDel()
	return t
}

func (t *Base) SetPatch() *Base {
	t.Api.SetPatch()
	return t
}

func (t *Base) SetPost() *Base {
	t.Api.SetPost()
	return t
}

func (t *Base) SetPut() *Base {
	t.Api.SetPut()
	return t
}

func (t *Base) Err() *string    { return t.Api.Err() }
func (t *Base) Name() string    { return t.Property.Name }
func (t *Base) Ok() bool        { return t.Api.Ok() }
func (t *Base) Output() *string { return t.Api.Output() }
func (t *Base) Repo() *string   { return &t.Property.Repo }
