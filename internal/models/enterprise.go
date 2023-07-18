package models

type EnterpriseAttck struct {
	Type        string        `json:"type"`
	ID          string        `json:"id"`
	Objects     []interface{} `json:"objects"`
	SpecVersion string        `json:"spec_version"`
}
