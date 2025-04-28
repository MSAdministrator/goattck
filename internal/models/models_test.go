package models

import (
	"testing"
)

func TestNewCampaign(t *testing.T) {
	campaign, err := NewCampaign(loadTestJSON("campaign.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if campaign.Name != "SolarWinds Compromise" {
		t.Errorf("Error, could not load Campaign data models: %v", err)
	}
	if campaign.ExternalReferences[0].SourceName != "mitre-attack" {
		t.Errorf("Error, could not load Campaign data models: %v", err)
	}
}

func TestDataComponent(t *testing.T) {
	dataComponent, err := NewDataComponent(loadTestJSON("data_component.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if dataComponent.Name != "User Account Authentication" {
		t.Errorf("Error, could not load DataComponent data models: %v", err)
	}
	if dataComponent.XMitreAttackSpecVersion != "2.1.0" {
		t.Errorf("Error, could not load DataComponent data models: %v", err)
	}
}

func TestDataSource(t *testing.T) {
	dataSource, err := NewDataSource(loadTestJSON("data_source.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if dataSource.Name != "User Account" {
		t.Errorf("Error, could not load DataSource data models: %v", err)
	}
	if dataSource.XMitreVersion != "1.1" {
		t.Errorf("Error, could not load DataSource data models: %v", err)
	}
	if len(dataSource.XMitrePlatforms) != 9 {
		t.Errorf("Error, could not load DataSource data models: %v", err)
	}
}

func TestNewMarkingDefinition(t *testing.T) {
	markingDefinition, err := NewMarkingDefinition(loadTestJSON("definition.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if markingDefinition.DefinitionType != "statement" {
		t.Errorf("Error, could not load MarkingDefinition data models: %v", err)
	}
}

func TestNewMalware(t *testing.T) {
	malware, err := NewMalware(loadTestJSON("malware.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if malware.Name != "HDoor" {
		t.Errorf("Error, could not load Malware data models: %v", err)
	}
	if malware.XMitreVersion != "1.0" {
		t.Errorf("Error, could not load Malware data models: %v", err)
	}
	if len(malware.XMitrePlatforms) != 1 {
		t.Errorf("Error, could not load Malware data models: %v", err)
	}
	if malware.Windows != false {
		t.Errorf("Error, could not load Malware data models: %v", err)
	}
}

func TestNewMatrix(t *testing.T) {
	matrix, err := NewMatrix(loadTestJSON("matrix.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if matrix.Name != "Enterprise ATT&CK" {
		t.Errorf("Error, could not load Matrix data models: %v", err)
	}
	if len(matrix.TacticRefs) != 14 {
		t.Errorf("Error, could not load Matrix data models: %v", err)
	}
	if matrix.XMitreVersion != "1.0" {
		t.Errorf("Error, could not load Matrix data models: %v", err)
	}
}

func TestNewMitigation(t *testing.T) {
	mitigation, err := NewMitigation(loadTestJSON("mitigation.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if mitigation.Name != "Data from Information Repositories Mitigation" {
		t.Errorf("Error, could not load Mitigation data models: %v", err)
	}
	if len(mitigation.ExternalReferences) != 1 {
		t.Errorf("Error, could not load Mitigation data models: %v", err)
	}
	if mitigation.XMitreVersion != "1.0" {
		t.Errorf("Error, could not load Mitigation data models: %v", err)
	}
}

func TestNewRelationship(t *testing.T) {
	relationship, err := NewRelationship(loadTestJSON("relationship.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if relationship.Id != "relationship--74f14668-7111-4f96-a307-4aac00d91cf4" {
		t.Errorf("Error, could not load Relationship data models: %v", err)
	}
	if relationship.ExternalReferences[0].SourceName != "Carbon Black HotCroissant April 2020" {
		t.Errorf("Error, could not load Relationship data models: %v", err)
	}
}

func TestNewTactic(t *testing.T) {
	tactic, err := NewTactic(loadTestJSON("tactic.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if tactic.Name != "Credential Access" {
		t.Errorf("Error, could not load Tactic data models: %v", err)
	}
	if tactic.Modified != "2019-07-19T17:43:41.967Z" {
		t.Errorf("Error, could not load Tactic data models: %v", err)
	}
	if tactic.XMitreDomains[0] != "enterprise-attack" {
		t.Errorf("Error, could not load Tactic data models: %v", err)
	}
}

func TestNewTechnique(t *testing.T) {
	technique, err := NewTechnique(loadTestJSON("technique.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if technique.Name != "Scheduled Task" {
		t.Errorf("Error, could not load Technique data models: %v", err)
	}
	if len(technique.XMitreDataSources) != 5 {
		t.Errorf("Error, could not load Technique data models: %v", err)
	}
	if technique.KillChainPhases[0].PhaseName != "execution" {
		t.Errorf("Error, could not load Technique data models: %v", err)
	}
	if len(technique.KillChainPhases) != 3 {
		t.Errorf("Error, could not load Technique data models: %v", err)
	}
	if technique.TechniqueId != "T1053.005" {
		t.Errorf("Error, could not load Technique data models: %v", err)
	}
}

func TestNewTool(t *testing.T) {
	tool, err := NewTool(loadTestJSON("tool.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if tool.Name != "Empire" {
		t.Errorf("Error, could not load Tool data models: %v", err)
	}
	if tool.XMitreVersion != "1.6" {
		t.Errorf("Error, could not load Tool data models: %v", err)
	}
	if tool.XMitrePlatforms[0] != "Linux" {
		t.Errorf("Error, could not load Tool data models: %v", err)
	}
	if tool.ExternalReferences[5].SourceName != "NCSC Joint Report Public Tools" {
		t.Errorf("Error, could not load Tool data models: %v", err)
	}
	if tool.Server != "Python" {
		t.Errorf("Error, could not load Tool data models: %v", err)
	}
	if tool.MultiUser != false {
		t.Errorf("Error, could not load Tool data models: %v", err)
	}
}
