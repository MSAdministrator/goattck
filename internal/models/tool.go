package models

import (
	"encoding/json"
	"strings"
)

type tool interface{}

type Tool struct {
	BaseModel
	BaseAttributes
	XMitrePlatforms         []string            `json:"x_mitre_platforms,omitempty"`
	XMitreDomains           []string            `json:"x_mitre_domains,omitempty"`
	XMitreContributors      []string            `json:"x_mitre_contributors,omitempty"`
	XMitreAliases           []string            `json:"x_mitre_aliases,omitempty"`
	ExternalReferences      []ExternalReference `json:"external_references"`
	Labels                  []string            `json:"labels"`
	XMitreAttackSpecVersion string              `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string              `json:"x_mitre_modified_by_ref"`
	// Additional [] of related entities
	Actors     []Actor
	Campaigns  []Campaign
	Techniques []Technique
	// External data
	C2Data             interface{}   `json:"c2_data"`
	ExternalDataset    []interface{} `json:"external_dataset"`
	AdditionalNames    []string      `json:"additional_names"`
	AttributionLinks   []string      `json:"attribution_links"`
	AdditionalComments []string      `json:"additional_comments"`
	Names              []string      `json:"names"`
	Comments           []string      `json:"comments"`
	Family             []string      `json:"family"`
	Links              []string      `json:"links"`
	License            string        `json:"license"`
	Price              string        `json:"price"`
	Github             string        `json:"github"`
	Site               string        `json:"site"`
	Twitter            string        `json:"twitter"`
	Evaluator          string        `json:"evaluator"`
	Date               string        `json:"date"`
	Version            string        `json:"version"`
	Implementation     string        `json:"implementation"`
	HowTo              string        `json:"how_to"`
	Slingshot          string        `json:"slingshot"`
	Kali               string        `json:"kali"`
	Server             string        `json:"server"`
	Implant            string        `json:"implant"`
	MultiUser          bool          `json:"multi_user"`
	UI                 bool          `json:"ui"`
	DarkMode           bool          `json:"dark_mode"`
	Api                bool          `json:"api"`
	Windows            bool          `json:"windows"`
	Linux              bool          `json:"linux"`
	Macos              bool          `json:"macos"`
	Tcp                bool          `json:"tcp"`
	Http               bool          `json:"http"`
	Http2              bool          `json:"http2"`
	Http3              bool          `json:"http3"`
	Dns                bool          `json:"dns"`
	Doh                bool          `json:"doh"`
	Icmp               bool          `json:"icmp"`
	Ftp                bool          `json:"ftp"`
	Imap               bool          `json:"imap"`
	Mapi               bool          `json:"mapi"`
	SMB                bool          `json:"smb"`
	Ldap               bool          `json:"ldap"`
	KeyExchange        bool          `json:"key_exchange"`
	Stego              bool          `json:"stego"`
	ProxyAware         bool          `json:"proxy_aware"`
	Domainfront        bool          `json:"domainfront"`
	CustomProfile      bool          `json:"custom_profile"`
	Jitter             bool          `json:"jitter"`
	WorkingHours       bool          `json:"working_hours"`
	KillDate           bool          `json:"kill_date"`
	Chaining           bool          `json:"chaining"`
	Logging            bool          `json:"logging"`
	InWild             bool          `json:"in_wild"`
	AttckMapping       bool          `json:"attck_mapping"`
	Dashboard          bool          `json:"dashboard"`
	Blog               string        `json:"blog"`
	C2MatrixIndicators string        `json:"c2_matrix_indicators"`
	Jarm               bool          `json:"jarm"`
	ActivelyMaint      bool          `json:"actively_maint"`
	Slack              bool          `json:"slack"`
	SlackMembers       bool          `json:"slack_members"`
	GhIssues           bool          `json:"gh_issues"`
	Notes              string        `json:"notes"`
	SocksSupport       bool          `json:"socks_support"`
	XMitreOldAttackId  string        `json:"x_mitre_old_attack_id"`
}

var _ (tool) = new(Tool)

func NewTool(object map[string]interface{}) (Tool, error) {
	tool := Tool{}
	jsonString, _ := json.Marshal(object)
	json.Unmarshal(jsonString, &tool)
	return tool, nil
}

func (t *Tool) GetExternalID() string {
	mitreID := ""
	for _, ref := range t.ExternalReferences {
		if ref.ExternalId != "" && mitreID == "" {
			if strings.HasPrefix(ref.ExternalId, "S") {
				mitreID = ref.ExternalId
			}
		}
	}
	return mitreID
}
