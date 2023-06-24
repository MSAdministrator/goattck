package models

type Campaign struct {
	// Base fields
	BaseModel
	// Fields
	Type                    string               `json:"type"`
	Description             string               `json:"description"`
	FirstSeen               string               `json:"first_seen"`
	LastSeen                string               `json:"last_seen"`
	XMitreFirstSeenCitation string               `json:"x_mitre_first_seen_citation"`
	XMitreLastSeenCitation  string               `json:"x_mitre_last_seen_citation"`
	Aliases                 []string             `json:"aliases"`
	XMitreDeprecated        bool                 `json:"x_mitre_deprecated"`
	XMitreContributors      []string             `json:"x_mitre_contributors"`
	Revoked                 bool                 `json:"revoked"`
	CreatedByRef            string               `json:"created_by_ref"`
	ExternalReferences      []ExternalReferences `json:"external_references"`
	ObjectMarkingRefs       []string             `json:"object_marking_refs"`
	XMitreAttackSpecVersion string               `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string               `json:"x_mitre_modified_by_ref"`
	XMitreDomains           []string             `json:"x_mitre_domains"`
}

func (c *Campaign) Malwares() ([]Malware, error) {
	return nil, nil
}

func (c *Campaign) Tools() ([]Tool, error) {
	return nil, nil
}

func (c *Campaign) Techniques() ([]Technique, error) {
	return nil, nil
}
