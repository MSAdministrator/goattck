package models

import "time"

type EnterpriseAttck struct {
	Type    string `json:"type"`
	ID      string `json:"id"`
	Objects []struct {
		TacticRefs         []string  `json:"tactic_refs,omitempty"`
		ObjectMarkingRefs  []string  `json:"object_marking_refs,omitempty"`
		ID                 string    `json:"id"`
		Type               string    `json:"type"`
		Created            time.Time `json:"created"`
		CreatedByRef       string    `json:"created_by_ref,omitempty"`
		ExternalReferences []struct {
			ExternalID string `json:"external_id"`
			SourceName string `json:"source_name"`
			URL        string `json:"url"`
		} `json:"external_references,omitempty"`
		Modified                time.Time `json:"modified,omitempty"`
		Name                    string    `json:"name,omitempty"`
		Description             string    `json:"description,omitempty"`
		XMitreVersion           string    `json:"x_mitre_version,omitempty"`
		XMitreAttackSpecVersion string    `json:"x_mitre_attack_spec_version,omitempty"`
		XMitreModifiedByRef     string    `json:"x_mitre_modified_by_ref,omitempty"`
		XMitreDomains           []string  `json:"x_mitre_domains,omitempty"`
		XMitreDeprecated        bool      `json:"x_mitre_deprecated,omitempty"`
		Revoked                 bool      `json:"revoked,omitempty"`
		XMitrePlatforms         []string  `json:"x_mitre_platforms,omitempty"`
		XMitreAliases           []string  `json:"x_mitre_aliases,omitempty"`
		Labels                  []string  `json:"labels,omitempty"`
		XMitreContributors      []string  `json:"x_mitre_contributors,omitempty"`
		XMitreShortname         string    `json:"x_mitre_shortname,omitempty"`
		KillChainPhases         []struct {
			KillChainName string `json:"kill_chain_name"`
			PhaseName     string `json:"phase_name"`
		} `json:"kill_chain_phases,omitempty"`
		XMitreDetection            string    `json:"x_mitre_detection,omitempty"`
		XMitreIsSubtechnique       bool      `json:"x_mitre_is_subtechnique,omitempty"`
		XMitreDataSources          []string  `json:"x_mitre_data_sources,omitempty"`
		XMitreDefenseBypassed      []string  `json:"x_mitre_defense_bypassed,omitempty"`
		XMitrePermissionsRequired  []string  `json:"x_mitre_permissions_required,omitempty"`
		XMitreRemoteSupport        bool      `json:"x_mitre_remote_support,omitempty"`
		XMitreSystemRequirements   []string  `json:"x_mitre_system_requirements,omitempty"`
		XMitreImpactType           []string  `json:"x_mitre_impact_type,omitempty"`
		XMitreEffectivePermissions []string  `json:"x_mitre_effective_permissions,omitempty"`
		XMitreNetworkRequirements  bool      `json:"x_mitre_network_requirements,omitempty"`
		RelationshipType           string    `json:"relationship_type,omitempty"`
		SourceRef                  string    `json:"source_ref,omitempty"`
		TargetRef                  string    `json:"target_ref,omitempty"`
		Aliases                    []string  `json:"aliases,omitempty"`
		FirstSeen                  time.Time `json:"first_seen,omitempty"`
		LastSeen                   time.Time `json:"last_seen,omitempty"`
		XMitreFirstSeenCitation    string    `json:"x_mitre_first_seen_citation,omitempty"`
		XMitreLastSeenCitation     string    `json:"x_mitre_last_seen_citation,omitempty"`
		XMitreDataSourceRef        string    `json:"x_mitre_data_source_ref,omitempty"`
		XMitreCollectionLayers     []string  `json:"x_mitre_collection_layers,omitempty"`
		IdentityClass              string    `json:"identity_class,omitempty"`
		Definition                 struct {
			Statement string `json:"statement"`
		} `json:"definition,omitempty"`
		DefinitionType string `json:"definition_type,omitempty"`
	} `json:"objects"`
	SpecVersion string `json:"spec_version"`
}
