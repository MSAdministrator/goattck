package models

type DataComponent struct {
	// Base fields
	BaseModel
	// Fields
	Type                    string               `json:"type"`
	Description             string               `json:"description"`
	CreatedByRef            string               `json:"created_by_ref"`
	XMitreModifiedByRef     string               `json:"x_mitre_modified_by_ref"`
	XMitreDataSourceRef     string               `json:"x_mitre_data_source_ref"`
	ObjectMarkingRefs       []string             `json:"object_marking_refs"`
	XMitreDomains           []string             `json:"x_mitre_domains"`
	XMitreAttackSpecVersion string               `json:"x_mitre_attack_spec_version"`
	XMitreDeprecated        bool                 `json:"x_mitre_deprecated"`
	Revoked                 bool                 `json:"revoked"`
	ExternalReferences      []ExternalReferences `json:"external_references"`
}

func (d *DataComponent) Techniques() ([]Technique, error) {
	return nil, nil
}
