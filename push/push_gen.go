package push

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// All Push all branches (i.e. refs under refs/heads/); cannot be used with other <refspec>.
// --all
func All(g *types.Cmd) {
	g.AddOptions("--all")
}

// Atomic Use an atomic transaction on the remote side if available.
// Either all refs are updated, or on error, no refs are updated.
// If the server does not support atomic pushes the push will fail.
// --atomic
func Atomic(g *types.Cmd) {
	g.AddOptions("--atomic")
}

// Delete All listed refs are deleted from the remote repository.
// This is the same as prefixing all refs with a colon.
// --delete
func Delete(g *types.Cmd) {
	g.AddOptions("--delete")
}

// DryRun Do everything except actually send the updates.
// -n, --dry-run
func DryRun(g *types.Cmd) {
	g.AddOptions("--dry-run")
}

// Exec Path to the git-receive-pack program on the remote end.
// Sometimes useful when pushing to a remote repository over ssh, and you do not have the program in a directory on the default $PATH.
// --exec=<git-receive-pack>
func Exec(gitReceivePack string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--exec=%s", gitReceivePack))
	}
}

// FollowTags Push all the refs that would be pushed without this option, and also push annotated tags in refs/tags that are missing from the remote but are pointing at commit-ish that are reachable from the refs being pushed.
// This can also be specified with configuration variable push.followTags.
// For more information, see push.followTags in git-config(1).
// --follow-tags
func FollowTags(g *types.Cmd) {
	g.AddOptions("--follow-tags")
}

// Force Usually, the command refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it.
// Also, when --force-with-lease option is used, the command refuses to update a remote ref whose current value does not match what is expected.
// -f, --force
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// Ipv4 Use IPv4 addresses only, ignoring IPv6 addresses.
// -4, --ipv4
func Ipv4(g *types.Cmd) {
	g.AddOptions("--ipv4")
}

// Ipv6 Use IPv6 addresses only, ignoring IPv4 addresses.
// -6, --ipv6
func Ipv6(g *types.Cmd) {
	g.AddOptions("--ipv6")
}

// Mirror Instead of naming each ref to push, specifies that all refs under refs/ (which includes but is not limited to refs/heads/, refs/remotes/, and refs/tags/) be mirrored to the remote repository.
// Newly created local refs will be pushed to the remote end, locally updated refs will be force updated on the remote end, and deleted refs will be removed from the remote end.
// This is the default if the configuration option remote.<remote>.mirror is set.
// --mirror
func Mirror(g *types.Cmd) {
	g.AddOptions("--mirror")
}

// NoAtomic Use an atomic transaction on the remote side if available.
// Either all refs are updated, or on error, no refs are updated.
// If the server does not support atomic pushes the push will fail.
// --no-atomic
func NoAtomic(g *types.Cmd) {
	g.AddOptions("--no-atomic")
}

// NoRecurseSubmodules May be used to make sure all submodule commits used by the revisions to be pushed are available on a remote-tracking branch.
// If check is used Git will verify that all submodule commits that changed in the revisions to be pushed are available on at least one remote of the submodule.
// If any commits are missing the push will be aborted and exit with non-zero status.
// If on-demand is used all submodules that changed in the revisions to be pushed will be pushed.
// If on-demand was not able to push all necessary revisions it will also be aborted and exit with non-zero status.
// If only is used all submodules will be recursively pushed while the superproject is left unpushed.
// A value of no or using --no-recurse-submodules can be used to override the push.recurseSubmodules configuration variable when no submodule recursion is required.
// --no-recurse-submodules
func NoRecurseSubmodules(g *types.Cmd) {
	g.AddOptions("--no-recurse-submodules")
}

// NoSigned GPG-sign the push request to update refs on the receiving side, to allow it to be checked by the hooks and/or be logged.
// If false or --no-signed, no signing will be attempted.
// If true or --signed, the push will fail if the server does not support signed pushes.
// If set to if-asked, sign if and only if the server supports signed pushes.
// The push will also fail if the actual call to gpg --sign fails.
// See git-receive-pack(1) for the details on the receiving end.
// --no-signed
func NoSigned(g *types.Cmd) {
	g.AddOptions("--no-signed")
}

// NoThin These options are passed to git-send-pack(1).
// A thin transfer significantly reduces the amount of sent data when the sender and receiver share many of the same objects in common.
// The default is --thin.
// --no-thin
func NoThin(g *types.Cmd) {
	g.AddOptions("--no-thin")
}

// NoVerify Toggle the pre-push hook (see githooks(5)).
// The default is --verify, giving the hook a chance to prevent the push.
// With --no-verify, the hook is bypassed completely.
// --no-verify
func NoVerify(g *types.Cmd) {
	g.AddOptions("--no-verify")
}

