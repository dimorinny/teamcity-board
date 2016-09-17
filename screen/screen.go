package screen

import (
	ui "github.com/gizak/termui"
)

const (
	boardHeight = 2
)

type Screen interface {
	Content() []*ui.Row
	StartHandlers()
}
