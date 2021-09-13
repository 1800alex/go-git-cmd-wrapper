/*
Package go_git_cmd_wrapper A simple wrapper around `git` command.

		import (
			// ...
			"github.com/1800alex/go-git-cmd-wrapper/v2/git"
			// ...
			"github.com/1800alex/go-git-cmd-wrapper/v2/clone"
			"github.com/1800alex/go-git-cmd-wrapper/v2/config"
			"github.com/1800alex/go-git-cmd-wrapper/v2/fetch"
			"github.com/1800alex/go-git-cmd-wrapper/v2/remote"
		)

		// clone
		output, err := git.Clone(clone.Repository("https://github.com/1800alex/gcg"))
		// with debug option
		output, err := git.Clone(clone.Repository("https://github.com/1800alex/gcg"), git.Debug)
		output, err := git.Clone(clone.Repository("https://github.com/1800alex/gcg"), git.Debugger(true))

		// fetch
		output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"))
		output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("master"))

		// add a remote
		output, err = git.Remote(remote.Add, remote.Name("upstream"), remote.URL("https://github.com/1800alex/gcg"))

*/
package go_git_cmd_wrapper
