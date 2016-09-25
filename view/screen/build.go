package screen

import (
	"github.com/dimorinny/teamcity-board/data/model"
	"github.com/dimorinny/teamcity-board/view"
	"github.com/dimorinny/teamcity-board/view/widget"
	ui "github.com/gizak/termui"
	"strconv"
	"strings"
	"time"
	"log"
	"fmt"
)

type BuildScreen struct {
	log     widget.LogView
	context *view.Context

	build   model.DetailBuild
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
	buildScreen.LoadBuild()
	buildScreen.log.AddMessage(fmt.Sprintf("Build with id %d loaded", buildScreen.buildId))

	return []*ui.Row{
		ui.NewRow(
			ui.NewCol(
				6,
				0,
				buildScreen.GetLeftInfoBlock(),
			),
			ui.NewCol(
				12,
				0,
				buildScreen.log.GenerateView(),
			),
		),
	}
}

func (buildScreen *BuildScreen) LoadBuild() {
	build, err := buildScreen.context.Client.LoadBuild(buildScreen.buildId)
	if err != nil {
		log.Fatal(err)
	}

	buildScreen.build = *build
}

func (buildScreen *BuildScreen) GetLeftInfoBlock() *ui.Par {
	startTime, _ := strconv.ParseInt(buildScreen.build.StartDate, 10, 64)
	endTime, _ := strconv.ParseInt(buildScreen.build.FinishDate, 10, 64)

	messages := []string{
		"Result: " + buildScreen.build.StatusText,
		"Time: " + fmt.Sprintf(
			"%s - %s",
			time.Unix(startTime, 0).Format(dateFormat),
			time.Unix(endTime, 0).Format(dateFormat),
		),
		"Branch: " + buildScreen.build.BranchName,
	}

	par := ui.NewPar(strings.Join(messages, "\n"))
	par.Height = len(messages) + view.BoardHeight
	return par
}

func (buildScreen *BuildScreen) StartHandlers() {
	//for index := range agentScreen.agents {
	//	agentScreen.AddHandler("/sys/kbd/" + strconv.Itoa(index), func(event ui.Event) {
	//
	//	})
	//}
}
