package fetch

import (
	"github.com/codysk/go-git-cmd-wrapper/v2/types"
)

// Remote name.
func Remote(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// Group name.
func Group(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// RefSpec name.
func RefSpec(ref string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(ref)
	}
}
