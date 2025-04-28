package models

import "encoding/json"

type mitigation interface{}

type Mitigation struct {
	BaseModel
	BaseAttributes
	// These are properties from the MITRE ATT&CK json
	ExternalReferences      []ExternalReference `json:"external_references"`
	XMitreModifiedByRef     string              `json:"x_mitre_modified_by_ref"`
	XMitreAttackSpecVersion string              `json:"x_mitre_attack_spec_version,omitempty"`
	Techniques              []Technique
}

var _ (mitigation) = new(Mitigation)

func NewMitigation(object map[string]interface{}) (Mitigation, error) {
	mitigation := Mitigation{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &mitigation)
	return mitigation, nil
}
