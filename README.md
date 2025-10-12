# go-gitapi  [![Paypal donate](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/donate/?business=HZF49NM9D35SJ&no_recurring=0&currency_code=CAD)

Golang Github/Gitea api library using [go-restapi](https://github.com/J-Siu/go-restapi).

### Table Of Content
<!-- TOC -->

- [Pro](#pro)
- [Doc](#doc)
- [Dependency](#dependency)
- [Supported git repository services](#supported-git-repository-services)
- [Usage Example](#usage-example)
- [Used By Project](#used-by-project)
- [Repository](#repository)
- [Contributors](#contributors)
- [Change Log](#change-log)
- [License](#license)

<!-- /TOC -->
<!--more-->

#### api.go

```go
func (t *GitApi) EndpointRepos() *GitApi
func (t *GitApi) EndpointReposActionsGithub() *GitApi
func (t *GitApi) EndpointReposSecrets() *GitApi
func (t *GitApi) EndpointReposSecretsPubkey() *GitApi
func (t *GitApi) EndpointReposTopics() *GitApi
func (t *GitApi) EndpointUserRepos() *GitApi
func (t *GitApi) HeaderGithub() *GitApi
func (t *GitApi) HeaderInit() *GitApi
func (t *GitApi) New(property *Property) *GitApi
func New(property *Property) *GitApi
```

```go
type IInfo interface
```

#### Repo Struct

```go
type EncryptedPair struct
type PublicKey struct
type Private struct
type Visibility struct
type Description struct
type Topics struct
type Info struct
type Error struct
type InfoList []IInfo
```

### Pro

- Easy to extend
- Small size

### Doc

- https://pkg.go.dev/github.com/J-Siu/go-gitapi

### Dependency

- [go-helper](https://github.com/J-Siu/go-helper)

### Supported git repository services
- gitea
- github
- gogs

### Usage Example

Following is code to create a new repository:

1. Prepare a GitApi data structure
    ```go
    var info gitapi.RepoInfo
    info.Name = "test"
    info.Private = remote.Private
    ```

2. Setup and execute
    ```go
    property = gitapi.Property{
      // Debug:      true,
      EntryPoint: "https://api.github.com",
      Info:       &repoList,
      Name:       "Test",
      User:       "User",
      Token:      "01234567890123456789012345678912",
      Vendor:     gitapi.VendorGithub,
      SkipVerify: false,
    }

    // Get instance
    gitApi := gitapi.New(&property)
    // Setup endpoint
    gitApi.EndpointRepos()
    // Setup Github header
    gitApi.HeaderGithub()
    // Set http method: Post
    gitApi.SetPost()
    // Do request
    success := gitApi.Do().Ok()
    ```

### Used By Project

- [go-mygit](https://github.com/J-Siu/go-mygit)
### Repository

- [go-gitapi](https://github.com/J-Siu/go-gitapi)

### Contributors

- [John, Sing Dao, Siu](https://github.com/J-Siu)

### Change Log

- v1.0.0
  - Feature complete
- v1.0.1
  - Fix data struct *.StringP() output
- v1.1.0
  - Consolidate output processing
- v1.2.0
  - GitApiOut
    - move output from GitApi
    - add Success
    GitApi struct
    - Change Header to non-pointer
    - Use path and url package to handle endpoint and url
    - Add HeaderGithub()
- v1.2.4
  - Update go-helper package for bug fix
- v1.2.5
  - GitApiIn
    - Add UrlVal(url.Values)
- v1.2.6
  - Update go-helper package for bug fix
- v1.3.0
  - All GitApi methods reutrn self pointer
  - Interface GitApiInfo remove type restrictions
  - Member GitApi.In -> GitApi.Req
  - Member GitApi.Out -> GitApi.Res
  - Type GitApiIn -> GitApiReq
  - Type GitApiOut -> GitApiRes
- v1.3.1
  - Improve README.md
- v1.3.2
  - Fix `GitApi.Do()` wiping `GitApi.Req.Data` if `GitApi.Info` is `nil`
- v1.4.0
  - upgrade helper to 1.1.6
  - `GitApi` struct and `GitApiMew()`
    - add `SkipVerify` to support self-signed cert
- v1.4.1
  - GitApi.ProcessOutput() - Fix output for full info
  - Use proper receiver name
- v1.5.0
  - Add RepoError struct
  - Handle http error: ProcessError()
  - Handle API error: ProcessOutputError()
- v1.6.0
  - gitApi.go
    - Move Method(http) from GitApiReq -> GitApi
    - GitApi struct
      - Add SetGet(), SetDel(), SetPatch(), SetPost(), SetPut()
      - Add .Res wrapper func Err(), Ok(), Output()
- v1.6.1
  - Update go-helper
- v1.6.2
  - Update go-helper/v2
- v1.6.3
  - Update go-helper/v2
- v1.6.4
  - Update go-helper/v2
  - Add GitApi.New()
  - Breaking: Standardize package level GitApiNew() -> New()
- v2.0.0
  - Split rest api to [go-restapi](https://github.com/J-Siu/go-restapi).
- v2.0.1
  - Add repo archived support
- v2.0.2
  - Update go-helper/v2, go-restapi
- v2.1.0
  - Add repo api support for discussions, projects and wiki
- v2.1.1
  - Add repo api support for actions and `EndpointReposActionsGithub`

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
