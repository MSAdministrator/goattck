package models

import (
	"fmt"
	"reflect"
)

type Actor interface {
	Malwares() ([]Malware, error)
	Tools() ([]Tool, error)
	Techniques() ([]Technique, error)
}

type ActorObject struct {
	BaseModel
	BaseAttributes
	BaseExternalModel
	// These are properties from the MITRE ATT&CK json
	XMitreContributors []string `json:"x_mitre_contributors,omitempty"`
	// These are properties unique to pyattck-data
	actorExternalAttributes
	MitreAttckId string `json:"mitre_attck_id"`
}

type actorExternalAttributes struct {
	Names               []string `json:"names"`
	ExternalTools       []string `json:"external_tools"`
	Country             []string `json:"country"`
	Operations          []string `json:"operations"`
	Links               []string `json:"links"`
	Targets             []string `json:"targets"`
	ExternalDescription []string `json:"external_description"`
	AttckID				string   `json:"attck_id"`
	Comment             string   `json:"comment"`
	Comments			[]string `json:"comments"`
}

func NewActor(object map[string]interface{}) (ActorObject, error) {
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	baseAttributes, err := parseBaseAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base attributes: %s", err))
	}
	baseExternalModel, err := parseExternalModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing external model: %s", err))
	}
	aExternalAttributes, err := parseActorExternalAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing external attributes: %s", err))
	}
	actor := ActorObject{
		actorExternalAttributes: aExternalAttributes,
	}
	actor.BaseModel = baseModel
	actor.BaseAttributes = baseAttributes
	actor.BaseExternalModel = baseExternalModel
	if object["x_mitre_contributors"] != nil {
		actor.XMitreContributors = ConvertInterfaceArrayToStringArray(object["x_mitre_contributors"].([]interface{}))
	}
	for _, extRef := range actor.ExternalReferences {
		if extRef.SourceName == "mitre-attack" {
			actor.MitreAttckId = extRef.ExternalId
		}
	}
	return actor, nil
}

func (a *ActorObject) Malwares() ([]Malware, error) {
	return nil, nil
}

func (a *ActorObject) Tools() ([]Tool, error) {
	return nil, nil
}

func (a *ActorObject) Techniques() ([]Technique, error) {
	return nil, nil
}

func (a *ActorObject) ToConsole() {
	headers := []string{"Id", "Name", "Type", "Description"}
	rows := []string{a.Id, a.Name, a.Type, a.Description}
	DisplayOutput(headers, rows)
}

// Parses the external attributes of the actor
func parseActorExternalAttributes(object map[string]interface{}) (actorExternalAttributes, error) {
	aExternalAttributes := actorExternalAttributes{}
	if object["external_tools"] != nil {
		aExternalAttributes.ExternalTools = ConvertInterfaceArrayToStringArray(object["external_tools"].([]interface{}))
	}
	if object["country"] != nil {
		aExternalAttributes.Country = ConvertInterfaceArrayToStringArray(object["country"].([]interface{}))
	}
	if object["operations"] != nil {
		aExternalAttributes.Operations = ConvertInterfaceArrayToStringArray(object["operations"].([]interface{}))
	}
	if object["links"] != nil {
		aExternalAttributes.Links = ConvertInterfaceArrayToStringArray(object["links"].([]interface{}))
	}
	if object["targets"] != nil {
		aExternalAttributes.Targets = ConvertInterfaceArrayToStringArray(object["targets"].([]interface{}))
	}
	if object["external_description"] != nil {
		aExternalAttributes.ExternalDescription = ConvertInterfaceArrayToStringArray(object["external_description"].([]interface{}))
	}
	if object["attck_id"] != nil {
		aExternalAttributes.AttckID = object["attck_id"].(string)
	}
	if object["comment"] != nil {
		aExternalAttributes.Comment = object["comment"].(string)
	}
	if object["comments"] != nil {
		aExternalAttributes.Comments = ConvertInterfaceArrayToStringArray(object["comments"].([]interface{}))
	}
	return aExternalAttributes, nil
}
