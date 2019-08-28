package entity

type Tasks struct {
	View           string `json:"view"`
	Fields         string `json:"fields"`
	ExcludedFields string `json:"excludedFields"`
	Ids            string `json:"ids"`
	ProjectId      string `json:"projectId"`
	OutdatedTasks  string `json:"outdatedTasks"`
}
