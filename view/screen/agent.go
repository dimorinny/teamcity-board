package screen

import (
	"github.com/dimorinny/teamcity-board/data/model"
	"github.com/dimorinny/teamcity-board/view"
	"github.com/dimorinny/teamcity-board/view/widget"
	ui "github.com/gizak/termui"
)

type AgentScreen struct {
	context *view.Context
	agent   model.Agent
}

func NewAgentScreen(context *view.Context) view.Screen {
	return &AgentScreen{
		context: context,
	}
}

func (agentScreen *AgentScreen) Content() []*ui.Row {
	return []*ui.Row{
		ui.NewRow(
			ui.NewCol(12, 0, widget.GetInfoView()),
		),
	}
}

func (agentScreen *AgentScreen) StartHandlers() {
	agentScreen.context.AddHandler("/sys/kbd/q", func(event ui.Event) {
		agentScreen.context.Exit()
	})

	//for index := range agentScreen.agents {
	//	agentScreen.AddHandler("/sys/kbd/" + strconv.Itoa(index), func(event ui.Event) {
	//
	//	})
	//}
}
