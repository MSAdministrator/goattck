package menu

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	overviewTab      = mitreTab{name: "Overview"}
	relationshipsTab = mitreTab{name: "Relationships"}
	referencesTab    = mitreTab{name: "References"}
	extendedTab      = mitreTab{name: "Extended"}
	externalTab      = mitreTab{name: "External"}

	mitreTabs = []mitreTab{overviewTab, relationshipsTab, referencesTab, extendedTab, externalTab}
)

type tabName string

type mitreTab struct {
	name tabName
}

func (t mitreTab) GetRow() string {
	var otherTabs []string

	for _, tabName := range mitreTabs {
		if tabName.name != t.name {
			otherTabs = append(otherTabs, tab.Render(string(tabName.name)))
		} else {
			otherTabs = append(otherTabs, activeTab.Render(string(t.name)))
		}
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		otherTabs...,
	)
}
