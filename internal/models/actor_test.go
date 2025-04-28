package models

import (
	"testing"
)

func TestNewActor(t *testing.T) {
	actor, err := NewActor(loadTestJSON("actor.example.json"))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if actor.Name != "APT1" {
		t.Errorf("Error, could not load Actor data models: %v", err)
	}
	if actor.XMitreVersion != "1.4" {
		t.Errorf("Error, could not load Actor data models: %v", err)
	}
	if actor.ExternalReferences[0].SourceName != "mitre-attack" {
		t.Errorf("Error, could not load Actor data models: %v", err)
	}
}
