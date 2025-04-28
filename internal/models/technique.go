package models

import (
	"encoding/json"
	"strings"
)

type technique interface{}

type Technique struct {
	BaseModel
	BaseAttributes
	// These are properties from the MITRE ATT&CK json
	XMitrePlatforms    []string            `json:"x_mitre_platforms"`
	ExternalReferences []ExternalReference `json:"external_references"`
	KillChainPhases    []struct {
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
	Actors         []Actor
	Campaigns      []Campaign
	DataComponents []DataComponent
	DataSources    []DataSource
	Malwares       []Malware
	Mitigations    []Mitigation
	Tactics        []Tactic
	Techniques     []Technique
	Tools          []Tool
}

var _ (technique) = new(Technique)

type techniqueExternalAttributes struct {
	// These are properties external from the MITRE ATT&CK json definitions
	CommandList        []string `json:"command_list"`
	Commands           []string `json:"commands"`
	Queries            []string `json:"queries"`
	ParsedDatasets     []string `json:"parsed_datasets"`
	PossibleDetections []string `json:"possible_detections"`
	ExternalReference  []string `json:"external_reference"`
	Controls           []string `json:"controls"`
	TechniqueId        string   `json:"technique_id"`
}

func NewTechnique(object map[string]interface{}) (Technique, error) {
	technique := Technique{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &technique)
	return technique, nil
}

func (t *Technique) GetExternalID() string {
	mitreID := ""
	for _, ref := range t.ExternalReferences {
		if ref.ExternalId != "" && mitreID == "" {
			if strings.HasPrefix(ref.ExternalId, "T") {
				mitreID = ref.ExternalId
			}
		}
	}
	return mitreID
}
