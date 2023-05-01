package controllers

import (
	"github.com/BSteffaniak/lazyaws/pkg/commands"
	"github.com/BSteffaniak/lazyaws/pkg/commands/oscommands"
	"github.com/BSteffaniak/lazyaws/pkg/gui/context"
	"github.com/BSteffaniak/lazyaws/pkg/gui/controllers/helpers"
	"github.com/BSteffaniak/lazyaws/pkg/gui/types"
)

type controllerCommon struct {
	c        *types.HelperCommon
	os       *oscommands.OSCommand
	git      *commands.GitCommand
	helpers  *helpers.Helpers
	model    *types.Model
	contexts *context.ContextTree
	modes    *types.Modes
	mutexes  *types.Mutexes
}

func NewControllerCommon(
	c *types.HelperCommon,
	os *oscommands.OSCommand,
	git *commands.GitCommand,
	helpers *helpers.Helpers,
	model *types.Model,
	contexts *context.ContextTree,
	modes *types.Modes,
	mutexes *types.Mutexes,
) *controllerCommon {
	return &controllerCommon{
		c:        c,
		os:       os,
		git:      git,
		helpers:  helpers,
		model:    model,
		contexts: contexts,
		modes:    modes,
		mutexes:  mutexes,
	}
}
