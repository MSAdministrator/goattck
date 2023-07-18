package models

import (
	"fmt"
)

type Relationship interface {
}

type RelationshipObject struct {
	BaseAttributes
	Id 				   string               `json:"id"`
	Type 			   string               `json:"type"`
	Created            string               `json:"created"`
	Modified           string               `json:"modified"`
	SourceRef          string               `json:"source_ref"`
	TargetRef          string               `json:"target_ref"`
	RelationshipType   string               `json:"relationship_type"`
	XMitreVersion      string               `json:"x_mitre_version"`
	XMitreAttackSpecVersion string         `json:"x_mitre_attack_spec_version"`
	ExternalReferences []ExternalReference `json:"external_references"`
	XMitreModifiedByRef string              `json:"x_mitre_modified_by_ref"`
}

func NewRelationship(object map[string]interface{}) (RelationshipObject, error) {
	relationship := RelationshipObject{}
	baseAttributes, err := parseBaseAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base attributes: %s", err))
	}
	relationship.BaseAttributes = baseAttributes
	if object["id"] != nil {
		relationship.Id = object["id"].(string)
	}
	if object["type"] != nil {
		relationship.Type = object["type"].(string)
	}
	if object["created"] != nil {
		relationship.Created = object["created"].(string)
	}
	if object["modified"] != nil {
		relationship.Modified = object["modified"].(string)
	}
	if object["source_ref"] != nil {
		relationship.SourceRef = object["source_ref"].(string)
	}
	if object["target_ref"] != nil {
		relationship.TargetRef = object["target_ref"].(string)
	}
	if object["relationship_type"] != nil {
		relationship.RelationshipType = object["relationship_type"].(string)
	}
	if object["x_mitre_version"] != nil {
		relationship.XMitreVersion = object["x_mitre_version"].(string)
	}
	if object["x_mitre_attack_spec_version"] != nil {
		relationship.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		relationship.ExternalReferences = refs
	}
	if object["x_mitre_modified_by_ref"] != nil {
		relationship.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	return relationship, nil
}
