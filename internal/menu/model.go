package menu

import (
	"fmt"
	"os"
	"strings"

	bubblelist "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lipglosslist "github.com/charmbracelet/lipgloss/list"
	"golang.org/x/term"

	"github.com/msadministrator/goattck"
)

type menuState int

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

const (
	stateSelectEntity menuState = iota
	stateSelectItem
	stateViewDetails
)

type model struct {
	list           bubblelist.Model
	tabs           []mitreTab
	selectedTab    int
	enterprise     goattck.Enterprise
	entitySelected string
	selectedItem   string
	state          menuState
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			switch m.state {
			case stateSelectEntity:
				m.entitySelected = ""
				m.state = stateSelectEntity
				m.selectedItem = ""
			case stateSelectItem:
				m.state = stateSelectEntity
				m.selectedItem = ""
			case stateViewDetails:
				m.state = stateSelectItem
				m.selectedItem = ""
			}
		case "enter":
			switch m.state {
			case stateSelectEntity:
				if i, ok := m.list.SelectedItem().(item); ok {
					m.entitySelected = i.title
					m.state = stateSelectItem
					// Update list items based on selected entity
					m.updateEntityList()
				}
			case stateSelectItem:
				if i, ok := m.list.SelectedItem().(item); ok {
					m.selectedItem = i.title
					m.state = stateViewDetails
				}
			}
		case "right":
			if m.selectedTab < len(mitreTabs)-1 {
				m.selectedTab++
			}
		case "left":
			if m.selectedTab > 0 {
				m.selectedTab--
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *model) updateEntityList() {
	var items []bubblelist.Item
	switch m.entitySelected {
	case "actors":
		for _, actor := range m.enterprise.Actors {
			items = append(items, item{
				title: actor.Name,
				desc:  actor.Description,
			})
		}
	case "campaigns":
		for _, campaign := range m.enterprise.Campaigns {
			items = append(items, item{
				title: campaign.Name,
				desc:  campaign.Description,
			})
		}
	case "controls":
		for _, control := range m.enterprise.Controls {
			items = append(items, item{
				title: control.Name,
				desc:  "",
			})
		}
	case "data sources":
		for _, source := range m.enterprise.DataSources {
			items = append(items, item{
				title: source.Name,
				desc:  source.Description,
			})
		}
	case "malwares":
		for _, malware := range m.enterprise.Malwares {
			items = append(items, item{
				title: malware.Name,
				desc:  malware.Description,
			})
		}
	case "mitigations":
		for _, mitigation := range m.enterprise.Mitigations {
			items = append(items, item{
				title: mitigation.Name,
				desc:  mitigation.Description,
			})
		}
	case "tactics":
		for _, tactic := range m.enterprise.Tactics {
			items = append(items, item{
				title: tactic.Name,
				desc:  tactic.Description,
			})
		}
	case "techniques":
		for _, technique := range m.enterprise.Techniques {
			items = append(items, item{
				title: technique.Name,
				desc:  technique.Description,
			})
		}
	case "tools":
		for _, tool := range m.enterprise.Tools {
			items = append(items, item{
				title: tool.Name,
				desc:  tool.Description,
			})
		}
	}
	m.list.SetItems(items)
}

func (m model) View() string {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	// Main content area
	switch m.state {
	case stateSelectEntity:
		// Selects the actors, techniques, tools, etc. entity
		doc.WriteString("Select an entity type:\n\n")
		doc.WriteString(m.list.View())
	case stateSelectItem:
		// Selects the specific entity
		doc.WriteString(fmt.Sprintf("Select a %s:\n\n", m.entitySelected))
		doc.WriteString(m.list.View())
	case stateViewDetails:
		// View the details of the specific entity

		// We get our title on our detailed view
		doc.WriteString(m.getEntityTitle() + "\n\n")

		// Add the selected entity string
		doc.WriteString(fmt.Sprintf("Details for %s: %s\n\n", m.entitySelected, m.selectedItem))

		// Build our tabs
		doc.WriteString(m.getEntityTabs() + "\n\n")

		// Display entity details based on type
		doc.WriteString(m.getEntityDetails())
	}

	// At the end of each detailed entity view, we display a status bar
	doc.WriteString(m.getStaturBar())

	// Now we set the width and write our doc string to the console
	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	return docStyle.Render(doc.String())
}

// Builds the bottom status bar
func (m model) getStaturBar() string {
	w := lipgloss.Width

	statusKey := statusStyle.Render("STATUS")
	encoding := encodingStyle.Render("UTF-8")
	state := statusText.Render(m.getStateString())
	statusVal := statusText.
		Width(width - w(statusKey) - w(encoding)).
		Render(state)

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		encoding,
	)

	return statusBarStyle.Width(width).Render(bar)
}

