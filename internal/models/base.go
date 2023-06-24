package models

type BaseModel struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Created       string   `json:"created"`
	Modified      string   `json:"modified"`
	XMitreVersion string   `json:"x_mitre_version"`
	XMitreDomains []string `json:"x_mitre_domains"`
}

type ExternalReferences struct {
	SourceName  string `json:"source_name"`
	Url         string `json:"url"`
	ExternalId  string `json:"external_id"`
	Description string `json:"description"`
}
