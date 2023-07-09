package models

type MarkingDefinition struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	Definition struct {
		Statement string `json:"statement"`
	} `json:"definition"`
	CreatedByRef            string `json:"created_by_ref"`
	DefinitionType          string `json:"definition_type"`
	XMitreAttackSpecVersion string `json:"x_mitre_attack_spec_version"`
}
