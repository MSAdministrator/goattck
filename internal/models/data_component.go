package models

type DataComponent struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	Description             string   `json:"description"`
	XMitreDataSourceRef     string   `json:"x_mitre_data_source_ref"`
	XMitreDeprecated        bool     `json:"x_mitre_deprecated,omitempty"`
	Type                    string   `json:"type"`
	CreatedByRef            string   `json:"created_by_ref"`
	Revoked                 bool     `json:"revoked,omitempty"`
	ObjectMarkingRefs       []string `json:"object_marking_refs"`
	XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref"`
}

func (d *DataComponent) Techniques() ([]Technique, error) {
	return nil, nil
}
