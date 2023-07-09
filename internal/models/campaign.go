package models

import "time"

type Campaign struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	Description             string    `json:"description"`
	Aliases                 []string  `json:"aliases"`
	FirstSeen               time.Time `json:"first_seen"`
	LastSeen                time.Time `json:"last_seen"`
	XMitreFirstSeenCitation string    `json:"x_mitre_first_seen_citation"`
	XMitreLastSeenCitation  string    `json:"x_mitre_last_seen_citation"`
	XMitreDeprecated        bool      `json:"x_mitre_deprecated"`
	CreatedByRef            string    `json:"created_by_ref"`
	Revoked                 bool      `json:"revoked"`
	ExternalReferences      []struct {
		SourceName  string `json:"source_name"`
		URL         string `json:"url"`
		ExternalID  string `json:"external_id,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"external_references"`
	ObjectMarkingRefs       []string `json:"object_marking_refs"`
	XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref"`
	XMitreContributors      []string `json:"x_mitre_contributors,omitempty"`
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
