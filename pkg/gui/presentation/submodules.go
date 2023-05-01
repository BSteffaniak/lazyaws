package presentation

import (
	"github.com/jesseduffield/generics/slices"
	"github.com/BSteffaniak/lazyaws/pkg/commands/models"
	"github.com/BSteffaniak/lazyaws/pkg/theme"
)

func GetSubmoduleListDisplayStrings(submodules []*models.SubmoduleConfig) [][]string {
	return slices.Map(submodules, func(submodule *models.SubmoduleConfig) []string {
		return getSubmoduleDisplayStrings(submodule)
	})
}

func getSubmoduleDisplayStrings(s *models.SubmoduleConfig) []string {
	return []string{theme.DefaultTextColor.Sprint(s.Name)}
}
