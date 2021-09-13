# Go Git Cmd Wrapper

[![Build Status](https://github.com/1800alex/go-git-cmd-wrapper/workflows/Main/badge.svg?branch=master)](https://github.com/1800alex/go-git-cmd-wrapper/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/1800alex/go-git-cmd-wrapper)](https://pkg.go.dev/github.com/1800alex/go-git-cmd-wrapper/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/1800alex/go-git-cmd-wrapper)](https://goreportcard.com/report/github.com/1800alex/go-git-cmd-wrapper)

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink.svg)](https://github.com/sponsors/1800alex)

It's a simple wrapper around `git` command.

Import `github.com/1800alex/go-git-cmd-wrapper/v2/git`.

```go
// clone
output, err := git.Clone(clone.Repository("https://github.com/1800alex/prm"))
// with debug option
output, err := git.Clone(clone.Repository("https://github.com/1800alex/prm"), git.Debug)
output, err := git.Clone(clone.Repository("https://github.com/1800alex/prm"), git.Debugger(true))

// fetch
output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"))
output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("master"))

// add a remote
output, err = git.Remote(remote.Add, remote.Name("upstream"), remote.URL("https://github.com/1800alex/prm"))
```

More examples: [Documentation](https://pkg.go.dev/github.com/1800alex/go-git-cmd-wrapper/v2/git)
