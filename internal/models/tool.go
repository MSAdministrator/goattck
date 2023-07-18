package models

import (
	"fmt"
)

type Tool interface {
	Actors() ([]Actor, error)
	Campaigns() ([]Campaign, error)
	Techniques() ([]Technique, error)
}

type ToolObject struct {
	BaseModel
	BaseAttributes
	XMitrePlatforms    []string `json:"x_mitre_platforms,omitempty"`
	XMitreContributors []string `json:"x_mitre_contributors,omitempty"`
	XMitreAliases      []string `json:"x_mitre_aliases,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references"`
	Labels                  []string `json:"labels"`
	XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref"`
}

func NewTool(object map[string]interface{}) (ToolObject, error) {
	tool := ToolObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	tool.BaseModel = baseModel
	baseAttributes, err := parseBaseAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base attributes: %s", err))
	}
	tool.BaseAttributes = baseAttributes
	if object["x_mitre_platforms"] != nil {
		tool.XMitrePlatforms = ConvertInterfaceArrayToStringArray(object["x_mitre_platforms"].([]interface{}))
	}
	if object["x_mitre_contributors"] != nil {
		tool.XMitreContributors = ConvertInterfaceArrayToStringArray(object["x_mitre_contributors"].([]interface{}))
	}
	if object["x_mitre_aliases"] != nil {
		tool.XMitreAliases = ConvertInterfaceArrayToStringArray(object["x_mitre_aliases"].([]interface{}))
	}
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		tool.ExternalReferences = refs
	}
	if object["labels"] != nil {
		tool.Labels = ConvertInterfaceArrayToStringArray(object["labels"].([]interface{}))
	}
	if object["x_mitre_attack_spec_version"] != nil {
		tool.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	if object["x_mitre_modified_by_ref"] != nil {
		tool.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	return tool, nil
}

func (t *ToolObject) Actors() ([]Actor, error) {
	return nil, nil
}

func (t *ToolObject) Campaigns() ([]Campaign, error) {
	return nil, nil
}

func (t *ToolObject) Techniques() ([]Technique, error) {
	return nil, nil
}
