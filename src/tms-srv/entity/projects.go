package entity

type Projects []struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WorkGroupID int    `json:"workgroupId"`
}