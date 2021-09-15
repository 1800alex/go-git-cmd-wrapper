package lfs

import "github.com/1800alex/go-git-cmd-wrapper/v2/types"

// Display the Git LFS environment.
// usage: git lfs env
func Env(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("env")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Populate working copy with real content from Git LFS files.
// usage: git lfs checkout
func Checkout(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("checkout")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// De-duplicate Git LFS files.
// usage: git lfs dedup
func Dedup(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("dedup")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Display Git LFS extension details.
// usage: git lfs ext
func Ext(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("ext")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Download Git LFS files from a remote.
// usage: git lfs fetch
func Fetch(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("fetch")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Check Git LFS files for consistency.
// usage: git lfs fsck
func Fsck(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("fsck")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Install Git LFS configuration.
// usage: git lfs install
func Install(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("install")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Set a file as "locked" on the Git LFS server.
// usage: git lfs lock
func Lock(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("lock")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// List currently "locked" files from the Git LFS server.
// usage: git lfs locks
func Locks(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("locks")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Show errors from the Git LFS command.
// usage: git lfs logs
func Logs(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("logs")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Show information about Git LFS files in the index and working tree.
// usage: git lfs ls-files
func LsFiles(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("ls-files")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Migrate history to or from Git LFS
// usage: git lfs migrate
func Migrate(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("migrate")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Delete old Git LFS files from local storage
// usage: git lfs prune
func Prune(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("prune")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Fetch Git LFS changes from the remote & checkout any required working tree files.
// usage: git lfs pull
func Pull(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("pull")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Push queued large files to the Git LFS endpoint.
// usage: git lfs push
func Push(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("push")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Show the status of Git LFS files in the working tree.
// usage: git lfs status
func Status(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("status")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// View or add Git LFS paths to Git attributes.
// usage: git lfs track
func Track(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("track")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Uninstall Git LFS by removing hooks and smudge/clean filter configuration.
// usage: git lfs uninstall
func Uninstall(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("uninstall")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Remove "locked" setting for a file on the Git LFS server.
// usage: git lfs unlock
func Unlock(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("unlock")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Remove Git LFS paths from Git Attributes.
// usage: git lfs untrack
func Untrack(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("untrack")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Update Git hooks for the current Git repository.
// usage: git lfs update
func Update(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("update")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Report the version number.
// usage: git lfs version
func Version(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("version")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Low level commands (plumbing)

// Git clean filter that converts large files to pointers.
// usage: git lfs clean
func Clean(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("clean")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Git process filter that converts between large files and pointers.
// usage: git lfs process
func Process(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("process")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Build and compare pointers.
// usage: git lfs pointer
func Pointer(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("pointer")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Git post-checkout hook implementation.
// usage: git lfs post-checkout
func PostCheckout(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("post-checkout")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Git post-commit hook implementation.
// usage: git lfs post-commit
func PostCommit(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("post-commit")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Git post-merge hook implementation.
// usage: git lfs post-merge
func PostMerge(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("post-merge")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Git pre-merge hook implementation.
// usage: git lfs pre-merge
func PreMerge(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("pre-merge")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Git pre-push hook implementation.
// usage: git lfs pre-push
func PrePush(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("pre-push")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Git smudge filter that converts pointer in blobs to the actual content.
// usage: git lfs smudge
func Smudge(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("smudge")

		if object != "" {
			g.AddOptions(object)
		}
	}
}

// Git LFS standalone transfer adapter for file URLs (local paths).
// usage: git lfs standalone-file
func StandaloneFile(object string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions("standalone-file")

		if object != "" {
			g.AddOptions(object)
		}
	}
}
