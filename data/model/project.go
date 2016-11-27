package model

const (
	buildPostfix = "_Build"
)

type Project struct {
	ID      string
	BuildID string
}

func NewProject(ID string) Project {
	return Project{
		ID:      ID,
		BuildID: ID + buildPostfix,
	}
}
