package main

import (
	"github.com/dimorinny/teamcity-board/config"
	"github.com/dimorinny/teamcity-board/data"
	"github.com/dimorinny/teamcity-board/screen"
	"os"
)

var (
	configuration config.Config
)

func initConfig() {
	configuration = config.Config{
		Interval: 2,
		Host: config.HostConfig{
			Host:       "http://teamcity",
			Port:       8111,
			AuthHeader: os.Getenv("AUTH_HEADER"),
		},
	}
}

func init() {
	initConfig()
}

func main() {
	context := screen.NewContext(
		configuration.Interval,
		data.NewTeamcity(configuration.Host),
	)
	context.Init()
	defer context.Close()

	context.StartScreen(screen.NewBoardScreen(context))

	context.StartLoop()
}
