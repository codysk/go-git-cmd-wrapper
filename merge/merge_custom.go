package merge

import "github.com/codysk/go-git-cmd-wrapper/v2/types"

// Commits <commit>...
// Commits, usually other branch heads, to merge into our branch.
// Specifying more than one commit will create a merge with more than two parents (affectionately called an Octopus merge).
func Commits(values ...string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		for _, value := range values {
			g.AddOptions(value)
		}
	}
}
