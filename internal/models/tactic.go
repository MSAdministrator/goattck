package models

import "encoding/json"

type tactic interface{}

type Tactic struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	ObjectMarkingRefs       []string            `json:"object_marking_refs"`
	CreatedByRef            string              `json:"created_by_ref"`
	ExternalReferences      []ExternalReference `json:"external_references"`
	Description             string              `json:"description"`
	XMitreAttackSpecVersion string              `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string              `json:"x_mitre_modified_by_ref"`
	XMitreShortname         string              `json:"x_mitre_shortname"`
	Techniques              []Technique
}

var _ (tactic) = new(Tactic)

func NewTactic(object map[string]interface{}) (Tactic, error) {
	tactic := Tactic{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &tactic)
	return tactic, nil
}
