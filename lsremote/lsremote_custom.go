package lsremote

import "github.com/codysk/go-git-cmd-wrapper/v2/types"

// repository <repository>.
// The "remote" repository to query. This parameter can be either a URL or the name of a remote (see the GIT URLS and REMOTES sections of git-fetch[1]).
func Repository(url string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(url)
	}
}

// patterns [<patterns>…​]
// When unspecified, all references, after filtering done with --heads and --tags, are shown. When <patterns>…​ are specified, only references matching one or more of the given patterns are displayed. Each pattern is interpreted as a glob (see glob in gitglossary[7]) which is matched against the "tail" of a ref, starting either from the start of the ref (so a full name like refs/heads/foo matches) or from a slash separator (so bar matches refs/heads/bar but not refs/heads/foobar).
func Patterns(patterns ...string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		for _, pattern := range patterns {
			g.AddOptions(pattern)
		}
	}
}
