package screen

import (
	"fmt"
	"github.com/dimorinny/teamcity-board/data"
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
				12,
				0,
				buildScreen.StatusBlock(),
			),
		),
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
				buildScreen.ProgressBlock(),
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

func (buildScreen *BuildScreen) StatusBlock() *ui.Par {
	icon := buildScreen.GetBuildIcon()

	if buildScreen.build.Status == data.BuildStatusFailure {
		icon = crossIcon
	} else if buildScreen.build.Status == data.BuildStatusSuccess {
		icon = checkmarkIcon
	} else {
		icon = ""
	}

	par := ui.NewPar(fmt.Sprintf("STATUS: %s %s", buildScreen.build.Status, icon))
	par.Height = 3
	par.TextFgColor = buildScreen.GetBuildColor()

	return par
}

func (buildScreen *BuildScreen) ProgressBlock() *ui.Gauge {
	g := ui.NewGauge()
	g.Percent = buildScreen.build.Percentage
	g.Float = ui.AlignLeft

	if buildScreen.build.State == data.StateRunning {
		g.Height = 3
	} else {
		g.Height = 0
	}

	if buildScreen.build.Status == data.StatusFail {
		g.BarColor = ui.ColorRed
	} else {
		g.BarColor = ui.ColorGreen
	}

	return g
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
		err                   error
		startDate, finishDate time.Time
	)

	startDate, err = time.Parse(teamcityDateFormat, buildScreen.build.StartDate)

	if buildScreen.build.State == data.StateRunning {
		finishDate = time.Now()
	} else {
		finishDate, err = time.Parse(teamcityDateFormat, buildScreen.build.FinishDate)
	}

	if err != nil {
		log.Fatal(err)
	}

	messages := []string{
		buildScreen.build.StatusText,
		"Time: " + fmt.Sprintf(
			"%s - %s",
			startDate,
			finishDate,
		),
		"Branch: " + buildScreen.build.BranchName,
	}

	par := ui.NewPar(strings.Join(messages, "\n"))
	par.Height = len(messages) + view.BoardHeight

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

	return par
}

func (buildScreen *BuildScreen) GetBuildColor() ui.Attribute {
	var color ui.Attribute

	if buildScreen.build.Status == data.BuildStatusFailure {
		color = ui.ColorRed
	} else if buildScreen.build.Status == data.BuildStatusSuccess {
		color = ui.ColorGreen
	} else {
		color = ui.ColorWhite
	}

	return color
}

func (buildScreen *BuildScreen) GetBuildIcon() string {
	var icon string

	if buildScreen.build.Status == data.BuildStatusFailure {
		icon = crossIcon
	} else if buildScreen.build.Status == data.BuildStatusSuccess {
		icon = checkmarkIcon
	} else {
		icon = ""
	}

	return icon
}

func (buildScreen *BuildScreen) StartHandlers() {
	buildScreen.context.AddKeyHandler("o", func(e ui.Event) {
		buildScreen.context.Browser.OpenBuild(
			"AndroidProjects_AvitoPro_Build",
			buildScreen.buildId,
		)
	})
}
