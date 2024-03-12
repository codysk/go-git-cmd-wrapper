package commit

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/codysk/go-git-cmd-wrapper/v2/types"
)

// All Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.
// -a, --all
func All(g *types.Cmd) {
	g.AddOptions("--all")
}

// AllowEmpty Usually recording a commit that has the exact same tree as its sole parent commit is a mistake, and the command prevents you from making such a commit.
// This option bypasses the safety, and is primarily for use by foreign SCM interface scripts.
// --allow-empty
func AllowEmpty(g *types.Cmd) {
	g.AddOptions("--allow-empty")
}

// AllowEmptyMessage Like --allow-empty this command is primarily for use by foreign SCM interface scripts.
// It allows you to create a commit with an empty commit message without using plumbing commands like git-commit-tree(1).
// --allow-empty-message
func AllowEmptyMessage(g *types.Cmd) {
	g.AddOptions("--allow-empty-message")
}

// Amend Replace the tip of the current branch by creating a new commit.
// The recorded tree is prepared as usual (including the effect of the -i and -o options and explicit pathspec), and the message from the original commit is used as the starting point, instead of an empty message, when no other message is specified from the command line via options such as -m, -F, -c, etc.
// The new commit has the same parents and author as the current one (the --reset-author option can countermand this).
// --amend
func Amend(g *types.Cmd) {
	g.AddOptions("--amend")
}

// Author Override the commit author.
// Specify an explicit author using the standard A U Thor <author@example.com> format.
// Otherwise <author> is assumed to be a pattern and is used to search for an existing commit by that author (i.e. rev-list --all -i --author=<author>); the commit author is then copied from the first such commit found.
// --author=<author>
func Author(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--author=%s", value))
	}
}

// Branch Show the branch and tracking info even in short-format.
// --branch
func Branch(g *types.Cmd) {
	g.AddOptions("--branch")
}

// Cleanup This option determines how the supplied commit message should be cleaned up before committing.
// The <mode> can be strip, whitespace, verbatim, scissors or default.
// --cleanup=<mode>
func Cleanup(mode string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--cleanup=%s", mode))
	}
}

// Date Override the author date used in the commit.
// --date=<date>
func Date(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--date=%s", value))
	}
}

// DryRun Do not create a commit, but show a list of paths that are to be committed, paths with local changes that will be left uncommitted and paths that are untracked.
// --dry-run
func DryRun(g *types.Cmd) {
	g.AddOptions("--dry-run")
}

// Edit The message taken from file with -F, command line with -m, and from commit object with -C are usually used as the commit log message unmodified.
// This option lets you further edit the message taken from these sources.
// -e, --edit
func Edit(g *types.Cmd) {
	g.AddOptions("--edit")
}

// File Take the commit message from the given file.
// Use - to read the message from the standard input.
// -F <file>, --file=<file>
func File(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--file=%s", value))
	}
}

// Fixup Construct a commit message for use with rebase --autosquash.
// The commit message will be the subject line from the specified commit with a prefix of 'fixup! '.
// See git-rebase(1) for details.
// --fixup=<commit>
func Fixup(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--fixup=%s", commit))
	}
}

// GpgSign GPG-sign commits.
// The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space.
// -S[<keyid>], --gpg-sign[=<keyid>]
func GpgSign(keyID string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(keyID) == 0 {
			g.AddOptions("--gpg-sign")
		} else {
			g.AddOptions(fmt.Sprintf("--gpg-sign=%s", keyID))
		}
	}
}

// Include Before making a commit out of staged contents so far, stage the contents of paths given on the command line as well.
// This is usually not what you want unless you are concluding a conflicted merge.
// -i, --include
func Include(g *types.Cmd) {
	g.AddOptions("--include")
}

// Long When doing a dry-run, give the output in the long-format.
// Implies --dry-run.
// --long
func Long(g *types.Cmd) {
	g.AddOptions("--long")
}

// NoEdit Use the selected commit message without launching an editor.
// For example, git commit --amend --no-edit amends a commit without changing its commit message.
// --no-edit
func NoEdit(g *types.Cmd) {
	g.AddOptions("--no-edit")
}

// NoGpgSign Countermand commit.gpgSign configuration variable that is set to force each and every commit to be signed.
// --no-gpg-sign
func NoGpgSign(g *types.Cmd) {
	g.AddOptions("--no-gpg-sign")
}

// NoPostRewrite Bypass the post-rewrite hook.
// --no-post-rewrite
func NoPostRewrite(g *types.Cmd) {
	g.AddOptions("--no-post-rewrite")
}

// NoStatus Do not include the output of git-status(1) in the commit message template when using an editor to prepare the default commit message.
// --no-status
func NoStatus(g *types.Cmd) {
	g.AddOptions("--no-status")
}

