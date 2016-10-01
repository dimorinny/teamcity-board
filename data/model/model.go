package model

type Agent struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	TypeID int    `json:"typeId"`
	Href   string `json:"href"`
}

type Build struct {
	ID          int    `json:"id"`
	BuildTypeID string `json:"buildTypeId"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	State       string `json:"state"`
	Percentage  int    `json:"percentageComplete,omitempty"`
	BranchName  string `json:"branchName"`
	Href        string `json:"href"`
	URL         string `json:"webUrl"`
}

type BuildType struct {
	Href        string `json:"href"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	ProjectID   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	URL         string `json:"webUrl"`
}

type Events struct {
	Events []Event `json:"change"`
}

type Event struct {
	Date     string `json:"date"`
	Href     string `json:"href"`
	ID       int    `json:"id"`
	Username string `json:"username"`
	Version  string `json:"version"`
	URL      string `json:"webUrl"`
}

type Triggered struct {
	Type    string `json:"type"`
	Details string `json:"details"`
	Date    string `json:"date"`
}

type DetailBuild struct {
	Build
	BuildType  BuildType `json:"buildType"`
	Events     Events    `json:"lastChanges"`
	StatusText string    `json:"statusText"`
	StartDate  string    `json:"startDate"`
	FinishDate string    `json:"finishDate"`
	Agent      Agent     `json:"agent"`
	Triggered  Triggered `json:"triggered"`
}

type QueueItem struct {
	Id          int    `json:"id"`
	BuildTypeId string `json:"buildTypeId"`
	State       string `json:"state"`
	BranchName  string `json:"branchName"`
	Href        string `json:"href"`
	Url         string `json:"webUrl"`
}
