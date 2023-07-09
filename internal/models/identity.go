package models

type Identity struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	ObjectMarkingRefs []string `json:"object_marking_refs"`
	IdentityClass     string   `json:"identity_class"`
}
