package models

import (
	"encoding/json"
)

type dataComponent interface {
}

type DataComponent struct {
	BaseModel
	BaseAttributes
	// These are properties from the MITRE ATT&CK json
	XMitreDataSourceRef     string `json:"x_mitre_data_source_ref"`
	Type                    string `json:"type"`
	XMitreAttackSpecVersion string `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string `json:"x_mitre_modified_by_ref"`
	Techniques              []Technique
}

var _ (dataComponent) = new(DataComponent)

func NewDataComponent(object map[string]interface{}) (DataComponent, error) {
	dataComponent := DataComponent{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &dataComponent)
	return dataComponent, nil
}

func (dc *DataComponent) GetExternalID() string {
	return ""
}
