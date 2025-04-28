package models

import (
	"encoding/json"
	"strings"
)

type actor interface {
}

// ActorObject is a representation of the MITRE ATT&CK Actor json model
type Actor struct {
	BaseModel
	BaseAttributes
	BaseExternalModel
	// These are properties from the MITRE ATT&CK json
	XMitreContributors []string `json:"x_mitre_contributors,omitempty"`
	// These are properties unique to pyattck-data
	actorExternalAttributes
	MitreAttckId string `json:"mitre_attck_id"`
	Malwares     []Malware
	Techniques   []Technique
	Tools        []Tool
}

var _ (actor) = new(Actor)

// actorExternalAttributes are properties external from the MITRE ATT&CK json definitions
type actorExternalAttributes struct {
	Names               []string `json:"names"`
	ExternalTools       []string `json:"external_tools"`
	Country             []string `json:"country"`
	Operations          []string `json:"operations"`
	Links               []string `json:"links"`
	Targets             []string `json:"targets"`
	ExternalDescription []string `json:"external_description"`
	AttckID             string   `json:"attck_id"`
	Comment             string   `json:"comment"`
	Comments            []string `json:"comments"`
}

// NewActor is a function that takes in a map of data and returns a ActorObject
func NewActor(object map[string]interface{}) (Actor, error) {
	actor := Actor{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &actor)
	return actor, nil
}

func (a *Actor) GetExternalID() string {
	mitreID := ""
	for _, ref := range a.ExternalReferences {
		if ref.ExternalId != "" && mitreID == "" {
			if strings.HasPrefix(ref.ExternalId, "G") {
				mitreID = ref.ExternalId
			}
		}
	}
	return mitreID
}
