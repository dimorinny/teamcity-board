package model

type BaseResponse struct {
	Count int    `json:"count"`
	Href  string `json:"href"`
}

type AgentsResponse struct {
	BaseResponse
	Agents []Agent `json:"agent"`
}

type BuildsResponse struct {
	BaseResponse
	Builds []Build `json:"build"`
}

type QueueResponse struct {
	BaseResponse
	Queue []QueueItem `json:"build"`
}
