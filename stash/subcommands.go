package stash

import "github.com/1800alex/go-git-cmd-wrapper/v2/types"

// Show the changes recorded in the stash entry as a diff between the stashed contents and the commit back when the stash entry was first created.
// usage: git stash push [<options>] [<stash>]
func Push(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("push")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// List the stash entries that you currently have.
// usage: git stash list [<options>]
func List(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("list")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Show the changes recorded in the stash entry as a diff between the stashed contents and the commit back when the stash entry was first created.
// usage: git stash show [<options>] [<stash>]
func Show(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("show")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Remove a single stashed state from the stash list and apply it on top of the current working tree state
// usage: git stash pop [<stash>]
func Pop(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("pop")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Like pop, but do not remove the state from the stash list.
// usage: git stash apply [<stash>]
func Apply(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("apply")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Creates and checks out a new branch named <branchname> starting from the commit at which the <stash> was originally created, applies the changes recorded in <stash> to the new working tree and index.
// usage: git stash branch <branchname> [<stash>]
func Branch(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("branch")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Remove all the stash entries.
// usage: git stash clear
func Clear(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("clear")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Remove a single stash entry from the list of stash entries.
// usage: git stash drop
func Drop(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("drop")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Create a stash entry (which is a regular commit object) and return its object name, without storing it anywhere in the ref namespace.
// usage: git stash create
func Create(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("create")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

//Store a given stash created via git stash create (which is a dangling merge commit) in the stash ref, updating the stash reflog.
// usage: git stash store
func Store(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("store")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// GetRef Print the current notes ref. This provides an easy way to retrieve the current notes ref (e.g. from scripts).
// usage: git notes get-ref
func GetRef(opts ...types.Option) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("get-ref")
	}
}
