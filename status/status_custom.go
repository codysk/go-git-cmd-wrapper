package status

import "github.com/codysk/go-git-cmd-wrapper/v2/types"

// PathSpec <pathspec>...
// See the pathspec entry in gitglossary(7).
func PathSpec(paths ...string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		for _, path := range paths {
			g.AddOptions(path)
		}
	}
}
