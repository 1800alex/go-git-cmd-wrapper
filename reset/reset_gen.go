package reset

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import "github.com/1800alex/go-git-cmd-wrapper/v2/types"

// Hard Resets the index and working tree. Any changes to tracked files in the working tree since <commit> are discarded.
// --hard
func Hard(g *types.Cmd) {
	g.AddOptions("--hard")
}

// Keep Resets index entries and updates files in the working tree that are different between <commit> and HEAD. If a file that is different between <commit> and HEAD has local changes, reset is aborted.
// --keep
func Keep(g *types.Cmd) {
	g.AddOptions("--keep")
}

// Merge Resets the index and updates the files in the working tree that are different between <commit> and HEAD, but keeps those which are different between the index and working tree (i.e. which have changes which have not been added).
// If a file that is different between <commit> and the index has unstaged changes, reset is aborted.
//  In other words, --merge does something like a git read-tree -u -m <commit>, but carries forward unmerged index entries.
// --merge
func Merge(g *types.Cmd) {
	g.AddOptions("--merge")
}

// Mixed Resets the index but not the working tree (i.e., the changed files are preserved but not marked for commit) and reports what has not been updated.
// This is the default action.
// If -N is specified, removed paths are marked as intent-to-add (see git-add(1)).
// --mixed
func Mixed(g *types.Cmd) {
	g.AddOptions("--mixed")
}

// Quiet Be quiet, only report errors.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// Soft Does not touch the index file or the working tree at all (but resets the head to <commit>, just like all modes do).
// This leaves all your changed files 'Changes to be committed', as git status would put it.
// --soft
func Soft(g *types.Cmd) {
	g.AddOptions("--soft")
}
