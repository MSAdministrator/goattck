package models

import (
	"encoding/json"
	"strings"
	"time"
)

type campaign interface {
}

type Campaign struct {
	BaseModel
	BaseAttributes
	BaseExternalModel
	// These are properties from the MITRE ATT&CK json
	FirstSeen               time.Time `json:"first_seen"`
	LastSeen                time.Time `json:"last_seen"`
	XMitreFirstSeenCitation string    `json:"x_mitre_first_seen_citation"`
	XMitreLastSeenCitation  string    `json:"x_mitre_last_seen_citation"`
	XMitreContributors      []string  `json:"x_mitre_contributors,omitempty"`
	Malwares                []Malware
	Techniques              []Technique
	Tools                   []Tool
}

var _ (campaign) = new(Campaign)

func NewCampaign(object map[string]interface{}) (Campaign, error) {
	campaign := Campaign{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &campaign)
	return campaign, nil
}

func (m *Campaign) GetExternalID() string {
	mitreID := ""
	for _, ref := range m.ExternalReferences {
		if ref.ExternalId != "" && mitreID == "" {
			if strings.HasPrefix(ref.ExternalId, "C") {
				mitreID = ref.ExternalId
			}
		}
	}
	return mitreID
}
