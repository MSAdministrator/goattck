package models

import (
	"fmt"
)

type MarkingDefinition interface {
}

type MarkingDefinitionObject struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	Definition struct {
		Statement string `json:"statement"`
	} `json:"definition"`
	CreatedByRef            string `json:"created_by_ref"`
	DefinitionType          string `json:"definition_type"`
	XMitreAttackSpecVersion string `json:"x_mitre_attack_spec_version"`
}

func NewMarkingDefinition(object map[string]interface{}) (MarkingDefinitionObject, error) {
	markingDefinition := MarkingDefinitionObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	markingDefinition.BaseModel = baseModel
	if object["definition"] != nil {
		definition := object["definition"].(map[string]interface{})
		if definition["statement"] != nil {
			markingDefinition.Definition.Statement = definition["statement"].(string)
		}
	}
	if object["created_by_ref"] != nil {
		markingDefinition.CreatedByRef = object["created_by_ref"].(string)
	}
	if object["definition_type"] != nil {
		markingDefinition.DefinitionType = object["definition_type"].(string)
	}
	if object["x_mitre_attack_spec_version"] != nil {
		markingDefinition.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	return markingDefinition, nil
}
