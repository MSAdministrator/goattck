package models

import "encoding/json"

type control interface {
}

type Control struct {
	// Base fields
	BaseModel
	// Fields
	Revoked            bool                `json:"revoked"`
	XMitreFamily       string              `json:"x_mitre_family"`
	XMitreImpact       []string            `json:"x_mitre_impact"`
	XMitrePriority     string              `json:"x_mitre_priority"`
	ObjectMarkingRefs  []string            `json:"object_marking_refs"`
	ExternalReferences []ExternalReference `json:"external_references"`
	Techniques         []Technique
}

var _ (control) = new(Control)

func NewControl(object map[string]interface{}) (Control, error) {
	control := Control{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &control)
	return control, nil
}
