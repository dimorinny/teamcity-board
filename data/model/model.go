package model

type BaseResponse struct {
	Count int `json:"count"`
	Href string `json:"href"`
}

type AgentsResponse struct {
	BaseResponse
	Agents []Agent `json:"agent"`
}

type BuildsResponse struct {
	BaseResponse
	Builds []Build `json:"build"`
}

type Agent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	TypeId int `json:"typeId"`
	Href string `json:"href"`
}

type Build struct {
	Id int `json:"id"`
	BuildTypeId string `json:"buildTypeId"`
	Number string `json:"number"`
	Status string `json:"status"`
	State string `json:"state"`
	Percentage *int  `json:"percentageComplete,omitempty"`
	BranchName string `json:"branchName"`
	Href string `json:"href"`
	Url string `json:"webUrl"`
}