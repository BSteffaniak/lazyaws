package types

import (
	"github.com/BSteffaniak/lazyaws/pkg/gui/modes/cherrypicking"
	"github.com/BSteffaniak/lazyaws/pkg/gui/modes/diffing"
	"github.com/BSteffaniak/lazyaws/pkg/gui/modes/filtering"
)

type Modes struct {
	Filtering     filtering.Filtering
	CherryPicking *cherrypicking.CherryPicking
	Diffing       diffing.Diffing
}
