package models

type Identity struct {
	// Base fields
	BaseModel
	// Fields
	Type               string               `json:"type"`
	IdentityClass      string               `json:"identity_class"`
	Created            string               `json:"created"`
	Modified           string               `json:"modified"`
	Name               string               `json:"name"`
	ObjectMarkingRefs  []string             `json:"object_marking_refs"`
	Roles              []string             `json:"roles"`
	Sectors            []string             `json:"sectors"`
	ExternalReferences []ExternalReferences `json:"external_references"`
}
