package presentation

import "github.com/BSteffaniak/lazyaws/pkg/gui/style"

func OpensMenuStyle(str string) string {
	return style.FgMagenta.Sprintf("%s...", str)
}
