package models

import (
	"fmt"
)

type Technique interface {
	Actors() ([]Actor, error)
	Campaigns() ([]Campaign, error)
	DataComponents() ([]DataComponent, error)
	DataSources() ([]DataSource, error)
	Malwares() ([]Malware, error)
	Mitigations() ([]Mitigation, error)
	Tactics() ([]Tactic, error)
	Techniques() ([]Technique, error)
	Tools() ([]Tool, error)
}

type TechniqueObject struct {
	BaseModel
	BaseAttributes
	// These are properties from the MITRE ATT&CK json
	XMitrePlatforms    []string `json:"x_mitre_platforms"`
	ExternalReferences []ExternalReference `json:"external_references"`
	KillChainPhases []struct {
		KillChainName string `json:"kill_chain_name"`
		PhaseName     string `json:"phase_name"`
	} `json:"kill_chain_phases"`
	XMitreDetection            string   `json:"x_mitre_detection,omitempty"`
	XMitreIsSubtechnique       bool     `json:"x_mitre_is_subtechnique"`
	XMitreModifiedByRef        string   `json:"x_mitre_modified_by_ref"`
	XMitreDataSources          []string `json:"x_mitre_data_sources,omitempty"`
	XMitreDefenseBypassed      []string `json:"x_mitre_defense_bypassed,omitempty"`
	XMitreContributors         []string `json:"x_mitre_contributors,omitempty"`
	XMitrePermissionsRequired  []string `json:"x_mitre_permissions_required,omitempty"`
	XMitreRemoteSupport        bool     `json:"x_mitre_remote_support,omitempty"`
	XMitreAttackSpecVersion    string   `json:"x_mitre_attack_spec_version,omitempty"`
	XMitreSystemRequirements   []string `json:"x_mitre_system_requirements,omitempty"`
	XMitreImpactType           []string `json:"x_mitre_impact_type,omitempty"`
	XMitreEffectivePermissions []string `json:"x_mitre_effective_permissions,omitempty"`
	XMitreNetworkRequirements  bool     `json:"x_mitre_network_requirements,omitempty"`
	techniqueExternalAttributes
}

type techniqueExternalAttributes struct {
	// These are properties external from the MITRE ATT&CK json definitions
	CommandList []string `json:"command_list"`
	Commands []string `json:"commands"`
	Queries []string `json:"queries"`
	ParsedDatasets []string `json:"parsed_datasets"`
	PossibleDetections []string`json:"possible_detections"`
	ExternalReference []string `json:"external_reference"`
	Controls []string `json:"controls"`
}

func parseTechniqueExternalAttributes(object map[string]interface{}) (techniqueExternalAttributes, error) {
	techniqueExternalAttributes := techniqueExternalAttributes{}
	if object["command_list"] != nil {
		techniqueExternalAttributes.CommandList = ConvertInterfaceArrayToStringArray(object["command_list"].([]interface{}))
	}
	if object["commands"] != nil {
		techniqueExternalAttributes.Commands = ConvertInterfaceArrayToStringArray(object["commands"].([]interface{}))
	}
	if object["queries"] != nil {
		techniqueExternalAttributes.Queries = ConvertInterfaceArrayToStringArray(object["queries"].([]interface{}))
	}
	if object["parsed_datasets"] != nil {
		techniqueExternalAttributes.ParsedDatasets = ConvertInterfaceArrayToStringArray(object["parsed_datasets"].([]interface{}))
	}
	if object["possible_detections"] != nil {
		techniqueExternalAttributes.PossibleDetections = ConvertInterfaceArrayToStringArray(object["possible_detections"].([]interface{}))
	}
	if object["external_reference"] != nil {
		techniqueExternalAttributes.ExternalReference = ConvertInterfaceArrayToStringArray(object["external_reference"].([]interface{}))
	}
	if object["controls"] != nil {
		techniqueExternalAttributes.Controls = ConvertInterfaceArrayToStringArray(object["controls"].([]interface{}))
	}
	return techniqueExternalAttributes, nil
}

func parseKillChainPhases(object map[string]interface{}) []struct {
	KillChainName string `json:"kill_chain_name"`
	PhaseName     string `json:"phase_name"`
} {
	killChainPhases := []struct {
		KillChainName string `json:"kill_chain_name"`
		PhaseName     string `json:"phase_name"`
	}{}
	for _, killChainPhase := range object["kill_chain_phases"].([]interface{}) {
		killChainPhases = append(killChainPhases, struct {
			KillChainName string `json:"kill_chain_name"`
			PhaseName     string `json:"phase_name"`
		}{
			KillChainName: killChainPhase.(map[string]interface{})["kill_chain_name"].(string),
			PhaseName:     killChainPhase.(map[string]interface{})["phase_name"].(string),
		})
	}
	return killChainPhases
}