// Builds the different tabs on the detail entity view
func (m model) getEntityTabs() string {
	var row string
	switch m.selectedTab {
	case 0:
		row = overviewTab.GetRow()
	case 1:
		row = relationshipsTab.GetRow()
	case 2:
		row = referencesTab.GetRow()
	case 3:
		row = extendedTab.GetRow()
	case 4:
		row = externalTab.GetRow()
	}

	gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
	return lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
}

// Builds the entity title section of our details display
func (m model) getEntityTitle() string {
	var (
		colors = colorGrid(1, 5)
		title  strings.Builder
	)

	// We get the string value berfore loading our menu
	titleStyle = titleStyle.SetString(m.getStateStringForTitle())

	for i, v := range colors {
		const offset = 2
		c := lipgloss.Color(v[0])
		fmt.Fprint(&title, titleStyle.MarginLeft(i*offset).Background(c))
		if i < len(colors)-1 {
			title.WriteRune('\n')
		}
	}

	desc := lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("MITRE ATT&CK Enterprise Framework"),
		infoStyle.Render("Interactive Menu System"),
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
}

// Returns a string for use when building the entity title
func (m model) getStateStringForTitle() string {
	switch m.entitySelected {
	case "actors":
		return "Actor"
	case "campaigns":
		return "Campaign"
	case "data sources":
		return "Data Source"
	case "malwares":
		return "Malware"
	case "mitigations":
		return "Mitigation"
	case "tactics":
		return "Tactic"
	case "techniques":
		return "Technique"
	case "tools":
		return "Tool"
	default:
		return "Unknown"
	}
}

// Displays the bottom status bar representing the current state
func (m model) getStateString() string {
	switch m.state {
	case stateSelectEntity:
		return "Select Entity Type"
	case stateSelectItem:
		return fmt.Sprintf("Select %s", m.entitySelected)
	case stateViewDetails:
		return fmt.Sprintf("Viewing %s: %s", m.entitySelected, m.selectedItem)
	default:
		return "Unknown State"
	}
}

