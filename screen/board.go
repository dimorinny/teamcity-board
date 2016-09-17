package screen

import (
	"bytes"
	"github.com/dimorinny/teamcity-board/data"
	"github.com/dimorinny/teamcity-board/data/model"
	ui "github.com/gizak/termui"
	"log"
	"strconv"
)

const (
	agentsTitle     = "Agents"
	buildQueueTitle = "Queue"
)

type BoardScreen struct {
	context *Context
	agents  []model.Agent
	builds  []model.Build
}

func NewBoardScreen(context *Context) Screen {
	return BoardScreen{
		context: context,
	}
}

func (boardScreen BoardScreen) Content() []*ui.Row {
	boardScreen.loadAgents()
	boardScreen.loadBuilds()

	return []*ui.Row{
		ui.NewRow(
			ui.NewCol(3, 0, boardScreen.getBuildList()),
			ui.NewCol(2, 0, boardScreen.getBuildProgresses()...),
		),
		ui.NewRow(
			ui.NewCol(2, 0, boardScreen.getAgentList()),
			ui.NewCol(2, 0, boardScreen.getBuildQueue()),
		),
	}
}

func (boardScreen *BoardScreen) loadAgents() {
	agents, err := boardScreen.context.client.LoadAgents()
	if err != nil {
		log.Fatal(err)
	}

	boardScreen.agents = agents
}

func (boardScreen *BoardScreen) loadBuilds() {
	builds, err := boardScreen.context.client.LoadBuilds(
		"AndroidProjects_AvitoPro_Build",
		10,
	)
	if err != nil {
		log.Fatal(err)
	}

	boardScreen.builds = builds
}

func (boardScreen BoardScreen) getAgentList() *ui.List {
	ls := ui.NewList()
	ls.Border = true
	ls.BorderLabel = agentsTitle
	for index, agent := range boardScreen.agents {
		buffer := bytes.NewBufferString("[")
		buffer.WriteString(strconv.Itoa(index))
		buffer.WriteString("] ")
		buffer.WriteString(agent.Name)

		ls.Items = append(ls.Items, buffer.String())
	}
	ls.ItemFgColor = ui.ColorYellow
	ls.Height = len(boardScreen.agents) + boardHeight

	return ls
}

func (boardScreen BoardScreen) getBuildList() *ui.List {
	length := len(boardScreen.builds)

	builds := ui.NewList()
	for index, build := range boardScreen.builds {
		buffer := bytes.NewBufferString("[")
		buffer.WriteString(strconv.Itoa(index))
		buffer.WriteString("] ")
		buffer.WriteString(build.BranchName)
		buffer.WriteString(" ")
		buffer.WriteString(build.Status)

		builds.Items = append(builds.Items, buffer.String())
	}
	builds.Border = false
	builds.PaddingLeft = 1
	builds.Height = length

	return builds
}

func (boardScreen BoardScreen) getBuildProgresses() []ui.GridBufferer {
	bars := []ui.GridBufferer{}

	for _, build := range boardScreen.builds {
		if build.State == data.StateRunning {
			g := ui.NewGauge()
			g.Border = false
			g.Percent = build.Percentage
			g.Float = ui.AlignLeft
			g.Height = 1

			if build.Status == data.StatusFail {
				g.BarColor = ui.ColorRed
			} else {
				g.BarColor = ui.ColorGreen
			}
			bars = append(bars, ui.GridBufferer(g))
		}
	}

	return bars
}

func (boardScreen BoardScreen) getBuildQueue() *ui.List {
	ls := ui.NewList()
	ls.Border = true
	ls.BorderLabel = buildQueueTitle
	for index, agent := range boardScreen.agents {
		buffer := bytes.NewBufferString("[")
		buffer.WriteString(strconv.Itoa(index))
		buffer.WriteString("] ")
		buffer.WriteString(agent.Name)

		ls.Items = append(ls.Items, buffer.String())
	}

	length := len(boardScreen.agents)

	if length <= 5 {
		ls.ItemFgColor = ui.ColorGreen
	} else if length <= 10 {
		ls.ItemFgColor = ui.ColorYellow
	} else {
		ls.ItemFgColor = ui.ColorRed
	}

	ls.Height = length + boardHeight

	return ls
}

func (boardScreen BoardScreen) StartHandlers() {
	boardScreen.context.AddKeyHandler("w", func(event ui.Event) {
		boardScreen.context.StartScreen(NewAgentScreen(boardScreen.context))
	})

	//for index := range agentScreen.agents {
	//	agentScreen.AddHandler("/sys/kbd/" + strconv.Itoa(index), func(event ui.Event) {
	//
	//	})
	//}
}
