package gitapi

import (
	"testing"

	"github.com/J-Siu/go-helper"
)

func TestGetGithubRepository(t *testing.T) {

	// helper.Debug = true

	var repoList RepoInfoList

	// var gitApi = &GitApi{
	// 	Name: "Test",
	// 	Info: &repoList,
	// 	In: GitApiIn{
	// 		Entrypoint: "https://api.github.com",
	// 	},
	// }
	var gitApi = GitApiNew("Test", "", "https://api.github.com", "", "", "", &repoList)

	// Setup endpoint
	gitApi.In.Endpoint = "repositories"
	// Setup Github header
	gitApi.HeaderGithub()

	// Get request
	success := gitApi.Get()
	helper.Report(gitApi.Out.Output, "List", true, false)
	helper.Report(gitApi.Out.Url, "Url", true, false)
	helper.Report(gitApi.Out.Url.String(), "", true, false)
	helper.Report(len(repoList), "Count", true, true)

	if !success {
		t.Fatalf(*helper.ReportSp(gitApi, "Failed", true, false))
	}
}
