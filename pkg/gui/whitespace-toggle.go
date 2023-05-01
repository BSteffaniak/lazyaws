package gui

import (
	"github.com/BSteffaniak/lazyaws/pkg/gui/types"
)

func (gui *Gui) toggleWhitespaceInDiffView() error {
	gui.IgnoreWhitespaceInDiffView = !gui.IgnoreWhitespaceInDiffView

	toastMessage := gui.c.Tr.ShowingWhitespaceInDiffView
	if gui.IgnoreWhitespaceInDiffView {
		toastMessage = gui.c.Tr.IgnoringWhitespaceInDiffView
	}
	gui.c.Toast(toastMessage)

	return gui.currentSideListContext().HandleFocus(types.OnFocusOpts{})
}
