package models

import (
	"fmt"
)

type DataComponent interface {
	Techniques() ([]Technique, error)
}

type DataComponentObject struct {
	BaseModel
	BaseAttributes
	// These are properties from the MITRE ATT&CK json
	XMitreDataSourceRef     string `json:"x_mitre_data_source_ref"`
	Type                    string `json:"type"`
	XMitreAttackSpecVersion string `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string `json:"x_mitre_modified_by_ref"`
}

func NewDataComponent(object map[string]interface{}) (DataComponentObject, error) {
	dataComponent := DataComponentObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	dataComponent.BaseModel = baseModel
	baseAttributes, err := parseBaseAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base attributes: %s", err))
	}
	dataComponent.BaseAttributes = baseAttributes
	if object["x_mitre_data_source_ref"] != nil {
		dataComponent.XMitreDataSourceRef = object["x_mitre_data_source_ref"].(string)
	}
	if object["type"] != nil {
		dataComponent.Type = object["type"].(string)
	}
	if object["x_mitre_attack_spec_version"] != nil {
		dataComponent.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	if object["x_mitre_modified_by_ref"] != nil {
		dataComponent.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	return dataComponent, nil
}

func (d *DataComponentObject) Techniques() ([]Technique, error) {
	return nil, nil
}
