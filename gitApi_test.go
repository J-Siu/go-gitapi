package gitapi

import (
	"testing"

	"github.com/J-Siu/go-helper"
)

func TestGetGithubRepository(t *testing.T) {

	// helper.Debug = true

	var repoList RepoInfoList

	// Get instance
	// https://api.github.com/repositories
	gitApi := GitApiNew(
		"Test",                   // Connection name for debug print out purpose
		"",                       // API token,
		"https://api.github.com", // API entrypoint
		"",                       // user
		"",                       // vendor/brand
		"",                       // Repo
		&repoList)                // data for request
	// Setup endpoint
	gitApi.In.Endpoint = "repositories"
	// Setup Github header
	gitApi.HeaderGithub()
	gitApi.In.UrlValInit()
	//gitApi.In.UrlVal.Add("per_page", "75")

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
