package entity

type Tasks struct {
	View           string `json:"view"`
	Fields         string `json:"fields"`
	ExcludedFields string `json:"excludedFields"`
	Ids            string `json:"ids"`
	ProjectId      string `json:"projectId"`
	OutdatedTasks  string `json:"outdatedTasks"`
}

type TasksClaim []struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Properties Properties `json:"properties"`
}

type TasksComplete struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Properties Properties `json:"properties"`
}
type TransitionID struct {
	Type string `json:"type"`
}
type Comment struct {
	Type string `json:"type"`
}
type ID struct {
	Type string `json:"type"`
}
type Properties struct {
	TransitionID TransitionID `json:"transitionId"`
	Comment      Comment      `json:"comment"`
	ID           ID           `json:"id"`
}
