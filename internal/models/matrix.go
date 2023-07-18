package models

import (
	"fmt"
)

type Matrix interface {
}

type MatrixObject struct {
	BaseModel
	Type					string               `json:"type"`
	TacticRefs				[]string             `json:"tactic_refs"`
	CreatedByRef			string               `json:"created_by_ref"`
	Description				string               `json:"description"`
	Revoked					bool                 `json:"revoked"`
	XMitreDomains			[]string             `json:"x_mitre_domains"`
	ObjectMarkingRefs		[]string             `json:"object_marking_refs"`
	ExternalReferences		[]ExternalReference  `json:"external_references"`
	XMitreDeprecated		bool                 `json:"x_mitre_deprecated"`
	XMitreVersion			string               `json:"x_mitre_version"`
	XMitreModifiedByRef		string               `json:"x_mitre_modified_by_ref"`
	XMitreAttackSpecVersion string               `json:"x_mitre_attack_spec_version"`
}

func NewMatrix(object map[string]interface{}) (MatrixObject, error) {
	matrix := MatrixObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	matrix.BaseModel = baseModel
	if object["type"] != nil {
		matrix.Type = object["type"].(string)
	}
	if object["tactic_refs"] != nil {
		matrix.TacticRefs = ConvertInterfaceArrayToStringArray(object["tactic_refs"].([]interface{}))
	}
	if object["created_by_ref"] != nil {
		matrix.CreatedByRef = object["created_by_ref"].(string)
	}
	if object["description"] != nil {
		matrix.Description = object["description"].(string)
	}
	if object["revoked"] != nil {
		matrix.Revoked = object["revoked"].(bool)
	}
	if object["x_mitre_domains"] != nil {
		matrix.XMitreDomains = ConvertInterfaceArrayToStringArray(object["x_mitre_domains"].([]interface{}))
	}
	if object["object_marking_refs"] != nil {
		matrix.ObjectMarkingRefs = ConvertInterfaceArrayToStringArray(object["object_marking_refs"].([]interface{}))
	}
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		matrix.ExternalReferences = refs
	}
	if object["x_mitre_deprecated"] != nil {
		matrix.XMitreDeprecated = object["x_mitre_deprecated"].(bool)
	}
	if object["x_mitre_version"] != nil {
		matrix.XMitreVersion = object["x_mitre_version"].(string)
	}
	if object["x_mitre_modified_by_ref"] != nil {
		matrix.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	if object["x_mitre_attack_spec_version"] != nil {
		matrix.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	return matrix, nil
}
