package entity

import "time"

type ProjectGroup []struct {
	PersistToken  string       `json:"persistToken"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	ProjectTypeID int          `json:"projectTypeId"`
	ClientID      int          `json:"clientId"`
	SystemFiles   []string     `json:"systemFiles"`
	AisFiles      []string     `json:"aisFiles"`
	Locales       []Locales    `json:"locales"`
	Attributes    []Attributes `json:"attributes"`
}

type Locales struct {
	Id      int       `json:"id"`
	DueDate time.Time `json:"dueDate"`
}

type Attributes struct {
	Attribute Attribute `json:"attribute"`
	Value     string    `json:"value"`
}

type Attribute struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
