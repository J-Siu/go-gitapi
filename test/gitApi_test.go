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
	"testing"

	"github.com/J-Siu/go-gitapi"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/str"
)

func TestGetGithubRepository(t *testing.T) {

	// helper.Debug = true

	var repoList gitapi.RepoInfoList

	// var gitApi = &GitApi{
	// 	Name: "Test",
	// 	Info: &repoList,
	// 	In: GitApiIn{
	// 		Entrypoint: "https://api.github.com",
	// 	},
	// }
	var gitApi = gitapi.GitApiNew("Test", "", "https://api.github.com", "", "", false, "", &repoList)

	// Setup endpoint
	gitApi.Req.Endpoint = "repositories"
	// Setup Github header
	gitApi.HeaderGithub()

	// Get request
	success := gitApi.Get().Res.Ok()
	ezlog.Log().NameLn("List").Msg(gitApi.Res.Output).Out()
	ezlog.Log().NameLn("Url").Msg(gitApi.Res.Url).Out()
	ezlog.Log().Msg(gitApi.Res.Url.String()).Out()
	ezlog.Log().NameLn("Count").Msg(len(repoList)).Out()

	if !success {
		t.Fatalf("Failed:\n%s", new(str.Any).New().Str(gitApi))
	}
}
