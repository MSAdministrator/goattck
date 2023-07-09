package models

import (
	"fmt"
	"reflect"
)

type Actor []struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	Description        string   `json:"description,omitempty"`
	Aliases            []string `json:"aliases,omitempty"`
	XMitreDeprecated   bool     `json:"x_mitre_deprecated,omitempty"`
	XMitreContributors []string `json:"x_mitre_contributors,omitempty"`
	CreatedByRef       string   `json:"created_by_ref,omitempty"`
	Revoked            bool     `json:"revoked,omitempty"`
	ExternalReferences []struct {
		SourceName  string `json:"source_name"`
		URL         string `json:"url,omitempty"`
		ExternalID  string `json:"external_id,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"external_references"`
	ObjectMarkingRefs       []string `json:"object_marking_refs,omitempty"`
	XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version,omitempty"`
	XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref,omitempty"`
	// These are properties unique to pyattck-data
	Names               []string `json:"names"`
	ExternalTools       []string `json:"external_tools"`
	Country             []string `json:"country"`
	Operations          []string `json:"operations"`
	Links               []string `json:"links"`
	Targets             []string `json:"targets"`
	ExternalDescription []string `json:"external_description"`
	AttckID             string   `json:"attck_id"`
	Comment             string   `json:"comment"`
}

func NewActor(object struct{}) (Actor, error) {
	actor := Actor{
		EnterpriseAttck: EnterpriseAttck{
			// These are properties from the MITRE ATT&CK json
			Type:        object.Type,
			ID:          object.ID,
			Name:        object.Name,
			Description: object.Description,
		},
	}
	slogger.Info(fmt.Sprintf("Inside NewActor. Created actor: %+v", actor.AttckID))
	return actor, nil
}

func (a *Actor) Malwares() ([]Malware, error) {
	return nil, nil
}

func (a *Actor) Tools() ([]Tool, error) {
	return nil, nil
}

func (a *Actor) Techniques() ([]Technique, error) {
	return nil, nil
}

func ObjectAssign(target interface{}, object interface{}) {
	// object atributes values in target atributes values
	// using pattern matching (https://golang.org/pkg/reflect/#Value.FieldByName)
	// https://stackoverflow.com/questions/35590190/how-to-use-the-spread-operator-in-golang
	t := reflect.ValueOf(target).Elem()
	o := reflect.ValueOf(object).Elem()
	for i := 0; i < o.NumField(); i++ {
		for j := 0; j < t.NumField(); j++ {
			if t.Field(j) == o.Field(i) {
				fmt.Printf("Field %s is equal to %s\n", t.Field(j), o.Field(i))
				t.Field(j).Set(o.Field(i))
			}
		}
	}
}
