package gui

import (
	"github.com/BSteffaniak/lazyaws/pkg/commands/git_commands"
	"github.com/BSteffaniak/lazyaws/pkg/commands/oscommands"
	"github.com/BSteffaniak/lazyaws/pkg/config"
	"github.com/BSteffaniak/lazyaws/pkg/updates"
	"github.com/BSteffaniak/lazyaws/pkg/utils"
)

// NewDummyGui creates a new dummy GUI for testing
func NewDummyUpdater() *updates.Updater {
	newAppConfig := config.NewDummyAppConfig()
	dummyUpdater, _ := updates.NewUpdater(utils.NewDummyCommon(), newAppConfig, oscommands.NewDummyOSCommand())
	return dummyUpdater
}

func NewDummyGui() *Gui {
	newAppConfig := config.NewDummyAppConfig()
	dummyGui, _ := NewGui(utils.NewDummyCommon(), newAppConfig, &git_commands.GitVersion{}, NewDummyUpdater(), false, "")
	return dummyGui
}
