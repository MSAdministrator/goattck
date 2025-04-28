package models

import (
	"fmt"
)

type identity interface{}

type Identity struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	ObjectMarkingRefs []string `json:"object_marking_refs"`
	IdentityClass     string   `json:"identity_class"`
}

var _ (identity) = new(Identity)

func NewIdentity(object map[string]interface{}) (Identity, error) {
	identity := Identity{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	identity.BaseModel = baseModel
	if object["object_marking_refs"] != nil {
		identity.ObjectMarkingRefs = ConvertInterfaceArrayToStringArray(object["object_marking_refs"].([]interface{}))
	}
	if object["identity_class"] != nil {
		identity.IdentityClass = object["identity_class"].(string)
	}
	return identity, nil
}
