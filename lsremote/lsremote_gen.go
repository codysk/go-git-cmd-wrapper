package lsremote

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/codysk/go-git-cmd-wrapper/v2/types"
)

// ExitCode Exit with status "2" when no matching refs are found in the remote repository. Usually the command exits with status "0" to indicate it successfully talked with the remote repository, whether it found any matching refs.
// --exit-code
func ExitCode(g *types.Cmd) {
	g.AddOptions("--exit-code")
}

// GetUrl Expand the URL of the given remote repository taking into account any "url.<base>.insteadOf" config setting (See git-config[1]) and exit without talking to the remote.
// --get-url
func GetUrl(g *types.Cmd) {
	g.AddOptions("--get-url")
}

// Heads Limit to only refs/heads and refs/tags, respectively. These options are not mutually exclusive; when given both, references stored in refs/heads and refs/tags are displayed. Note that git ls-remote -h used without anything else on the command line gives help, consistent with other git subcommands.
// -h, --heads
func Heads(g *types.Cmd) {
	g.AddOptions("--heads")
}

// Quiet Do not print remote URL to stderr.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// Refs Do not show peeled tags or pseudorefs like HEAD in the output.
// --refs
func Refs(g *types.Cmd) {
	g.AddOptions("--refs")
}

// ServerOption Transmit the given string to the server when communicating using protocol version 2. The given string must not contain a NUL or LF character. When multiple --server-option=<option> are given, they are all sent to the other side in the order listed on the command line.
// -o <option>, --server-option=<option>
func ServerOption(option string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--server-option=%s", option))
	}
}

// Sort Sort based on the key given. Prefix - to sort in descending order of the value. Supports "version:refname" or "v:refname" (tag names are treated as versions). The "version:refname" sort order can also be affected by the "versionsort.suffix" configuration variable. See git-for-each-ref[1] for more sort options, but be aware keys like committerdate that require access to the objects themselves will not work for refs whose objects have not yet been fetched from the remote, and will give a missing object error.
// --sort=<key>
func Sort(key string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--sort=%s", key))
	}
}

// Symref In addition to the object pointed by it, show the underlying ref pointed by it when showing a symbolic ref. Currently, upload-pack only shows the symref HEAD, so it will be the only one shown by ls-remote.
// --symref
func Symref(g *types.Cmd) {
	g.AddOptions("--symref")
}

// UploadPack Specify the full path of git-upload-pack on the remote host. This allows listing references from repositories accessed via SSH and where the SSH daemon does not use the PATH configured by the user.
// --upload-pack=<exec>
func UploadPack(exec string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--upload-pack=%s", exec))
	}
}
