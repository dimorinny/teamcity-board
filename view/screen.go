package view

import (
	ui "github.com/gizak/termui"
)

const (
	BoardHeight = 2
)

type Screen interface {
	Content() []*ui.Row
	StartHandlers()
}
