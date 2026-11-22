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

package gitApi_test

import (
	"fmt"
	"testing"

	"github.com/J-Siu/go-gitapi/v3/base"
	"github.com/J-Siu/go-gitapi/v3/info"
	"github.com/J-Siu/go-helper-mini/str"
)

func TestGetGithubRepository(t *testing.T) {
	var (
		repoList info.InfoList
		property = base.Property{
			EntryPoint: "https://api.github.com",
			Info:       &repoList,
			Name:       "Test",
			SkipVerify: false,
		}
		gitApi = base.New(&property)
		req    = gitApi.Api.Req
		res    = gitApi.Api.Res
	)
	// Setup endpoint
	req.Endpoint = "repositories"
	// Setup Github header
	gitApi.HeaderGithub()

	// Get request
	success := gitApi.Api.Get().Res.Ok()
	fmt.Println(str.Struct2String("List", res.Output, true))
	fmt.Println(str.Struct2String("Url", res.Url, true))
	fmt.Println(res.Url.String())
	fmt.Println("Count:", len(repoList))

	if !success {
		t.Fatalf("%s", str.Struct2String("Failed", gitApi, true))
	}
}
