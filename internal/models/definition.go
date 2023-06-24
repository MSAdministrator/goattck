package models

type MarkingDefinition struct {
	// Base fields
	BaseModel
	// Fields
	Type                    string               `json:"type"`
	Id                      string               `json:"id"`
	Created                 string               `json:"created"`
	Definition              string               `json:"definition"`
	DefinitionType          string               `json:"definition_type"`
	CreatedByRef            string               `json:"created_by_ref"`
	XMitreAttackSpecVersion string               `json:"x_mitre_attack_spec_version"`
	ObjectMarkingRefs       []string             `json:"object_marking_refs"`
	ExternalReferences      []ExternalReferences `json:"external_references"`
	Revoked                 bool                 `json:"revoked"`
}
