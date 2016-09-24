package widget

import (
	"github.com/dimorinny/teamcity-board/view"
	ui "github.com/gizak/termui"
)

const (
	emptyText = "There's nothing"
)

func GetEmptyView(title string, height int) *ui.Par {
	par := ui.NewPar(emptyText)
	par.Height = height + view.BoardHeight
	par.BorderLabel = title
	return par
}

func ListOrEmpty(list *ui.List, title string, height int) ui.GridBufferer {
	if len(list.Items) != 0 {
		return list
	} else {
		return GetEmptyView(title, height)
	}
}
