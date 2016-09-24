package screen

import (
	"github.com/dimorinny/teamcity-board/view"
	"github.com/dimorinny/teamcity-board/view/widget"
	ui "github.com/gizak/termui"
)

type BuildScreen struct {
	log     widget.LogView
	context *view.Context

	buildId int
}

func NewBuildScreen(context *view.Context, buildId int) view.Screen {
	return &BuildScreen{
		context: context,
		buildId: buildId,
		log:     widget.NewLogView(logMessagesCount),
	}
}

func (buildScreen *BuildScreen) Content() []*ui.Row {
	//boardScreen.loadAgents()
	//boardScreen.log.AddMessage("Agents loaded")
	//boardScreen.loadBuilds()
	//boardScreen.log.AddMessage("Builds loaded")
	//boardScreen.loadQueue()
	//boardScreen.log.AddMessage("Queue loaded")

	return []*ui.Row{
		ui.NewRow(
			ui.NewCol(
				12,
				0,
				buildScreen.log.GenerateView(),
			),
		),
	}
}

func (buildScreen *BuildScreen) StartHandlers() {
	//for index := range agentScreen.agents {
	//	agentScreen.AddHandler("/sys/kbd/" + strconv.Itoa(index), func(event ui.Event) {
	//
	//	})
	//}
}
