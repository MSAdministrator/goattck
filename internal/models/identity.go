package models

import (
	"fmt"
)

type Identity interface {
}

type IdentityObject struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	ObjectMarkingRefs []string `json:"object_marking_refs"`
	IdentityClass     string   `json:"identity_class"`
}

func NewIdentity(object map[string]interface{}) (IdentityObject, error) {
	identity := IdentityObject{}
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
