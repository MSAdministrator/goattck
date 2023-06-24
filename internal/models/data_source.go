package models

type DataSource struct {
	// Base fields
	BaseModel
	// Fields
	Type                    string               `json:"type"`
	Description             string               `json:"description"`
	XMitreModifiedByRef     string               `json:"x_mitre_modified_by_ref"`
	XMitreAttackSpecVersion string               `json:"x_mitre_attack_spec_version"`
	XMitreCollectionLayers  []string             `json:"x_mitre_collection_layers"`
	XMitreDomains           []string             `json:"x_mitre_domains"`
	CreatedByRef            string               `json:"created_by_ref"`
	ExternalReferences      []ExternalReferences `json:"external_references"`
	ObjectMarkingRefs       []string             `json:"object_marking_refs"`
	Aliases                 []string             `json:"aliases"`
	Revoked                 bool                 `json:"revoked"`
	XMitreDeprecated        bool                 `json:"x_mitre_deprecated"`
	XMitreContributors      []string             `json:"x_mitre_contributors"`
	XMitrePlatforms         []string             `json:"x_mitre_platforms"`
}

func (d *DataSource) DataComponents() ([]DataComponent, error) {
	return nil, nil
}

func (d *DataSource) Techniques() ([]Technique, error) {
	return nil, nil
}
