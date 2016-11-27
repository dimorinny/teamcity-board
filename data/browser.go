package data

import (
	"fmt"
	"github.com/dimorinny/teamcity-board/config"
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

func (browser *TeamcityBrowser) OpenBoard(projectID string) {
	open.Run(
		fmt.Sprintf(
			"%s:%d/project.html?projectId=%s",
			browser.configuration.Host,
			browser.configuration.Port,
			projectID,
		),
	)
}

func (browser *TeamcityBrowser) OpenBuild(buildTypeID string, id int) {
	open.Run(
		fmt.Sprintf(
			"%s:%d/viewLog.html?buildId=%d&buildTypeId=%s",
			browser.configuration.Host,
			browser.configuration.Port,
			id,
			buildTypeID,
		),
	)
}
