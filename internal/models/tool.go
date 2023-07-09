package models

type Tool struct {
	Description        string   `json:"description"`
	XMitrePlatforms    []string `json:"x_mitre_platforms,omitempty"`
	XMitreDeprecated   bool     `json:"x_mitre_deprecated,omitempty"`
	XMitreContributors []string `json:"x_mitre_contributors,omitempty"`
	XMitreAliases      []string `json:"x_mitre_aliases,omitempty"`
	CreatedByRef       string   `json:"created_by_ref"`
	Revoked            bool     `json:"revoked,omitempty"`
	ExternalReferences []struct {
		SourceName  string `json:"source_name"`
		URL         string `json:"url"`
		ExternalID  string `json:"external_id,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"external_references"`
	ObjectMarkingRefs       []string `json:"object_marking_refs"`
	Labels                  []string `json:"labels"`
	XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref"`
}