func NewTechnique(object map[string]interface{}) (TechniqueObject, error) {
	technique := TechniqueObject{}
	baseModel, err := parseBaseModel(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base model: %s", err))
	}
	technique.BaseModel = baseModel
	baseAttributes, err := parseBaseAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing base attributes: %s", err))
	}
	technique.BaseAttributes = baseAttributes
	if object["x_mitre_platforms"] != nil {
		technique.XMitrePlatforms = ConvertInterfaceArrayToStringArray(object["x_mitre_platforms"].([]interface{}))
	}
	if object["external_references"] != nil {
		refs, err := parseExternalReferences(object)
		if err != nil {
			slogger.Error(fmt.Sprintf("Error parsing external references: %s", err))
		}
		technique.ExternalReferences = refs
	}
	if object["kill_chain_phases"] != nil {
		technique.KillChainPhases = parseKillChainPhases(object)
	}
	if object["x_mitre_detection"] != nil {
		technique.XMitreDetection = object["x_mitre_detection"].(string)
	}
	if object["x_mitre_is_subtechnique"] != nil {
		technique.XMitreIsSubtechnique = object["x_mitre_is_subtechnique"].(bool)
	}
	if object["x_mitre_modified_by_ref"] != nil {
		technique.XMitreModifiedByRef = object["x_mitre_modified_by_ref"].(string)
	}
	if object["x_mitre_data_sources"] != nil {
		technique.XMitreDataSources = ConvertInterfaceArrayToStringArray(object["x_mitre_data_sources"].([]interface{}))
	}
	if object["x_mitre_defense_bypassed"] != nil {
		technique.XMitreDefenseBypassed = ConvertInterfaceArrayToStringArray(object["x_mitre_defense_bypassed"].([]interface{}))
	}
	if object["x_mitre_contributors"] != nil {
		technique.XMitreContributors = ConvertInterfaceArrayToStringArray(object["x_mitre_contributors"].([]interface{}))
	}
	if object["x_mitre_permissions_required"] != nil {
		technique.XMitrePermissionsRequired = ConvertInterfaceArrayToStringArray(object["x_mitre_permissions_required"].([]interface{}))
	}
	if object["x_mitre_remote_support"] != nil {
		technique.XMitreRemoteSupport = object["x_mitre_remote_support"].(bool)
	}
	if object["x_mitre_attack_spec_version"] != nil {
		technique.XMitreAttackSpecVersion = object["x_mitre_attack_spec_version"].(string)
	}
	if object["x_mitre_system_requirements"] != nil {
		technique.XMitreSystemRequirements = ConvertInterfaceArrayToStringArray(object["x_mitre_system_requirements"].([]interface{}))
	}
	if object["x_mitre_impact_type"] != nil {
		technique.XMitreImpactType = ConvertInterfaceArrayToStringArray(object["x_mitre_impact_type"].([]interface{}))
	}
	if object["x_mitre_effective_permissions"] != nil {
		technique.XMitreEffectivePermissions = ConvertInterfaceArrayToStringArray(object["x_mitre_effective_permissions"].([]interface{}))
	}
	if object["x_mitre_network_requirements"] != nil {
		technique.XMitreNetworkRequirements = object["x_mitre_network_requirements"].(bool)
	}
	techniqueExternalAttributes, err := parseTechniqueExternalAttributes(object)
	if err != nil {
		slogger.Error(fmt.Sprintf("Error parsing technique external attributes: %s", err))
	}
	technique.techniqueExternalAttributes = techniqueExternalAttributes
	return technique, nil
}

func (t *TechniqueObject) Actors() ([]Actor, error) {
	return nil, nil
}

func (t *TechniqueObject) Campaigns() ([]Campaign, error) {
	return nil, nil
}

func (t *TechniqueObject) DataComponents() ([]DataComponent, error) {
	return nil, nil
}

func (t *TechniqueObject) DataSources() ([]DataSource, error) {
	return nil, nil
}

func (t *TechniqueObject) Malwares() ([]Malware, error) {
	return nil, nil
}

func (t *TechniqueObject) Mitigations() ([]Mitigation, error) {
	return nil, nil
}

func (t *TechniqueObject) Tactics() ([]Tactic, error) {
	return nil, nil
}

func (t *TechniqueObject) Techniques() ([]Technique, error) {
	return nil, nil
}

func (t *TechniqueObject) Tools() ([]Tool, error) {
	return nil, nil
}
