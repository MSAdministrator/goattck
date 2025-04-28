package menu

import (
	"fmt"
	"os"

	bubblelist "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/msadministrator/goattck"
)

const (
	// Just setting this to a standard width for now.
	width = 120
)

// Style definitions.
// Most of this was borrowed from their example
// https://github.com/charmbracelet/lipgloss/blob/master/examples/layout/main.go
var (

	// General.
	normal    = lipgloss.Color("#EEEEEE")
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	base = lipgloss.NewStyle().Foreground(normal)

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	url = lipgloss.NewStyle().Foreground(special).Render

	// Tabs.
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlight).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true)

	tabGap = tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	// Title.
	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB"))

	descStyle = base.MarginTop(1)

	infoStyle = base.
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

	// description.
	descriptionStyle = lipgloss.NewStyle().
				Align(lipgloss.Center).
				Foreground(lipgloss.Color("#FAFAFA")).
				Background(highlight).
				Margin(1, 3, 0, 0).
				Padding(1, 2)

	// Status Bar.
	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	// Page.
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

// The main function which loads our attck data model
func Load(attck goattck.Enterprise) {
	items := []bubblelist.Item{
		item{title: "actors", desc: "View threat actors and groups"},
		item{title: "campaigns", desc: "View threat campaigns"},
		item{title: "controls", desc: "View security controls"},
		item{title: "data sources", desc: "View data sources"},
		item{title: "malwares", desc: "View malware"},
		item{title: "mitigations", desc: "View mitigations"},
		item{title: "tactics", desc: "View tactics"},
		item{title: "techniques", desc: "View techniques"},
		item{title: "tools", desc: "View tools"},
	}

	m := model{
		list:       bubblelist.New(items, bubblelist.NewDefaultDelegate(), 0, 0),
		tabs:       mitreTabs,
		enterprise: attck,
		state:      stateSelectEntity,
	}
	m.list.Title = "MITRE ATT&CK Enterprise Framework"
	m.list.SetShowStatusBar(false)
	m.list.SetFilteringEnabled(true)
	m.list.Styles.Title = titleStyle

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
