package models

import (
	"fmt"
)

type Mitigation interface {
	Techniques() ([]Technique, error)
}

type MitigationObject struct {
	BaseModel
	BaseAttributes
	// These are properties from the MITRE ATT&CK json
	ExternalReferences []ExternalReference `json:"external_references"`
	XMitreModifiedByRef     string `json:"x_mitre_modified_by_ref"`
	XMitreAttackSpecVersion string `json:"x_mitre_attack_spec_version,omitempty"`
}

func NewMitigation(object map[string]interface{}) (MitigationObject, error) {
	mitigation := MitigationObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	mitigation.BaseModel = baseModel
	baseAttributes, err := parseBaseAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base attributes: %s", err))
	}
	mitigation.BaseAttributes = baseAttributes
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		mitigation.ExternalReferences = refs
	}
	if object["x_mitre_modified_by_ref"] != nil {
		mitigation.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	if object["x_mitre_attack_spec_version"] != nil {
		mitigation.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	return mitigation, nil
}

func (m *MitigationObject) Techniques() ([]Technique, error) {
	return nil, nil
}
