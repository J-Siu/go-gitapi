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
	"strconv"

	"github.com/J-Siu/go-gitapi/v3/base"
	"github.com/J-Siu/go-gitapi/v3/info"
)

// Github repository(creation) info structure
type InfoList struct {
	*base.Base
	Info info.InfoList
}

func (t *InfoList) New(property *base.Property, vendor base.Vendor, page int) *InfoList {
	property.Info = &t.Info
	t.Base = new(base.Base).New(property).HeaderGithub().EndpointUserRepos()

	t.Req.UrlValInit()
	// switch vendor {
	// case base.VendorGithub:
	t.Req.UrlVal.Add("per_page", strconv.Itoa(100)) // github
	// case base.VendorGitea:
	t.Req.UrlVal.Add("limit", strconv.Itoa(100)) //gitea
	// }
	t.Req.UrlVal.Add("page", strconv.Itoa(page))
	*t.Repo() = ""
	return t
}
func (t *InfoList) Get() *InfoList {
	t.SetGet()
	return t
}
