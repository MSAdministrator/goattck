package models

type Technique struct {
	BaseModel
	// These are properties from the MITRE ATT&CK json
	XMitrePlatforms    []string `json:"x_mitre_platforms"`
	ObjectMarkingRefs  []string `json:"object_marking_refs"`
	CreatedByRef       string   `json:"created_by_ref"`
	ExternalReferences []struct {
		SourceName  string `json:"source_name"`
		ExternalID  string `json:"external_id,omitempty"`
		URL         string `json:"url"`
		Description string `json:"description,omitempty"`
	} `json:"external_references"`
	Description     string `json:"description"`
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
	XMitreDeprecated           bool     `json:"x_mitre_deprecated,omitempty"`
	XMitrePermissionsRequired  []string `json:"x_mitre_permissions_required,omitempty"`
	XMitreRemoteSupport        bool     `json:"x_mitre_remote_support,omitempty"`
	Revoked                    bool     `json:"revoked,omitempty"`
	XMitreAttackSpecVersion    string   `json:"x_mitre_attack_spec_version,omitempty"`
	XMitreSystemRequirements   []string `json:"x_mitre_system_requirements,omitempty"`
	XMitreImpactType           []string `json:"x_mitre_impact_type,omitempty"`
	XMitreEffectivePermissions []string `json:"x_mitre_effective_permissions,omitempty"`
	XMitreNetworkRequirements  bool     `json:"x_mitre_network_requirements,omitempty"`
}