// Porcelain Produce machine-readable output.
// The output status line for each ref will be tab-separated and sent to stdout instead of stderr.
// The full symbolic names of the refs will be given.
// --porcelain
func Porcelain(g *types.Cmd) {
	g.AddOptions("--porcelain")
}

// Progress Progress status is reported on the standard error stream by default when it is attached to a terminal, unless -q is specified.
// This flag forces progress status even if the standard error stream is not directed to a terminal.
// --progress
func Progress(g *types.Cmd) {
	g.AddOptions("--progress")
}

// Prune Remove remote branches that don’t have a local counterpart.
// For example a remote branch tmp will be removed if a local branch with the same name doesn’t exist any more. This also respects refspecs, e.g.  git push --prune remote refs/heads/*:refs/tmp/* would make sure that remote refs/tmp/foo will be removed if refs/heads/foo doesn’t exist.
// --prune
func Prune(g *types.Cmd) {
	g.AddOptions("--prune")
}

// PushOption Transmit the given string to the server, which passes them to the pre-receive as well as the post-receive hook.
// The given string must not contain a NUL or LF character.
// -o, --push-option
func PushOption(g *types.Cmd) {
	g.AddOptions("--push-option")
}

// Quiet Suppress all output, including the listing of updated refs, unless an error occurs.
// Progress is not reported to the standard error stream.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// ReceivePack Path to the git-receive-pack program on the remote end.
// Sometimes useful when pushing to a remote repository over ssh, and you do not have the program in a directory on the default $PATH.
// --receive-pack=<git-receive-pack>
func ReceivePack(gitReceivePack string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--receive-pack=%s", gitReceivePack))
	}
}

// RecurseSubmodules May be used to make sure all submodule commits used by the revisions to be pushed are available on a remote-tracking branch.
// If check is used Git will verify that all submodule commits that changed in the revisions to be pushed are available on at least one remote of the submodule.
// If any commits are missing the push will be aborted and exit with non-zero status.
// If on-demand is used all submodules that changed in the revisions to be pushed will be pushed.
// If on-demand was not able to push all necessary revisions it will also be aborted and exit with non-zero status.
// If only is used all submodules will be recursively pushed while the superproject is left unpushed.
// A value of no or using --no-recurse-submodules can be used to override the push.recurseSubmodules configuration variable when no submodule recursion is required.
// --recurse-submodules=(check|on-demand|only|no)
func RecurseSubmodules(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--recurse-submodules=%s", value))
	}
}

// Repo This option is equivalent to the <repository> argument.
// If both are specified, the command-line argument takes precedence.
// --repo=<repository>
func Repo(repository string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--repo=%s", repository))
	}
}

// SetUpstream For every branch that is up to date or successfully pushed, add upstream (tracking) reference, used by argument-less git-pull(1) and other commands.
// For more information, see branch.<name>.merge in git-config(1).
// -u, --set-upstream
func SetUpstream(g *types.Cmd) {
	g.AddOptions("--set-upstream")
}

// Sign GPG-sign the push request to update refs on the receiving side, to allow it to be checked by the hooks and/or be logged.
// If false or --no-signed, no signing will be attempted.
// If true or --signed, the push will fail if the server does not support signed pushes.
// If set to if-asked, sign if and only if the server supports signed pushes.
// The push will also fail if the actual call to gpg --sign fails.
// See git-receive-pack(1) for the details on the receiving end.
// --sign=(true|false|if-asked)
func Sign(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--sign=%s", value))
	}
}

// Signed GPG-sign the push request to update refs on the receiving side, to allow it to be checked by the hooks and/or be logged.
// If false or --no-signed, no signing will be attempted.
// If true or --signed, the push will fail if the server does not support signed pushes.
// If set to if-asked, sign if and only if the server supports signed pushes.
// The push will also fail if the actual call to gpg --sign fails.
// See git-receive-pack(1) for the details on the receiving end.
// --signed
func Signed(g *types.Cmd) {
	g.AddOptions("--signed")
}

// Tags All refs under refs/tags are pushed, in addition to refspecs explicitly listed on the command line.
// --tags
func Tags(g *types.Cmd) {
	g.AddOptions("--tags")
}

// Thin These options are passed to git-send-pack(1).
// A thin transfer significantly reduces the amount of sent data when the sender and receiver share many of the same objects in common.
// The default is --thin.
// --thin
func Thin(g *types.Cmd) {
	g.AddOptions("--thin")
}

// Verbose Run verbosely.
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

// Verify Toggle the pre-push hook (see githooks(5)).
// The default is --verify, giving the hook a chance to prevent the push.
// With --no-verify, the hook is bypassed completely.
// --verify
func Verify(g *types.Cmd) {
	g.AddOptions("--verify")
}

// WorkingDir Sets the working dir use for the git command.
// --working-dir=<dir>
func WorkingDir(dir string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--working-dir=%s", dir))
	}
}
