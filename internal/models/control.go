package models

type Control struct {
	// Base fields
	BaseModel
	// Fields
	Revoked            bool                 `json:"revoked"`
	XMitreFamily       string               `json:"x_mitre_family"`
	XMitreImpact       []string             `json:"x_mitre_impact"`
	XMitrePriority     string               `json:"x_mitre_priority"`
	ObjectMarkingRefs  []string             `json:"object_marking_refs"`
	ExternalReferences []ExternalReferences `json:"external_references"`
}

func (c *Control) Techniques() ([]Technique, error) {
	return nil, nil
}
