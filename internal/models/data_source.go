package models

import "encoding/json"

type dataSource interface {
}

type DataSource struct {
	BaseModel
	BaseAttributes
	BaseExternalModel
	// These are properties from the MITRE ATT&CK json
	XMitrePlatforms         []string            `json:"x_mitre_platforms"`
	XMitreContributors      []string            `json:"x_mitre_contributors,omitempty"`
	XMitreCollectionLayers  []string            `json:"x_mitre_collection_layers"`
	ExternalReferences      []ExternalReference `json:"external_references"`
	XMitreAttackSpecVersion string              `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string              `json:"x_mitre_modified_by_ref"`
	DataComponents          []DataComponent
	Techniques              []Technique
}

var _ (dataSource) = new(DataSource)

func NewDataSource(object map[string]interface{}) (DataSource, error) {
	dataSource := DataSource{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &dataSource)
	return dataSource, nil
}
