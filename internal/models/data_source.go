package models

type DataSource struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	Description            string   `json:"description"`
	XMitrePlatforms        []string `json:"x_mitre_platforms"`
	XMitreDeprecated       bool     `json:"x_mitre_deprecated,omitempty"`
	XMitreContributors     []string `json:"x_mitre_contributors,omitempty"`
	XMitreCollectionLayers []string `json:"x_mitre_collection_layers"`
	CreatedByRef           string   `json:"created_by_ref"`
	Revoked                bool     `json:"revoked,omitempty"`
	ExternalReferences     []struct {
		SourceName string `json:"source_name"`
		URL        string `json:"url"`
		ExternalID string `json:"external_id"`
	} `json:"external_references"`
	ObjectMarkingRefs       []string `json:"object_marking_refs"`
	XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref"`
}

func (d *DataSource) DataComponents() ([]DataComponent, error) {
	return nil, nil
}

func (d *DataSource) Techniques() ([]Technique, error) {
	return nil, nil
}
