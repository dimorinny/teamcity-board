package screen

import (
	"github.com/dimorinny/teamcity-board/data/model"
	ui "github.com/gizak/termui"
)

type AgentScreen struct {
	context *Context
	agent   model.Agent
}

func NewAgentScreen(context *Context) Screen {
	return AgentScreen{
		context: context,
	}
}

func (agentScreen AgentScreen) Content() []*ui.Row {
	return []*ui.Row{}
	//ui.NewRow(
	//	ui.NewCol(4, 0, agentScreen.getAgentsList()),
	//	ui.NewCol(4, 0, agentScreen.getAgentsList()),
	//	ui.NewCol(4, 0, agentScreen.getAgentsList()),
	//),
}

func (agentScreen AgentScreen) StartHandlers() {
	agentScreen.context.AddHandler("/sys/kbd/q", func(event ui.Event) {
		agentScreen.context.Exit()
	})

	//for index := range agentScreen.agents {
	//	agentScreen.AddHandler("/sys/kbd/" + strconv.Itoa(index), func(event ui.Event) {
	//
	//	})
	//}
}
