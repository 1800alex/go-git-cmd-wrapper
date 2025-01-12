/*
Package lfs git-lfs - Work with large files in Git repositories.

SYNOPSIS

Reference: https://git-lfs.github.com/

		git lfs command [args]

DESCRIPTION

Git  LFS  is  a  system  for  managing and versioning large files in association with a Git repository.
Instead of storing the large files within the Git repository as blobs, Git LFS stores special  "pointer
files"  in  the repository, while storing the actual file contents on a Git LFS server. The contents of
the large file are downloaded automatically when needed, for example when a Git branch  containing  the
large file is checked out.

Git  LFS works by using a "smudge" filter to look up the large file contents based on the pointer file,
and a "clean" filter to create a new version of the pointer file when the large file's contents change.
It  also uses a pre-push hook to upload the large file contents to the Git LFS server whenever a commit
containing a new large file version is about to be pushed to the corresponding Git server.

COMMANDS

Like Git, Git LFS commands are separated into high level ("porcelain") commands and low level  ("plumb-
ing") commands.

*/
package lfs
