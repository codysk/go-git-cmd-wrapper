package rebase

import (
	"github.com/ldez/go-git-cmd-wrapper/types"
)

// Upstream Upstream branch to compare against. May be any valid commit, not just an existing branch name.
// Defaults to the configured upstream for the current branch.
func Upstream(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// Branch Working branch; defaults to HEAD.
func Branch(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}
