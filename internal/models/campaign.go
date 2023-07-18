package models

import (
	"fmt"
	"time"
)

type Campaign interface {
	Malwares() ([]Malware, error)
	Tools() ([]Tool, error)
	Techniques() ([]Technique, error)
}

type CampaignObject struct {
	BaseModel
	BaseAttributes
	BaseExternalModel
	// These are properties from the MITRE ATT&CK json
	FirstSeen               time.Time `json:"first_seen"`
	LastSeen                time.Time `json:"last_seen"`
	XMitreFirstSeenCitation string    `json:"x_mitre_first_seen_citation"`
	XMitreLastSeenCitation  string    `json:"x_mitre_last_seen_citation"`
	XMitreContributors      []string  `json:"x_mitre_contributors,omitempty"`
}

func NewCampaign(object map[string]interface{}) (CampaignObject, error) {
	campaign := CampaignObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	campaign.BaseModel = baseModel
	baseAttributes, err := parseBaseAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base attributes: %s", err))
	}
	campaign.BaseAttributes = baseAttributes
	baseExternalModel, err := parseExternalModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing external model: %s", err))
	}
	campaign.BaseExternalModel = baseExternalModel
	if object["first_seen"] != nil {
		firstSeen, err := time.Parse(time.RFC3339, object["first_seen"].(string))
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing first_seen: %s", err))
		}
		campaign.FirstSeen = firstSeen
	}
	if object["last_seen"] != nil {
		lastSeen, err := time.Parse(time.RFC3339, object["last_seen"].(string))
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing last_seen: %s", err))
		}
		campaign.LastSeen = lastSeen
	}
	if object["x_mitre_first_seen_citation"] != nil {
		campaign.XMitreFirstSeenCitation = object["x_mitre_first_seen_citation"].(string)
	}
	if object["x_mitre_last_seen_citation"] != nil {
		campaign.XMitreLastSeenCitation = object["x_mitre_last_seen_citation"].(string)
	}
	if object["x_mitre_contributors"] != nil {
		campaign.XMitreContributors = ConvertInterfaceArrayToStringArray(object["x_mitre_contributors"].([]interface{}))
	}
	return campaign, nil
}


func (c *CampaignObject) Malwares() ([]Malware, error) {
	return nil, nil
}

func (c *CampaignObject) Tools() ([]Tool, error) {
	return nil, nil
}

func (c *CampaignObject) Techniques() ([]Technique, error) {
	return nil, nil
}
