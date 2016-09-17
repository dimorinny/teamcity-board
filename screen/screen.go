package screen

import (
	ui "github.com/gizak/termui"
)

type Screen interface {
	Content() []*ui.Row
	StartHandlers()
}
