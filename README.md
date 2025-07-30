# go-gitapi  [![Paypal donate](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/donate/?business=HZF49NM9D35SJ&no_recurring=0&currency_code=CAD)

Golang Github/Gitea api library.

### Table Of Content
<!-- TOC -->

- [Table Of Content](#table-of-content)
- [Features](#features)
  - [gitApi.go](#gitapigo)
  - [gitApiDataStruct.go](#gitapidatastructgo)
- [Pro](#pro)
- [Doc](#doc)
- [Dependency](#dependency)
- [Supported git repository services](#supported-git-repository-services)
- [Usage Example](#usage-example)
  - [Debug](#debug)
- [Used By Project](#used-by-project)
- [Repository](#repository)
- [Contributors](#contributors)
- [Change Log](#change-log)
- [License](#license)

<!-- /TOC -->
<!--more-->
### Features

- API action
  - [x] Do
  - [x] Get
  - [x] Del
  - [x] Patch
  - [x] Post
  - [x] Put

#### gitApi.go
- type GitApiReq struct
- func (gaReq *GitApiReq) UrlValInit()
- type GitApiRes struct
- func (gaRes *GitApiRes) Ok() bool
- type GitApi struct
- func GitApiNew(
- func (ga *GitApi) Ok() bool
- func (ga *GitApi) Output() *string
- func (ga *GitApi) Err() *string
- func (ga *GitApi) EndpointUserRepos() *GitApi
- func (ga *GitApi) EndpointRepos() *GitApi
- func (ga *GitApi) EndpointReposTopics() *GitApi
- func (ga *GitApi) EndpointReposSecrets() *GitApi
- func (ga *GitApi) EndpointReposSecretsPubkey() *GitApi
- func (ga *GitApi) HeaderGithub() *GitApi
- func (ga *GitApi) HeaderInit() *GitApi
- func (ga *GitApi) Do() *GitApi
- func (ga *GitApi) Get() *GitApi
- func (ga *GitApi) Del() *GitApi
- func (ga *GitApi) Patch() *GitApi
- func (ga *GitApi) Post() *GitApi
- func (ga *GitApi) Put() *GitApi
- func (ga *GitApi) SetGet() *GitApi
- func (ga *GitApi) SetDel() *GitApi
- func (ga *GitApi) SetPatch() *GitApi
- func (ga *GitApi) SetPost() *GitApi
- func (ga *GitApi) SetPut() *GitApi
- func (ga *GitApi) ProcessOutput() *GitApi
- func (ga *GitApi) ProcessOutputError() *GitApi
- func (ga *GitApi) ProcessError() *GitApi
#### gitApiDataStruct.go
- type RepoEncryptedPair struct
- func (rEncryptedPair *RepoEncryptedPair) StringP() *string
- func (rEncryptedPair *RepoEncryptedPair) String() string
- type RepoPublicKey struct
- func (rPKey *RepoPublicKey) StringP() *string
- func (rPKey *RepoPublicKey) String() string
- type RepoPrivate struct
- func (rPrivate *RepoPrivate) StringP() *string
- func (rPrivate *RepoPrivate) String() string
- type RepoVisibility struct
- func (rVisibility *RepoVisibility) StringP() *string
- func (rVisibility *RepoVisibility) String() string
- type RepoDescription struct
- func (rDesc *RepoDescription) StringP() *string
- func (rDesc *RepoDescription) String() string
- type RepoTopics struct
- func (rTopics *RepoTopics) StringP() *string
- func (rTopics *RepoTopics) String() string
- type RepoInfo struct
- func (rInfo *RepoInfo) StringP() *string
- func (rInfo *RepoInfo) String() string
- type RepoError struct
- func (rError *RepoError) StringP() *string
- func (rError *RepoError) String() string
- type RepoInfoList []RepoInfo
- func (rInfoList *RepoInfoList) StringP() *string
- func (rInfoList *RepoInfoList) String() string
- type NilType struct
- func (nilT *NilType) StringP() *string
- func (nilT *NilType) String() string
- func Nil() *NilType
- type GitApiInfo interface

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
    // Get instance
    gitApi := gitapi.GitApiNew(
      "Test",   // Connection name for debug print out purpose
      "01234567890123456789012345678912", // API token,
      "https://api.github.com", // API entrypoint
      "User",  // user
      "github", // vendor/brand
      &info)    // data for request
    // Setup endpoint
    gitApi.EndpointRepos()
    // Setup Github header
    gitApi.HeaderGithub()
    // Set http method: Post
    gitApi.SetPost()
    // Do request
    success := gitApi.Do().Ok()
    ```

3. Print out using helper function
    ```go
    helper.ReportStatus(success, gitApi.Name, false, true)
    helper.Report(gitApi.Output(), gitApi.Name, false, true)
    ```

#### Debug

Enable debug
```go
helper.Debug = true
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

### License

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
