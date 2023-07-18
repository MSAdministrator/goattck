/*
Copyright © 2023 Josh Rickard @MSAdministrator
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"net/http"
	"io/ioutil"

	"github.com/msadministrator/goattck/internal/models"
	"github.com/spf13/cobra"
)

const attackURL = "https://raw.githubusercontent.com/mitre/cti/master/enterprise-attack/enterprise-attack.json"

var (
	EnterpriseAttck Enterprise
	EnterpriseAttckJson rawEnterpriseAttck
)

type Enterprise struct {
	Actors     		[]models.ActorObject
	Campaigns  		[]models.CampaignObject
	Controls   		[]models.ControlObject
	DataComponents 	[]models.DataComponentObject
	DataSources 	[]models.DataSourceObject
	Defintions 		[]models.MarkingDefinitionObject
	Identities 		[]models.IdentityObject
	Malwares   		[]models.MalwareObject
	Matrices   		[]models.MatrixObject
	Mitigations 	[]models.MitigationObject
	Relationships 	[]models.RelationshipObject
	Tactics 		[]models.TacticObject
	Techniques 		[]models.TechniqueObject
	Tools      		[]models.ToolObject
	rawData       	*rawEnterpriseAttck
}

type rawEnterpriseAttck struct {
	Type        string        `json:"type"`
	ID          string        `json:"id"`
	Objects     []interface{} `json:"objects"`
	SpecVersion string        `json:"spec_version"`
}

// enterpriseCmd represents the enterprise command
var enterpriseCmd = &cobra.Command{
	Use:   "enterprise",
	Short: "A brief description of your command",
	Long: ReturnLogo(),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		empty, err := IsStructEmpty(EnterpriseAttck)
		if err != nil {
			slogger.Error("Error, could not determine if Enterprise is empty")
		}
		if empty {
			bytesData, err := Fetch()
			if err != nil {
				slogger.Error("Error, could not fetch data")
			}
			eAttck := rawEnterpriseAttck{}
			json.Unmarshal(bytesData, &eAttck)
			EnterpriseAttckJson = eAttck
		}

		var enterprise Enterprise

		for _, value := range EnterpriseAttckJson.Objects {
			v, ok := value.(map[string]interface{})
			if !ok {
				slogger.Error("error casting value to map")
			}
			switch v["type"] {
			case "intrusion-set":
				actor, err := models.NewActor(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating actor: %s", err))
				}
				enterprise.Actors = append(enterprise.Actors, actor)
			case "campaign":
				campaign, err := models.NewCampaign(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating campaign: %s", err))
				}
				enterprise.Campaigns = append(enterprise.Campaigns, campaign)
			case "x-mitre-data-component":
				dataComponent, err := models.NewDataComponent(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating data component: %s", err))
				}
				enterprise.DataComponents = append(enterprise.DataComponents, dataComponent)
			case "x-mitre-data-source":
				dataSource, err := models.NewDataSource(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating data source: %s", err))
				}
				enterprise.DataSources = append(enterprise.DataSources, dataSource)
			case "marking-definition":
				markingDefinition, err := models.NewMarkingDefinition(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating marking definition: %s", err))
				}
				enterprise.Defintions = append(enterprise.Defintions, markingDefinition)
			case "identity":
				identity, err := models.NewIdentity(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating identity: %s", err))
				}
				enterprise.Identities = append(enterprise.Identities, identity)
			case "malware":
				malware, err := models.NewMalware(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating malware: %s", err))
				}
				enterprise.Malwares = append(enterprise.Malwares, malware)
			case "course-of-action":
				mitigation, err := models.NewMitigation(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating mitigation: %s", err))
				}
				enterprise.Mitigations = append(enterprise.Mitigations, mitigation)
			case "x-mitre-matrix":
				matrix, err := models.NewMatrix(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating matrix: %s", err))
				}
				enterprise.Matrices = append(enterprise.Matrices, matrix)
			case "relationship":
				relationship, err := models.NewRelationship(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating relationship: %s", err))
				}
				enterprise.Relationships = append(enterprise.Relationships, relationship)
			case "x-mitre-tactic":
				tactic, err := models.NewTactic(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating tactic: %s", err))
				}
				enterprise.Tactics = append(enterprise.Tactics, tactic)
			case "attack-pattern":
				technique, err := models.NewTechnique(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating technique: %s", err))
				}
				enterprise.Techniques = append(enterprise.Techniques, technique)
			case "tool":
				tool, err := models.NewTool(v)
				if err != nil {
					slogger.Error(fmt.Sprintf("Error creating tool: %s", err))
				}
				enterprise.Tools = append(enterprise.Tools, tool)
			}
		}
		EnterpriseAttck = enterprise
	},
}

func init() {
	rootCmd.AddCommand(enterpriseCmd)
}


// Fetch MITRE ATT&CK data
func Fetch() ([]byte, error) {
	slogger.Info("Fetching MITRE ATT&CK...")
	resp, err := http.Get(attackURL)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		slogger.Fatal(fmt.Sprintf("Error reading response body: %s", readErr))
		return nil, readErr
	}
	return body, nil
}

func IsStructEmpty(object interface{}) (bool, error) {
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, nil
}