func (m model) getEntityDetails() string {
	var details strings.Builder

	switch m.entitySelected {
	case "actors":
		for _, actor := range m.enterprise.Actors {
			if actor.Name == m.selectedItem {
				switch m.selectedTab {
				case 0: // Overview tab

					// Set the description first
					details.WriteString(descriptionStyle.Width(width).Render(fmt.Sprintf("%s", actor.Description)))
					details.WriteString("\n")

					// We now have a list of different attributers we want to display
					// in our table here
					attributes := map[string]string{
						"Id":                actor.Id,
						"Type":              actor.Type,
						"Created":           actor.Created,
						"Modified":          actor.Modified,
						"Version":           actor.XMitreVersion,
						"Domains":           strings.Join(actor.XMitreDomains, ","),
						"Revoked":           fmt.Sprintf("%v", actor.Revoked),
						"Deprecated":        fmt.Sprintf("%v", actor.XMitreDeprecated),
						"CreatedByRef":      actor.CreatedByRef,
						"ObjectMarkingrefs": strings.Join(actor.ObjectMarkingRefs, ""),
					}

					rows := [][]string{}
					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					// we display the table
					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))

					details.WriteString("\n")

					// If we have aliases, we display them
					if len(actor.Aliases) > 0 {
						aList := []string{}
						aList = append(aList, actor.Aliases...)
						details.WriteString(lipglosslist.New("Aliases", lipglosslist.New(aList)).String())
					}
					details.WriteString("\n\n")
				case 1: // Relationships tab
					var malwareList, toolList, techniqueList []string

					if len(actor.Malwares) > 0 {
						for _, malware := range actor.Malwares {
							malwareList = append(malwareList, fmt.Sprintf("%s (%s)", malware.Name, malware.GetExternalID()))
						}
					}

					if len(actor.Tools) > 0 {
						for _, tool := range actor.Tools {
							toolList = append(toolList, fmt.Sprintf("%s (%s)", tool.Name, tool.GetExternalID()))
						}
					}

					if len(actor.Techniques) > 0 {
						for _, technique := range actor.Techniques {
							techniqueList = append(techniqueList, fmt.Sprintf("%s (%s)", technique.Name, technique.GetExternalID()))
						}
					}

					l := lipglosslist.New(
						"Malwares", lipglosslist.New(malwareList),
						"Tools", lipglosslist.New(toolList),
						"Techniques", lipglosslist.New(techniqueList),
					)

					details.WriteString(l.String())
					details.WriteString("\n\n")
				case 2: // References tab
					if len(actor.ExternalReferences) > 0 {
						l := lipglosslist.New()
						for _, ref := range actor.ExternalReferences {
							if ref.Url != "" {
								l.Item(fmt.Sprintf("[%s](%s)", ref.SourceName, ref.Url))
							}
						}
						details.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, l.String()))
					}
					details.WriteString("\n\n")
				case 3: // Extended tab
					// This section displays extra attributes
					rows := [][]string{}

					attributes := map[string]string{
						"Contributors":            strings.Join(actor.XMitreContributors, ","),
						"XMitreModifiedByRef":     actor.XMitreModifiedByRef,
						"XMitreAttackSpecVersion": actor.XMitreAttackSpecVersion,
					}

					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))
					details.WriteString("\n\n")
				case 4: // External
					// TODO
				}
			}
		}
	case "campaigns":
		for _, campaign := range m.enterprise.Campaigns {
			if campaign.Name == m.selectedItem {
				switch m.selectedTab {
				case 0: // Overview
					// Set the description first
					details.WriteString(descriptionStyle.Width(width).Render(fmt.Sprintf("%s", campaign.Description)))
					details.WriteString("\n")

					// We now have a list of different attributers we want to display
					// in our table here
					attributes := map[string]string{
						"Id":                campaign.Id,
						"Type":              campaign.Type,
						"Created":           campaign.Created,
						"Modified":          campaign.Modified,
						"Version":           campaign.XMitreVersion,
						"Domains":           strings.Join(campaign.XMitreDomains, ","),
						"Revoked":           fmt.Sprintf("%v", campaign.Revoked),
						"Deprecated":        fmt.Sprintf("%v", campaign.XMitreDeprecated),
						"CreatedByRef":      campaign.CreatedByRef,
						"ObjectMarkingrefs": strings.Join(campaign.ObjectMarkingRefs, ""),
					}

					rows := [][]string{}
					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					// we display the table
					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))

					details.WriteString("\n\n")

				case 1: // Relationships
					var malwareList, toolList, techniqueList []string

					if len(campaign.Malwares) > 0 {
						for _, malware := range campaign.Malwares {
							malwareList = append(malwareList, fmt.Sprintf("%s (%s)", malware.Name, malware.GetExternalID()))
						}
					}

					if len(campaign.Tools) > 0 {
						for _, tool := range campaign.Tools {
							toolList = append(toolList, fmt.Sprintf("%s (%s)", tool.Name, tool.GetExternalID()))
						}
					}

					if len(campaign.Techniques) > 0 {
						for _, technique := range campaign.Techniques {
							techniqueList = append(techniqueList, fmt.Sprintf("%s (%s)", technique.Name, technique.GetExternalID()))
						}
					}

					l := lipglosslist.New(
						"Malwares", lipglosslist.New(malwareList),
						"Tools", lipglosslist.New(toolList),
						"Techniques", lipglosslist.New(techniqueList),
					)

					details.WriteString(l.String())
					details.WriteString("\n\n")

				case 2: // References
					if len(campaign.ExternalReferences) > 0 {
						l := lipglosslist.New()
						for _, ref := range campaign.ExternalReferences {
							if ref.Url != "" {
								l.Item(fmt.Sprintf("[%s](%s)", ref.SourceName, ref.Url))
							}
						}
						details.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, l.String()))
					}
					details.WriteString("\n\n")

				case 3: // Extended
					// This section displays extra attributes
					rows := [][]string{}

					attributes := map[string]string{
						"First Seen":            campaign.FirstSeen.String(),
						"First Seen Citation":   campaign.XMitreFirstSeenCitation,
						"Last Seen":             campaign.LastSeen.String(),
						"Last Seen Citation":    campaign.XMitreLastSeenCitation,
						"Modified by Reference": campaign.XMitreModifiedByRef,
						"Attack Spec Version":   campaign.XMitreAttackSpecVersion,
					}

					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))
					details.WriteString("\n\n")
				case 4: // External
					//TODO
				}
			}
		}
	case "data sources":
		for _, source := range m.enterprise.DataSources {
			if source.Name == m.selectedItem {
				switch m.selectedTab {
				case 0: // Overview
					// Set the description first
					details.WriteString(descriptionStyle.Width(width).Render(fmt.Sprintf("%s", source.Description)))
					details.WriteString("\n")

					// We now have a list of different attributers we want to display
					// in our table here
					attributes := map[string]string{
						"Id":                source.Id,
						"Type":              source.Type,
						"Created":           source.Created,
						"Modified":          source.Modified,
						"Version":           source.XMitreVersion,
						"Domains":           strings.Join(source.XMitreDomains, ","),
						"Revoked":           fmt.Sprintf("%v", source.Revoked),
						"Deprecated":        fmt.Sprintf("%v", source.XMitreDeprecated),
						"CreatedByRef":      source.CreatedByRef,
						"ObjectMarkingrefs": strings.Join(source.ObjectMarkingRefs, ""),
					}

					rows := [][]string{}
					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					// we display the table
					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))

					details.WriteString("\n\n")

				case 1: // Relationships
					var componentList, techniqueList []string

					if len(source.DataComponents) > 0 {
						for _, malware := range source.DataComponents {
							componentList = append(componentList, fmt.Sprintf("%s (%s)", malware.Name, malware.GetExternalID()))
						}
					}

					if len(source.Techniques) > 0 {
						for _, source := range source.Techniques {
							techniqueList = append(techniqueList, fmt.Sprintf("%s", source.Name))
						}
					}

					l := lipglosslist.New(
						"Data Components", lipglosslist.New(componentList),
						"Techniques", lipglosslist.New(techniqueList),
					)

					details.WriteString(l.String())
					details.WriteString("\n\n")

				case 2: // References
					if len(source.ExternalReferences) > 0 {
						l := lipglosslist.New()
						for _, ref := range source.ExternalReferences {
							if ref.Url != "" {
								l.Item(fmt.Sprintf("[%s](%s)", ref.SourceName, ref.Url))
							}
						}
						details.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, l.String()))
					}
					details.WriteString("\n\n")

				case 3: // Extended
					// This section displays extra attributes
					rows := [][]string{}

					attributes := map[string]string{
						"Contributors":          strings.Join(source.XMitreContributors, ","),
						"Platforms":             strings.Join(source.XMitrePlatforms, ","),
						"Modified by Reference": source.XMitreModifiedByRef,
						"Attack Spec Version":   source.XMitreAttackSpecVersion,
					}

					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))
					details.WriteString("\n\n")
				case 4: // External
					//TODO
				}
			}
		}
	case "techniques":
		for _, technique := range m.enterprise.Techniques {
			if technique.Name == m.selectedItem {
				switch m.selectedTab {
				case 0: // Overview
					// Set the description first
					details.WriteString(descriptionStyle.Width(width).Render(fmt.Sprintf("%s", technique.Description)))
					details.WriteString("\n")

					// We now have a list of different attributers we want to display
					// in our table here
					attributes := map[string]string{
						"Id":                technique.Id,
						"Type":              technique.Type,
						"Created":           technique.Created,
						"Modified":          technique.Modified,
						"Version":           technique.XMitreVersion,
						"Domains":           strings.Join(technique.XMitreDomains, ","),
						"Revoked":           fmt.Sprintf("%v", technique.Revoked),
						"Deprecated":        fmt.Sprintf("%v", technique.XMitreDeprecated),
						"CreatedByRef":      technique.CreatedByRef,
						"ObjectMarkingrefs": strings.Join(technique.ObjectMarkingRefs, ""),
						"IsSubtechnique":    fmt.Sprintf("%v", technique.XMitreIsSubtechnique),
					}

					rows := [][]string{}
					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					// we display the table
					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))

					details.WriteString("\n\n")

				case 1: // Relationships
					var actorList, malwareList, toolList, datasourceList []string

					if len(technique.Malwares) > 0 {
						for _, malware := range technique.Malwares {
							malwareList = append(malwareList, fmt.Sprintf("%s (%s)", malware.Name, malware.GetExternalID()))
						}
					}

					if len(technique.Tools) > 0 {
						for _, tool := range technique.Tools {
							toolList = append(toolList, fmt.Sprintf("%s (%s)", tool.Name, tool.GetExternalID()))
						}
					}

					if len(technique.Actors) > 0 {
						for _, actor := range technique.Actors {
							actorList = append(actorList, fmt.Sprintf("%s (%s)", actor.Name, actor.GetExternalID()))
						}
					}

					if len(technique.DataSources) > 0 {
						for _, source := range technique.DataSources {
							datasourceList = append(datasourceList, fmt.Sprintf("%s", source.Name))
						}
					}

					l := lipglosslist.New(
						"Malwares", lipglosslist.New(malwareList),
						"Tools", lipglosslist.New(toolList),
						"Actors", lipglosslist.New(actorList),
						"Data Sources", lipglosslist.New(datasourceList),
					)

					details.WriteString(l.String())
					details.WriteString("\n\n")

				case 2: // References
					if len(technique.ExternalReferences) > 0 {
						l := lipglosslist.New()
						for _, ref := range technique.ExternalReferences {
							if ref.Url != "" {
								l.Item(fmt.Sprintf("[%s](%s)", ref.SourceName, ref.Url))
							}
						}
						details.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, l.String()))
					}
					details.WriteString("\n\n")

				case 3: // Extended
					// This section displays extra attributes
					rows := [][]string{}

					attributes := map[string]string{
						"Contributors":          strings.Join(technique.XMitreContributors, ","),
						"Modified by Reference": technique.XMitreModifiedByRef,
						"Attack Spec Version":   technique.XMitreAttackSpecVersion,
						"Data Sources":          strings.Join(technique.XMitreDataSources, ","),
						"Effective Permissions": strings.Join(technique.XMitreEffectivePermissions, ","),
						"Remote Support":        fmt.Sprintf("%v", technique.XMitreRemoteSupport),
						"Permissions Required":  strings.Join(technique.XMitrePermissionsRequired, ","),
						"Detection":             technique.XMitreDetection,
						"Defense Bypassed":      strings.Join(technique.XMitreDefenseBypassed, ","),
						"System Requirements":   strings.Join(technique.XMitreSystemRequirements, ","),
						"Platforms":             strings.Join(technique.XMitrePlatforms, ","),
						"Network Requirements":  fmt.Sprintf("%v", technique.XMitreNetworkRequirements),
					}

					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))
					details.WriteString("\n\n")
				case 4: // External
					//TODO
				}
			}
		}
	case "malwares":
		for _, malware := range m.enterprise.Malwares {
			if malware.Name == m.selectedItem {
				switch m.selectedTab {
				case 0: // Overview
					// Set the description first
					details.WriteString(descriptionStyle.Width(width).Render(fmt.Sprintf("%s", malware.Description)))
					details.WriteString("\n")

					// We now have a list of different attributers we want to display
					// in our table here
					attributes := map[string]string{
						"Id":                malware.Id,
						"Type":              malware.Type,
						"Created":           malware.Created,
						"Modified":          malware.Modified,
						"Version":           malware.XMitreVersion,
						"Domains":           strings.Join(malware.XMitreDomains, ","),
						"Revoked":           fmt.Sprintf("%v", malware.Revoked),
						"Deprecated":        fmt.Sprintf("%v", malware.XMitreDeprecated),
						"CreatedByRef":      malware.CreatedByRef,
						"ObjectMarkingrefs": strings.Join(malware.ObjectMarkingRefs, ""),
					}

					rows := [][]string{}
					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					// we display the table
					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))

					details.WriteString("\n\n")
				case 1: // Relationships
					var actorList, campaignList, techniqueList []string

					if len(malware.Actors) > 0 {
						for _, malware := range malware.Actors {
							actorList = append(actorList, fmt.Sprintf("%s (%s)", malware.Name, malware.GetExternalID()))
						}
					}

					if len(malware.Campaigns) > 0 {
						for _, tool := range malware.Campaigns {
							campaignList = append(campaignList, fmt.Sprintf("%s (%s)", tool.Name, tool.GetExternalID()))
						}
					}

					if len(malware.Techniques) > 0 {
						for _, technique := range malware.Techniques {
							techniqueList = append(techniqueList, fmt.Sprintf("%s (%s)", technique.Name, technique.GetExternalID()))
						}
					}

					l := lipglosslist.New(
						"Actors", lipglosslist.New(actorList),
						"Campaigns", lipglosslist.New(campaignList),
						"Techniques", lipglosslist.New(techniqueList),
					)

					details.WriteString(l.String())
					details.WriteString("\n\n")
				case 2: // References
					if len(malware.ExternalReferences) > 0 {
						l := lipglosslist.New()
						for _, ref := range malware.ExternalReferences {
							if ref.Url != "" {
								l.Item(fmt.Sprintf("[%s](%s)", ref.SourceName, ref.Url))
							}
						}
						details.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, l.String()))
					}
					details.WriteString("\n\n")
				case 3: // Extended
					// This section displays extra attributes
					rows := [][]string{}

					attributes := map[string]string{
						"Contributors":          strings.Join(malware.XMitreContributors, ","),
						"Modified by Reference": malware.XMitreModifiedByRef,
						"Attack Spec Version":   malware.XMitreAttackSpecVersion,
					}

					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))
					details.WriteString("\n\n")
				case 4: // External
					// TODO
				}
			}
		}
		// Add more cases for other entity types as needed
	case "mitigations":
		for _, mitigation := range m.enterprise.Mitigations {
			if mitigation.Name == m.selectedItem {
				switch m.selectedTab {
				case 0: // Overview
					// Set the description first
					details.WriteString(descriptionStyle.Width(width).Render(fmt.Sprintf("%s", mitigation.Description)))
					details.WriteString("\n")

					// We now have a list of different attributers we want to display
					// in our table here
					attributes := map[string]string{
						"Id":                mitigation.Id,
						"Type":              mitigation.Type,
						"Created":           mitigation.Created,
						"Modified":          mitigation.Modified,
						"Version":           mitigation.XMitreVersion,
						"Domains":           strings.Join(mitigation.XMitreDomains, ","),
						"Revoked":           fmt.Sprintf("%v", mitigation.Revoked),
						"Deprecated":        fmt.Sprintf("%v", mitigation.XMitreDeprecated),
						"CreatedByRef":      mitigation.CreatedByRef,
						"ObjectMarkingrefs": strings.Join(mitigation.ObjectMarkingRefs, ""),
					}

					rows := [][]string{}
					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					// we display the table
					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))

					details.WriteString("\n\n")
				case 1: // Relationships
					var techniqueList []string

					if len(mitigation.Techniques) > 0 {
						for _, technique := range mitigation.Techniques {
							techniqueList = append(techniqueList, fmt.Sprintf("%s (%s)", technique.Name, technique.GetExternalID()))
						}
					}

					l := lipglosslist.New(
						"Techniques", lipglosslist.New(techniqueList),
					)

					details.WriteString(l.String())
					details.WriteString("\n\n")
				case 2: // References
					if len(mitigation.ExternalReferences) > 0 {
						l := lipglosslist.New()
						for _, ref := range mitigation.ExternalReferences {
							if ref.Url != "" {
								l.Item(fmt.Sprintf("[%s](%s)", ref.SourceName, ref.Url))
							}
						}
						details.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, l.String()))
					}
					details.WriteString("\n\n")
				case 3: // Extended
					// This section displays extra attributes
					rows := [][]string{}

					attributes := map[string]string{
						"Modified by Reference": mitigation.XMitreModifiedByRef,
						"Attack Spec Version":   mitigation.XMitreAttackSpecVersion,
					}

					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))
					details.WriteString("\n\n")
				case 4: // External
					// TODO
				}
			}
		}
	case "tactics":
		for _, tactic := range m.enterprise.Tactics {
			if tactic.Name == m.selectedItem {
				switch m.selectedTab {
				case 0: // Overview
					// Set the description first
					details.WriteString(descriptionStyle.Width(width).Render(fmt.Sprintf("%s", tactic.Description)))
					details.WriteString("\n")

					// We now have a list of different attributers we want to display
					// in our table here
					attributes := map[string]string{
						"Id":                tactic.Id,
						"Type":              tactic.Type,
						"Created":           tactic.Created,
						"Modified":          tactic.Modified,
						"Version":           tactic.XMitreVersion,
						"Domains":           strings.Join(tactic.XMitreDomains, ","),
						"CreatedByRef":      tactic.CreatedByRef,
						"ObjectMarkingrefs": strings.Join(tactic.ObjectMarkingRefs, ""),
					}

					rows := [][]string{}
					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					// we display the table
					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))

					details.WriteString("\n\n")
				case 1: // Relationships
					var techniqueList []string

					if len(tactic.Techniques) > 0 {
						for _, technique := range tactic.Techniques {
							techniqueList = append(techniqueList, fmt.Sprintf("%s (%s)", technique.Name, technique.GetExternalID()))
						}
					}

					l := lipglosslist.New(
						"Techniques", lipglosslist.New(techniqueList),
					)

					details.WriteString(l.String())
					details.WriteString("\n\n")
				case 2: // References
					if len(tactic.ExternalReferences) > 0 {
						l := lipglosslist.New()
						for _, ref := range tactic.ExternalReferences {
							if ref.Url != "" {
								l.Item(fmt.Sprintf("[%s](%s)", ref.SourceName, ref.Url))
							}
						}
						details.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, l.String()))
					}
					details.WriteString("\n\n")
				case 3: // Extended
					// This section displays extra attributes
					rows := [][]string{}

					attributes := map[string]string{
						"Modified by Reference": tactic.XMitreModifiedByRef,
						"Attack Spec Version":   tactic.XMitreAttackSpecVersion,
					}

					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))
					details.WriteString("\n\n")
				case 4: // External
					// TODO
				}
			}
		}
	case "tools":
		for _, tool := range m.enterprise.Tools {
			if tool.Name == m.selectedItem {
				switch m.selectedTab {
				case 0: // Overview
					// Set the description first
					details.WriteString(descriptionStyle.Width(width).Render(fmt.Sprintf("%s", tool.Description)))
					details.WriteString("\n")

					// We now have a list of different attributers we want to display
					// in our table here
					attributes := map[string]string{
						"Id":                tool.Id,
						"Type":              tool.Type,
						"Created":           tool.Created,
						"Modified":          tool.Modified,
						"Version":           tool.XMitreVersion,
						"Domains":           strings.Join(tool.XMitreDomains, ","),
						"Revoked":           fmt.Sprintf("%v", tool.Revoked),
						"Deprecated":        fmt.Sprintf("%v", tool.XMitreDeprecated),
						"CreatedByRef":      tool.CreatedByRef,
						"ObjectMarkingrefs": strings.Join(tool.ObjectMarkingRefs, ""),
					}

					rows := [][]string{}
					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					// we display the table
					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))

					details.WriteString("\n\n")
				case 1: // Relationships
					var actorList, campaignList, techniqueList []string

					if len(tool.Actors) > 0 {
						for _, technique := range tool.Actors {
							actorList = append(actorList, fmt.Sprintf("%s (%s)", technique.Name, technique.GetExternalID()))
						}
					}

					if len(tool.Campaigns) > 0 {
						for _, technique := range tool.Campaigns {
							campaignList = append(campaignList, fmt.Sprintf("%s (%s)", technique.Name, technique.GetExternalID()))
						}
					}

					if len(tool.Techniques) > 0 {
						for _, technique := range tool.Techniques {
							techniqueList = append(techniqueList, fmt.Sprintf("%s (%s)", technique.Name, technique.GetExternalID()))
						}
					}

					l := lipglosslist.New(
						"Actors", lipglosslist.New(actorList),
						"Campaigns", lipglosslist.New(campaignList),
						"Techniques", lipglosslist.New(techniqueList),
					)

					details.WriteString(l.String())
					details.WriteString("\n\n")
				case 2: // References
					if len(tool.ExternalReferences) > 0 {
						l := lipglosslist.New()
						for _, ref := range tool.ExternalReferences {
							if ref.Url != "" {
								l.Item(fmt.Sprintf("[%s](%s)", ref.SourceName, ref.Url))
							}
						}
						details.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, l.String()))
					}
					details.WriteString("\n\n")
				case 3: // Extended
					// This section displays extra attributes
					rows := [][]string{}

					attributes := map[string]string{
						"Modified by Reference": tool.XMitreModifiedByRef,
						"Attack Spec Version":   tool.XMitreAttackSpecVersion,
					}

					for key, val := range attributes {
						rows = append(rows, []string{
							key,
							val,
						})
					}

					details.WriteString(getFormattedTable([]string{"Attribute", "Value"}, rows))
					details.WriteString("\n\n")
				case 4: // External
					// TODO
				}
			}
		}
	}

	return details.String()
}
