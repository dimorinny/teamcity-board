package widget

import (
	"github.com/dimorinny/teamcity-board/view"
	ui "github.com/gizak/termui"
	"strings"
)

const (
	keyMap = "Keymap"
)

func GetInfoView() *ui.Par {
	messages := []string{
		"r - Reload",
		"b - Back",
		"<build-number> - open build info",
	}

	par := ui.NewPar(strings.Join(messages, "\n"))
	par.Height = len(messages) + view.BoardHeight
	par.BorderLabel = keyMap
	return par
}
