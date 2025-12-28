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

package api

import (
	"path"

	"github.com/J-Siu/go-gitapi/v3/base"
	"github.com/J-Siu/go-gitapi/v3/info"
)

// Github repository(creation) info structure
type Repo struct {
	*base.Base
	Info info.Info
}

func (t *Repo) New(property *base.Property) *Repo {
	property.Info = &t.Info
	t.Base = new(base.Base).New(property)
	return t
}

// Set action: create
func (t *Repo) Create() *Repo {
	t.EndpointUserRepos().SetPost()
	return t
}

// Set action: delete
func (t *Repo) Del() *Repo {
	t.Base.Info = nil
	t.Base.Api.Info = nil
	t.EndpointRepos().SetDel()
	return t
}

func (t *Repo) DelSecret(secret string) *Repo {
	t.EndpointReposSecrets().SetDel()
	t.Req.Endpoint = path.Join(t.Req.Endpoint, secret)
	return t
}
