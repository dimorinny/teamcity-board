package screen

import (
	"fmt"
	"github.com/dimorinny/teamcity-board/data"
	"github.com/dimorinny/teamcity-board/data/model"
	"github.com/dimorinny/teamcity-board/view"
	"github.com/dimorinny/teamcity-board/view/widget"
	ui "github.com/gizak/termui"
	"log"
)

type BoardScreen struct {
	log     widget.LogView
	context *view.Context
	project model.Project

	agents []model.Agent
	builds []model.Build
	queue  []model.QueueItem
}

func NewBoardScreen(context *view.Context, project model.Project) view.Screen {
	return &BoardScreen{
		context: context,
		project: project,
		log:     widget.NewLogView(logMessagesCount),
	}
}

func (boardScreen *BoardScreen) Content() []*ui.Row {
	boardScreen.loadAgents()
	boardScreen.log.AddMessage("Agents loaded")
	boardScreen.loadBuilds()
	boardScreen.log.AddMessage("Builds loaded")
	boardScreen.loadQueue()
	boardScreen.log.AddMessage("Queue loaded")

	return []*ui.Row{
		ui.NewRow(
			ui.NewCol(
				6,
				0,
				boardScreen.getBuildList(),
			),
			ui.NewCol(
				6,
				0,
				boardScreen.getBuildProgresses()...,
			),
		),
		ui.NewRow(
			ui.NewCol(
				4,
				0,
				widget.ListOrEmpty(
					boardScreen.getAgentList(),
					buildQueueTitle,
					emptyItemHeight,
				),
			),
			ui.NewCol(
				4,
				0,
				widget.ListOrEmpty(
					boardScreen.getQueueList(),
					buildQueueTitle,
					emptyItemHeight,
				),
			),
			ui.NewCol(
				4,
				0,
				widget.GetInfoView(),
			),
		),
		ui.NewRow(
			ui.NewCol(
				12,
				0,
				boardScreen.log.GenerateView(),
			),
		),
	}
}

func (boardScreen *BoardScreen) loadAgents() {
	agents, err := boardScreen.context.Client.LoadAgents()
	if err != nil {
		log.Fatal(err)
	}

	boardScreen.agents = agents
}

func (boardScreen *BoardScreen) loadBuilds() {
	builds, err := boardScreen.context.Client.LoadBuilds(
		boardScreen.project.BuildID,
		maximumBuildsCount,
	)
	if err != nil {
		log.Fatal(err)
	}

	boardScreen.builds = builds
}

func (boardScreen *BoardScreen) loadQueue() {
	queue, err := boardScreen.context.Client.LoadQueue()
	if err != nil {
		log.Fatal(err)
	}

	boardScreen.queue = queue
}

func (boardScreen *BoardScreen) getAgentList() *ui.List {
	ls := ui.NewList()
	ls.Border = true
	ls.BorderLabel = agentsTitle
	for _, agent := range boardScreen.agents {
		ls.Items = append(ls.Items, agent.Name)
	}
	ls.Height = len(boardScreen.agents) + view.BoardHeight

	return ls
}

func (boardScreen *BoardScreen) getBuildList() *ui.List {
	length := len(boardScreen.builds)

	builds := ui.NewList()
	for index, build := range boardScreen.builds {
		buildTitle := fmt.Sprintf(
			"[%d] %s %s",
			index,
			build.BranchName,
			build.Status,
		)
		builds.Items = append(builds.Items, buildTitle)
	}
	builds.Border = false
	builds.Height = length

	return builds
}

func (boardScreen *BoardScreen) getBuildProgresses() []ui.GridBufferer {
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

func (boardScreen *BoardScreen) getQueueList() *ui.List {
	ls := ui.NewList()
	ls.Border = true
	ls.BorderLabel = buildQueueTitle
	for _, queue := range boardScreen.queue {
		ls.Items = append(ls.Items, queue.BranchName)
	}

	length := len(boardScreen.queue)

	if length <= 5 {
		ls.ItemFgColor = ui.ColorGreen
	} else if length <= 10 {
		ls.ItemFgColor = ui.ColorYellow
	} else {
		ls.ItemFgColor = ui.ColorRed
	}

	ls.Height = length + view.BoardHeight

	return ls
}

func (boardScreen *BoardScreen) StartHandlers() {
	boardScreen.context.AddKeyHandler("o", func(e ui.Event) {
		boardScreen.context.Browser.OpenBoard(boardScreen.project.ID)
	})
	boardScreen.context.AddNumberHandler(func(key int) {
		if len(boardScreen.builds) > key {
			boardScreen.context.StartScreen(
				NewBuildScreen(
					boardScreen.context,
					boardScreen.project,
					boardScreen.builds[key].ID,
				),
				true,
			)
		}
	})
}
