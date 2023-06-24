package models

type Actor struct {
	Type                    string               `json:"type"`
	ID                      string               `json:"id"`
	Aliases                 []string             `json:"aliases"`
	XMitreContributors      []string             `json:"x_mitre_contributors"`
	Revoked                 bool                 `json:"revoked"`
	Description             string               `json:"description"`
	XMitreModifiedByRef     string               `json:"x_mitre_modified_by_ref"`
	XMitreDeprecated        bool                 `json:"x_mitre_deprecated"`
	XMitreAttackSpecVersion string               `json:"x_mitre_attack_spec_version"`
	CreatedByRef            string               `json:"created_by_ref"`
	XMitreDomains           []string             `json:"x_mitre_domains"`
	ObjectMarkingRefs       []string             `json:"object_marking_refs"`
	ExternalReferences      []ExternalReferences `json:"external_references"`
	// These are properties unique to pyattck-data
	Names               []string `json:"names"`
	ExternalTools       []string `json:"external_tools"`
	Country             []string `json:"country"`
	Operations          []string `json:"operations"`
	Links               []string `json:"links"`
	Targets             []string `json:"targets"`
	ExternalDescription []string `json:"external_description"`
	AttckID             string   `json:"attck_id"`
	Comment             string   `json:"comment"`
}

func (a *Actor) Malwares() ([]Malware, error) {
	return nil, nil
}

func (a *Actor) Tools() ([]Tool, error) {
	return nil, nil
}

func (a *Actor) Techniques() ([]Technique, error) {
	return nil, nil
}
