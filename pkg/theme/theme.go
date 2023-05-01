package theme

import (
	"github.com/jesseduffield/gocui"
	"github.com/BSteffaniak/lazyaws/pkg/config"
	"github.com/BSteffaniak/lazyaws/pkg/gui/style"
)

var (
	// GocuiDefaultTextColor does the same as DefaultTextColor but this one only colors gocui default text colors
	GocuiDefaultTextColor gocui.Attribute

	// ActiveBorderColor is the border color of the active frame
	ActiveBorderColor gocui.Attribute

	// InactiveBorderColor is the border color of the inactive active frames
	InactiveBorderColor gocui.Attribute

	// GocuiSelectedLineBgColor is the background color for the selected line in gocui
	GocuiSelectedLineBgColor gocui.Attribute

	OptionsColor gocui.Attribute

	// DefaultTextColor is the default text color
	DefaultTextColor = style.FgWhite

	// SelectedLineBgColor is the background color for the selected line
	SelectedLineBgColor = style.New()

	// SelectedRangeBgColor is the background color of the selected range of lines
	SelectedRangeBgColor = style.New()

	// CherryPickedCommitColor is the text style when cherry picking a commit
	CherryPickedCommitTextStyle = style.New()

	OptionsFgColor = style.New()

	DiffTerminalColor = style.FgMagenta

	UnstagedChangesColor = style.New()
)

// UpdateTheme updates all theme variables
func UpdateTheme(themeConfig config.ThemeConfig) {
	ActiveBorderColor = GetGocuiStyle(themeConfig.ActiveBorderColor)
	InactiveBorderColor = GetGocuiStyle(themeConfig.InactiveBorderColor)
	SelectedLineBgColor = GetTextStyle(themeConfig.SelectedLineBgColor, true)
	SelectedRangeBgColor = GetTextStyle(themeConfig.SelectedRangeBgColor, true)

	cherryPickedCommitBgTextStyle := GetTextStyle(themeConfig.CherryPickedCommitBgColor, true)
	cherryPickedCommitFgTextStyle := GetTextStyle(themeConfig.CherryPickedCommitFgColor, false)
	CherryPickedCommitTextStyle = cherryPickedCommitBgTextStyle.MergeStyle(cherryPickedCommitFgTextStyle)

	unstagedChangesTextStyle := GetTextStyle(themeConfig.UnstagedChangesColor, false)
	UnstagedChangesColor = unstagedChangesTextStyle

	GocuiSelectedLineBgColor = GetGocuiStyle(themeConfig.SelectedLineBgColor)
	OptionsColor = GetGocuiStyle(themeConfig.OptionsTextColor)
	OptionsFgColor = GetTextStyle(themeConfig.OptionsTextColor, false)

	DefaultTextColor = GetTextStyle(themeConfig.DefaultFgColor, false)
	GocuiDefaultTextColor = GetGocuiStyle(themeConfig.DefaultFgColor)
}
