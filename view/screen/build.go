package screen

import (
	"fmt"
	"github.com/dimorinny/teamcity-board/data/model"
	"github.com/dimorinny/teamcity-board/view"
	"github.com/dimorinny/teamcity-board/view/widget"
	ui "github.com/gizak/termui"
	"log"
	"strings"
	"time"
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
				6,
				0,
				buildScreen.GetRightInfoBlock(),
			),
		),
		ui.NewRow(
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
	var (
		icon                  string
		err                   error
		startDate, finishDate time.Time
	)

	if buildScreen.build.Status == model.BuildStatusFailure {
		icon = crossIcon
	} else if buildScreen.build.Status == model.BuildStatusSuccess {
		icon = checkmarkIcon
	} else {
		icon = ""
	}

	startDate, err = time.Parse(teamcityDateFormat, buildScreen.build.StartDate)
	finishDate, err = time.Parse(teamcityDateFormat, buildScreen.build.FinishDate)
	if err != nil {
		log.Fatal(err)
	}

	messages := []string{
		fmt.Sprintf("Result: %s %s", buildScreen.build.StatusText, icon),
		"Time: " + fmt.Sprintf(
			"%s - %s",
			startDate,
			finishDate,
		),
		"Branch: " + buildScreen.build.BranchName,
	}

	par := ui.NewPar(strings.Join(messages, "\n"))
	par.Height = len(messages) + view.BoardHeight
	par.TextFgColor = buildScreen.GetBuildColor()

	return par
}

func (buildScreen *BuildScreen) GetRightInfoBlock() *ui.Par {
	var (
		username string
		err      error
		date     time.Time
	)

	if len(buildScreen.build.Events.Events) != 0 {
		username = buildScreen.build.Events.Events[0].Username
	} else {
		username = ""
	}

	date, err = time.Parse(teamcityDateFormat, buildScreen.build.Triggered.Date)
	if err != nil {
		log.Fatal(err)
	}

	messages := []string{
		"Agent: " + buildScreen.build.Agent.Name,
		"Triggered: " + fmt.Sprintf(
			"%s %s",
			username,
			date,
		),
	}

	par := ui.NewPar(strings.Join(messages, "\n"))
	par.Height = 3 + view.BoardHeight
	par.TextFgColor = buildScreen.GetBuildColor()

	return par
}

func (buildScreen *BuildScreen) GetBuildColor() ui.Attribute {
	var color ui.Attribute

	if buildScreen.build.Status == model.BuildStatusFailure {
		color = ui.ColorRed
	} else if buildScreen.build.Status == model.BuildStatusSuccess {
		color = ui.ColorGreen
	} else {
		color = ui.ColorWhite
	}

	return color
}

func (buildScreen *BuildScreen) StartHandlers() {}
