package presentation

import (
	"github.com/jesseduffield/generics/slices"
	"github.com/BSteffaniak/lazyaws/pkg/commands/models"
	"github.com/BSteffaniak/lazyaws/pkg/gui/presentation/icons"
	"github.com/BSteffaniak/lazyaws/pkg/gui/style"
	"github.com/BSteffaniak/lazyaws/pkg/theme"
)

func GetRemoteListDisplayStrings(remotes []*models.Remote, diffName string) [][]string {
	return slices.Map(remotes, func(remote *models.Remote) []string {
		diffed := remote.Name == diffName
		return getRemoteDisplayStrings(remote, diffed)
	})
}

// getRemoteDisplayStrings returns the display string of branch
func getRemoteDisplayStrings(r *models.Remote, diffed bool) []string {
	branchCount := len(r.Branches)

	textStyle := theme.DefaultTextColor
	if diffed {
		textStyle = theme.DiffTerminalColor
	}

	res := make([]string, 0, 3)
	if icons.IsIconEnabled() {
		res = append(res, textStyle.Sprint(icons.IconForRemote(r)))
	}
	res = append(res, textStyle.Sprint(r.Name), style.FgBlue.Sprintf("%d branches", branchCount))
	return res
}