// NoVerify This option bypasses the pre-commit and commit-msg hooks.
// See also githooks(5).
// -n, --no-verify
func NoVerify(g *types.Cmd) {
	g.AddOptions("--no-verify")
}

// Null When showing short or porcelain status output, print the filename verbatim and terminate the entries with NUL, instead of LF.
// If no format is given, implies the --porcelain output format.
// Without the -z option, filenames with 'unusual' characters are quoted as explained for the configuration variable core.quotePath (see git-config(1)).
// -z, --null
func Null(g *types.Cmd) {
	g.AddOptions("--null")
}

// Only Make a commit by taking the updated working tree contents of the paths specified on the command line, disregarding any contents that have been staged for other paths.
// This is the default mode of operation of git commit if any paths are given on the command line, in which case this option can be omitted.
// If this option is specified together with --amend, then no paths need to be specified, which can be used to amend the last commit without committing changes that have already been staged.
// If used together with --allow-empty paths are also not required, and an empty commit will be created.
// -o, --only
func Only(g *types.Cmd) {
	g.AddOptions("--only")
}

// Patch Use the interactive patch selection interface to chose which changes to commit.
// See git-add(1) for details.
// -p, --patch
func Patch(g *types.Cmd) {
	g.AddOptions("--patch")
}

// Porcelain When doing a dry-run, give the output in a porcelain-ready format.
// See git-status(1) for details.
// Implies --dry-run.
// --porcelain
func Porcelain(g *types.Cmd) {
	g.AddOptions("--porcelain")
}

// Quiet Suppress commit summary message.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// ReeditMessage Like -C, but with -c the editor is invoked, so that the user can further edit the commit message.
// -c <commit>, --reedit-message=<commit>
func ReeditMessage(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--reedit-message=%s", commit))
	}
}

// ResetAuthor When used with -C/-c/--amend options, or when committing after a conflicting cherry-pick, declare that the authorship of the resulting commit now belongs to the committer.
// This also renews the author timestamp.
// --reset-author
func ResetAuthor(g *types.Cmd) {
	g.AddOptions("--reset-author")
}

// ReuseMessage Take an existing commit object, and reuse the log message and the authorship information (including the timestamp) when creating the commit.
// -C <commit>, --reuse-message=<commit>
func ReuseMessage(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--reuse-message=%s", commit))
	}
}

// Short When doing a dry-run, give the output in the short-format.
// See git-status(1) for details.
// Implies --dry-run.
// --short
func Short(g *types.Cmd) {
	g.AddOptions("--short")
}

// Signoff Add Signed-off-by line by the committer at the end of the commit log message.
// The meaning of a signoff depends on the project, but it typically certifies that committer has the rights to submit this work under the same license and agrees to a Developer Certificate of Origin (see https://developercertificate.org/ for more information).
// -s, --signoff
func Signoff(g *types.Cmd) {
	g.AddOptions("--signoff")
}

// Squash Construct a commit message for use with rebase --autosquash.
// The commit message subject line is taken from the specified commit with a prefix of 'squash! '.
// Can be used with additional commit message options (-m/-c/-C/-F).
// See git-rebase(1) for details.
// --squash=<commit>
func Squash(commit string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--squash=%s", commit))
	}
}

// Status Include the output of git-status(1) in the commit message template when using an editor to prepare the commit message.
// Defaults to on, but can be used to override configuration variable commit.status.
// --status
func Status(g *types.Cmd) {
	g.AddOptions("--status")
}

// Template When editing the commit message, start the editor with the contents in the given file.
// The commit.template configuration variable is often used to give this option implicitly to the command.
// This mechanism can be used by projects that want to guide participants with some hints on what to write in the message in what order.
// If the user exits the editor without editing the message, the commit is aborted.
// This has no effect when a message is given by other means, e.g. with the -m or -F options.
// -t <file>, --template=<file>
func Template(file string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--template=%s", file))
	}
}

// UntrackedFiles Show untracked files.
// The mode parameter is optional (defaults to all), and is used to specify the handling of untracked files; when -u is not used, the default is normal, i.e. show untracked files and directories.
// -u[<mode>], --untracked-files[=<mode>]
func UntrackedFiles(mode string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(mode) == 0 {
			g.AddOptions("--untracked-files")
		} else {
			g.AddOptions(fmt.Sprintf("--untracked-files=%s", mode))
		}
	}
}

// Verbose Show unified diff between the HEAD commit and what would be committed at the bottom of the commit message template to help the user describe the commit by reminding what changes the commit has.
// Note that this diff output doesn’t have its lines prefixed with #.
// This diff will not be a part of the commit message.
// See the commit.verbose configuration variable in git-config(1).
// If specified twice, show in addition the unified diff between what would be committed and the worktree files, i.e. the unstaged changes to tracked files.
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}
