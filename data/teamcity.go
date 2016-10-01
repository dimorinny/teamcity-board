package data

import (
	"encoding/json"
	"fmt"
	"github.com/dimorinny/teamcity-board/config"
	"github.com/dimorinny/teamcity-board/data/model"
	"net/http"
)

const (
	StateRunning = "running"
	StatusFail   = "FAILURE"

	BuildStatusFailure = "FAILURE"
	BuildStatusSuccess = "SUCCESS"
)

type Teamcity struct {
	configuration config.HostConfig
	client        http.Client
}

func NewTeamcity(configuration config.HostConfig) *Teamcity {
	return &Teamcity{
		configuration: configuration,
		client:        http.Client{},
	}
}

func (teamcity *Teamcity) LoadBuild(id int) (*model.DetailBuild, error) {
	build := &model.DetailBuild{}
	err := teamcity.load(
		"GET",
		fmt.Sprintf("builds/id:%d", id),
		build,
	)
	if err != nil {
		return nil, err
	}
	return build, nil
}

func (teamcity *Teamcity) LoadAgents() ([]model.Agent, error) {
	agents := &model.AgentsResponse{}
	err := teamcity.load(
		"GET",
		"agents",
		agents,
	)
	if err != nil {
		return nil, err
	}
	return agents.Agents, nil
}

func (teamcity *Teamcity) LoadBuilds(buildType string, count int) ([]model.Build, error) {
	builds := &model.BuildsResponse{}
	err := teamcity.load(
		"GET",
		fmt.Sprintf(
			"builds/?locator=buildType:%s,branch:(default:any),state:any,count:%d",
			buildType,
			count,
		),
		builds,
	)
	if err != nil {
		return nil, err
	}
	return builds.Builds, nil
}

func (teamcity *Teamcity) LoadQueue() ([]model.QueueItem, error) {
	queue := &model.QueueResponse{}
	err := teamcity.load(
		"GET",
		"buildQueue",
		queue,
	)
	if err != nil {
		return nil, err
	}
	return queue.Queue, nil
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
