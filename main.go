package main

import (
	"github.com/caarlos0/env"
	"github.com/dimorinny/teamcity-board/config"
	"github.com/dimorinny/teamcity-board/data"
	"github.com/dimorinny/teamcity-board/data/model"
	"github.com/dimorinny/teamcity-board/view"
	"github.com/dimorinny/teamcity-board/view/screen"
	"log"
)

var (
	configuration config.Config
)

func initConfig() {
	configuration = config.Config{}

	err := env.Parse(&configuration)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	initConfig()
}

func main() {
	context := view.NewContext(
		configuration.Interval,
		data.NewTeamcity(configuration),
		data.NewTeamcityBrowser(configuration),
	)
	project := model.NewProject(
		configuration.ProjectID,
	)
	context.Init()
	defer context.Close()

	context.StartScreen(
		screen.NewBoardScreen(
			context,
			project,
		),
		true,
	)

	context.StartLoop()
}
