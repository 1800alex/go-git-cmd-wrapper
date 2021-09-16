package stash

import (
	"github.com/1800alex/go-git-cmd-wrapper/v2/types"
)

// HyphenHyphen add `--`
// This option can be used to separate command-line options from the list of files, (useful when filenames might be mistaken for command-line options).
func HyphenHyphen(g *types.Cmd) {
	g.AddOptions("--")
}
