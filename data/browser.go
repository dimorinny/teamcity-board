package data

import (
	"github.com/dimorinny/teamcity-board/config"
	"fmt"
	"github.com/skratchdot/open-golang/open"
)

type TeamcityBrowser struct {
	configuration config.HostConfig
}

func NewTeamcityBrowser(configuration config.HostConfig) *TeamcityBrowser {
	return &TeamcityBrowser{
		configuration: configuration,
	}
}

func (browser *TeamcityBrowser) OpenBoard(buildType string) {
	open.Run(
		fmt.Sprintf(
			"%s:%d/project.html?projectId=%s",
			browser.configuration.Host,
			browser.configuration.Port,
			buildType,
		),
	)
}

func (browser *TeamcityBrowser) OpenBuild(id int) {

}
