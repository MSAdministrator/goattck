package models

import (
	"fmt"
)

type Tactic interface {
	Techniques() ([]Technique, error)
}

type TacticObject struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	ObjectMarkingRefs  []string `json:"object_marking_refs"`
	CreatedByRef       string   `json:"created_by_ref"`
	ExternalReferences []ExternalReference `json:"external_references"`
	Description             string `json:"description"`
	XMitreAttackSpecVersion string `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string `json:"x_mitre_modified_by_ref"`
	XMitreShortname         string `json:"x_mitre_shortname"`
}

func NewTactic(object map[string]interface{}) (TacticObject, error) {
	tactic := TacticObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	tactic.BaseModel = baseModel
	if object["object_marking_refs"] != nil {
		tactic.ObjectMarkingRefs = ConvertInterfaceArrayToStringArray(object["object_marking_refs"].([]interface{}))
	}
	if object["created_by_ref"] != nil {
		tactic.CreatedByRef = object["created_by_ref"].(string)
	}
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		tactic.ExternalReferences = refs
	}
	if object["description"] != nil {
		tactic.Description = object["description"].(string)
	}
	if object["x_mitre_attack_spec_version"] != nil {
		tactic.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	if object["x_mitre_modified_by_ref"] != nil {
		tactic.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	if object["x_mitre_shortname"] != nil {
		tactic.XMitreShortname = object["x_mitre_shortname"].(string)
	}
	return tactic, nil
}

func (t *TacticObject) Techniques() ([]Technique, error) {
	return nil, nil
}
