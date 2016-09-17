package data

import (
	"github.com/dimorinny/teamcity-board/config"
	"github.com/dimorinny/teamcity-board/data/model"
	"net/http"
	"encoding/json"
	"fmt"
)

const (
	StateRunning = "running"
	StatusFail = "FAILURE"
)

type Teamcity struct {
	configuration config.HostConfig
	client http.Client
}

func NewTeamcity(configuration config.HostConfig) *Teamcity {
	return &Teamcity{
		configuration: configuration,
		client: http.Client{},
	}
}

func (teamcity *Teamcity) LoadAgents() ([]model.Agent, error) {
	agents := &model.AgentsResponse{}
	err := teamcity.load("GET", "agents", agents)
	if err != nil {
		return nil, err
	}
	return agents.Agents, nil
}

func (teamcity *Teamcity) LoadBuilds() ([]model.Build, error) {
	builds := &model.BuildsResponse{}
	// TODO: remove build type hardcode
	err := teamcity.load(
		"GET",
		"builds/?locator=buildType:AndroidProjects_AvitoPro_Build,branch:(default:any),state:any",
		builds,
	)
	if err != nil {
		return nil, err
	}
	return builds.Builds, nil
}

func (teamcity *Teamcity) load(method, url string, response interface{}) error {
	req, err := http.NewRequest(
		method,
		fmt.Sprintf(
			"%s:%d/httpAuth/app/rest/%s",
			teamcity.configuration.Host,
			teamcity.configuration.Port,
			url,
		),
		nil,
	)
	if err != nil {
		return err
	}

	req.Header.Add(
		"Authorization",
		teamcity.configuration.AuthHeader,
	)

	req.Header.Add(
		"Accept",
		"application/json",
	)

	resp, err := teamcity.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(response)
}

