package models

import (
	"fmt"
)

type DataSource interface {
	DataComponents() ([]DataComponent, error)
	Techniques() ([]Technique, error)
}

type DataSourceObject struct {
	BaseModel
	BaseAttributes
	BaseExternalModel
	// These are properties from the MITRE ATT&CK json
	XMitrePlatforms        []string `json:"x_mitre_platforms"`
	XMitreContributors     []string `json:"x_mitre_contributors,omitempty"`
	XMitreCollectionLayers []string `json:"x_mitre_collection_layers"`
	ExternalReferences     []ExternalReference `json:"external_references"`
	XMitreAttackSpecVersion string `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string `json:"x_mitre_modified_by_ref"`
}

func NewDataSource(object map[string]interface{}) (DataSourceObject, error) {
	dataSource := DataSourceObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	dataSource.BaseModel = baseModel
	baseAttributes, err := parseBaseAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base attributes: %s", err))
	}
	dataSource.BaseAttributes = baseAttributes
	baseExternalModel, err := parseExternalModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing external model: %s", err))
	}
	dataSource.BaseExternalModel = baseExternalModel
	if object["x_mitre_platforms"] != nil {
		dataSource.XMitrePlatforms = ConvertInterfaceArrayToStringArray(object["x_mitre_platforms"].([]interface{}))
	}
	if object["x_mitre_contributors"] != nil {
		dataSource.XMitreContributors = ConvertInterfaceArrayToStringArray(object["x_mitre_contributors"].([]interface{}))
	}
	if object["x_mitre_collection_layers"] != nil {
		dataSource.XMitreCollectionLayers = ConvertInterfaceArrayToStringArray(object["x_mitre_collection_layers"].([]interface{}))
	}
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		dataSource.ExternalReferences = refs
	}
	if object["x_mitre_attack_spec_version"] != nil {
		dataSource.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	if object["x_mitre_modified_by_ref"] != nil {
		dataSource.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	return dataSource, nil
}

func (d *DataSourceObject) DataComponents() ([]DataComponent, error) {
	return nil, nil
}

func (d *DataSourceObject) Techniques() ([]Technique, error) {
	return nil, nil
}
