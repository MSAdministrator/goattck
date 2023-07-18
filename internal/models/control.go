package models

import "fmt"

type Control interface {
	Techniques() ([]Technique, error)
}

type ControlObject struct {
	// Base fields
	BaseModel
	// Fields
	Revoked            bool                 `json:"revoked"`
	XMitreFamily       string               `json:"x_mitre_family"`
	XMitreImpact       []string             `json:"x_mitre_impact"`
	XMitrePriority     string               `json:"x_mitre_priority"`
	ObjectMarkingRefs  []string             `json:"object_marking_refs"`
	ExternalReferences []ExternalReference `json:"external_references"`
}

func NewControl(object map[string]interface{}) (ControlObject, error) {
	control := ControlObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	control.BaseModel = baseModel
	if object["revoked"] != nil {
		control.Revoked = object["revoked"].(bool)
	}
	if object["x_mitre_family"] != nil {
		control.XMitreFamily = object["x_mitre_family"].(string)
	}
	if object["x_mitre_impact"] != nil {
		control.XMitreImpact = object["x_mitre_impact"].([]string)
	}
	if object["x_mitre_priority"] != nil {
		control.XMitrePriority = object["x_mitre_priority"].(string)
	}
	if object["object_marking_refs"] != nil {
		control.ObjectMarkingRefs = object["object_marking_refs"].([]string)
	}
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		control.ExternalReferences = refs
	}
	return control, nil
}

func (c *ControlObject) Techniques() ([]Technique, error) {
	return nil, nil
}
