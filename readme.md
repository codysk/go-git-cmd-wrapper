# Go Git Cmd Wrapper

> fork from github.com/ldez/go-git-cmd-wrapper

[![Build Status](https://github.com/codysk/go-git-cmd-wrapper/workflows/Main/badge.svg?branch=main)](https://github.com/codysk/go-git-cmd-wrapper/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/codysk/go-git-cmd-wrapper)](https://pkg.go.dev/github.com/codysk/go-git-cmd-wrapper/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/codysk/go-git-cmd-wrapper)](https://goreportcard.com/report/github.com/codysk/go-git-cmd-wrapper)

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink.svg)](https://github.com/sponsors/codysk)

It's a simple wrapper around `git` command.

Import `github.com/codysk/go-git-cmd-wrapper/v2/git`.

```go
// clone
output, err := git.Clone(clone.Repository("https://github.com/codysk/prm"))
// with debug option
output, err := git.Clone(clone.Repository("https://github.com/codysk/prm"), git.Debug)
output, err := git.Clone(clone.Repository("https://github.com/codysk/prm"), git.Debugger(true))

// fetch
output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"))
output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("master"))

// add a remote
output, err = git.Remote(remote.Add, remote.Name("upstream"), remote.URL("https://github.com/codysk/prm"))
```

More examples: [Documentation](https://pkg.go.dev/github.com/codysk/go-git-cmd-wrapper/v2/git)
