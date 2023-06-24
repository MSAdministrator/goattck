package models

type Mitigation struct {
	// Base fields
	BaseModel
	// Fields
	Type                string               `json:"type"`
	Description         string               `json:"description"`
	CreatedByRef        string               `json:"created_by_ref"`
	XMitreDeprecated    bool                 `json:"x_mitre_deprecated"`
	ObjectMarkingRefs   []string             `json:"object_marking_refs"`
	ExternalReferences  []ExternalReferences `json:"external_references"`
	MitigationId        string               `json:"mitigation_id"`
	XMitreModifiedByRef string               `json:"x_mitre_modified_by_ref"`
	XMitreDomains       []string             `json:"x_mitre_domains"`
	Revoked             bool                 `json:"revoked"`
}

func (m *Mitigation) Techniques() ([]Technique, error) {
	return nil, nil
}
