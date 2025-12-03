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
	"testing"

	"github.com/J-Siu/go-gitapi/v2/gitapi"
	"github.com/J-Siu/go-gitapi/v2/repo"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/strany"
)

func TestGetGithubRepository(t *testing.T) {

	// helper.Debug = true

	var (
		repoList repo.InfoList
		property = gitapi.Property{
			// Debug:      true,
			EntryPoint: "https://api.github.com",
			Info:       &repoList,
			Name:       "Test",
			SkipVerify: false,
		}
		gitApi = gitapi.New(&property)
		req    = gitApi.Req
		res    = gitApi.Res
	)
	// Setup endpoint
	req.Endpoint = "repositories"
	// Setup Github header
	gitApi.HeaderGithub()

	// Get request
	success := gitApi.Get().Res.Ok()
	ezlog.Log().N("List").Lm(res.Output).Out()
	ezlog.Log().N("Url").Lm(res.Url).Out()
	ezlog.Log().Lm(res.Url.String()).Out()
	ezlog.Log().N("Count").M(len(repoList)).Out()

	if !success {
		t.Fatalf("Failed:\n%s", *strany.Any(gitApi))
	}
}
