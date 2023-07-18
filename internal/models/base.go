package models

type BaseModel struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Created       string   `json:"created"`
	Modified      string   `json:"modified"`
	XMitreVersion string   `json:"x_mitre_version"`
	XMitreDomains []string `json:"x_mitre_domains"`
	Type          string   `json:"type"`
}

type BaseAttributes struct {
	Description       string   `json:"description"`
	XMitreDeprecated  bool     `json:"x_mitre_deprecated"`
	CreatedByRef      string   `json:"created_by_ref"`
	Revoked           bool     `json:"revoked"`
	ObjectMarkingRefs []string `json:"object_marking_refs"`
}

type BaseExternalModel struct {
	Aliases                 []string             `json:"aliases"`
	ExternalReferences      []ExternalReference  `json:"external_references"`
	XMitreAttackSpecVersion string               `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string               `json:"x_mitre_modified_by_ref"`
}

type ExternalReference struct {
	SourceName  string `json:"source_name"`
	Url         string `json:"url"`
	ExternalId  string `json:"external_id"`
	Description string `json:"description"`
}

func parseBaseAttributes(object map[string]interface{}) (BaseAttributes, error) {
	baseAttributes := BaseAttributes{}
	if object["description"] != nil {
		baseAttributes.Description = object["description"].(string)
	}
	if object["x_mitre_deprecated"] != nil {
		baseAttributes.XMitreDeprecated = object["x_mitre_deprecated"].(bool)
	}
	if object["created_by_ref"] != nil {
		baseAttributes.CreatedByRef = object["created_by_ref"].(string)
	}
	if object["revoked"] != nil {
		baseAttributes.Revoked = object["revoked"].(bool)
	}
	if object["object_marking_refs"] != nil {
		baseAttributes.ObjectMarkingRefs = ConvertInterfaceArrayToStringArray(object["object_marking_refs"].([]interface{}))
	}
	return baseAttributes, nil
}

func parseExternalReference(object map[string]interface{}) (ExternalReference, error) {
	externalReference := ExternalReference{}
	if object["source_name"] != nil {
		externalReference.SourceName = object["source_name"].(string)
	}
	if object["url"] != nil {
		externalReference.Url = object["url"].(string)
	}
	if object["external_id"] != nil {
		externalReference.ExternalId = object["external_id"].(string)
	}
	if object["description"] != nil {
		externalReference.Description = object["description"].(string)
	}
	return externalReference, nil
}

func parseExternalReferences(object map[string]interface{}) ([]ExternalReference, error) {
	extRefs := object["external_references"].([]interface{})
	externalReferences := make([]ExternalReference, len(extRefs))
	for i, v := range extRefs {
		ref, err := parseExternalReference(v.(map[string]interface{}))
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		externalReferences[i] = ref
	}
	return externalReferences, nil
}

func parseExternalModel(object map[string]interface{}) (BaseExternalModel, error) {
	baseExternalModel := BaseExternalModel{}
	if object["aliases"] != nil {
		baseExternalModel.Aliases = ConvertInterfaceArrayToStringArray(object["aliases"].([]interface{}))
	}
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		baseExternalModel.ExternalReferences = refs
	}
	if object["x_mitre_attack_spec_version"] != nil {
		baseExternalModel.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	if object["x_mitre_modified_by_ref"] != nil {
		baseExternalModel.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	return baseExternalModel, nil
}

func parseBaseModel(object map[string]interface{}) (BaseModel, error) {
	baseModel := BaseModel{}
	if object["id"] != nil {
		baseModel.Id = object["id"].(string)
	}
	if object["name"] != nil {
		baseModel.Name = object["name"].(string)
	}
	if object["created"] != nil {
		baseModel.Created = object["created"].(string)
	}
	if object["modified"] != nil {
		baseModel.Modified = object["modified"].(string)
	}
	if object["x_mitre_version"] != nil {
		baseModel.XMitreVersion = object["x_mitre_version"].(string)
	}
	if object["type"] != nil {
		baseModel.Type = object["type"].(string)
	}
	if object["x_mitre_domains"] != nil {
		baseModel.XMitreDomains = ConvertInterfaceArrayToStringArray(object["x_mitre_domains"].([]interface{}))
	}
	return baseModel, nil
}
