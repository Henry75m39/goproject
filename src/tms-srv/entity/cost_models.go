package entity

type CostModels struct {
	ExcludedFields string `json:"excludedFields"`
	Fields         string `json:"fields"`
	SortBy         string `json:"sortBy"`
	SortDirection  string `json:"sortDirection"`
	ScopingId      string `json:"scopingId"`
	LocaleId       string `json:"localeId"`
}
