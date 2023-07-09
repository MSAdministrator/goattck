package models

type Mitigation struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	ObjectMarkingRefs  []string `json:"object_marking_refs"`
	CreatedByRef       string   `json:"created_by_ref"`
	ExternalReferences []struct {
		SourceName  string `json:"source_name"`
		URL         string `json:"url"`
		ExternalID  string `json:"external_id,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"external_references"`
	Description             string `json:"description"`
	XMitreDeprecated        bool   `json:"x_mitre_deprecated,omitempty"`
	XMitreModifiedByRef     string `json:"x_mitre_modified_by_ref"`
	Revoked                 bool   `json:"revoked,omitempty"`
	XMitreAttackSpecVersion string `json:"x_mitre_attack_spec_version,omitempty"`
}

func (m *Mitigation) Techniques() ([]Technique, error) {
	return nil, nil
}
