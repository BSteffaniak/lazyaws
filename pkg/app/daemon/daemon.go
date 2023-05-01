package daemon

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BSteffaniak/lazyaws/pkg/common"
	"github.com/BSteffaniak/lazyaws/pkg/env"
)

// Sometimes lazyaws will be invoked in daemon mode from a parent lazyaws process.
// We do this when git lets us supply a program to run within a git command.
// For example, if we want to ensure that a git command doesn't hang due to
// waiting for an editor to save a commit message, we can tell git to invoke lazyaws
// as the editor via 'GIT_EDITOR=lazyaws', and use the env var
// 'LAZYAWS_DAEMON_KIND=EXIT_IMMEDIATELY' to specify that we want to run lazyaws
// as a daemon which simply exits immediately. Any additional arguments we want
// to pass to a daemon can be done via other env vars.

type DaemonKind string

const (
	InteractiveRebase DaemonKind = "INTERACTIVE_REBASE"
	ExitImmediately   DaemonKind = "EXIT_IMMEDIATELY"
)

const (
	DaemonKindEnvKey string = "LAZYAWS_DAEMON_KIND"
	RebaseTODOEnvKey string = "LAZYAWS_REBASE_TODO"

	// The `PrependLinesEnvKey` env variable is set to `true` to tell our daemon
	// to prepend the content of `RebaseTODOEnvKey` to the default `git-rebase-todo`
	// file instead of using it as a replacement.
	PrependLinesEnvKey string = "LAZYAWS_PREPEND_LINES"
)

type Daemon interface {
	Run() error
}

func Handle(common *common.Common) {
	d := getDaemon(common)
	if d == nil {
		return
	}

	if err := d.Run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func InDaemonMode() bool {
	return getDaemonKind() != ""
}

func getDaemon(common *common.Common) Daemon {
	switch getDaemonKind() {
	case InteractiveRebase:
		return &rebaseDaemon{c: common}
	case ExitImmediately:
		return &exitImmediatelyDaemon{c: common}
	}

	return nil
}

func getDaemonKind() DaemonKind {
	return DaemonKind(os.Getenv(DaemonKindEnvKey))
}

type rebaseDaemon struct {
	c *common.Common
}

func (self *rebaseDaemon) Run() error {
	self.c.Log.Info("Lazyaws invoked as interactive rebase demon")
	self.c.Log.Info("args: ", os.Args)
	path := os.Args[1]

	if strings.HasSuffix(path, "git-rebase-todo") {
		return self.writeTodoFile(path)
	} else if strings.HasSuffix(path, filepath.Join(gitDir(), "COMMIT_EDITMSG")) { // TODO: test
		// if we are rebasing and squashing, we'll see a COMMIT_EDITMSG
		// but in this case we don't need to edit it, so we'll just return
	} else {
		self.c.Log.Info("Lazyaws demon did not match on any use cases")
	}

	return nil
}

func (self *rebaseDaemon) writeTodoFile(path string) error {
	todoContent := []byte(os.Getenv(RebaseTODOEnvKey))

	prependLines := os.Getenv(PrependLinesEnvKey) != ""
	if prependLines {
		existingContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		todoContent = append(todoContent, existingContent...)
	}

	return os.WriteFile(path, todoContent, 0o644)
}

func gitDir() string {
	dir := env.GetGitDirEnv()
	if dir == "" {
		return ".git"
	}
	return dir
}

type exitImmediatelyDaemon struct {
	c *common.Common
}

func (self *exitImmediatelyDaemon) Run() error {
	return nil
}
