/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/msadministrator/goattck/internal/models"
	"github.com/msadministrator/goattck/internal/utils"
	"github.com/spf13/cobra"
)

var AttckJson models.EnterpriseAttck

// enterpriseCmd represents the enterprise command
var enterpriseCmd = &cobra.Command{
	Use:   "enterprise",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		resp, err := utils.Fetch()
		if err != nil {
			fmt.Println("error fetching MITRE ATT&CK")
		}
		AttckJson = resp
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enterprise called")
		loadStructs()

	},
}

func init() {
	rootCmd.AddCommand(enterpriseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enterpriseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enterpriseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Recast(a, b interface{}) (interface{}, error) {
	js, _ := json.Marshal(a)

	return json.Unmarshal(js, b), nil
}

func loadStructs() ([]interface{}, error) {
	//attackPatterns := []models.Technique{}
	//	relationships := map[string][]relationship{}
	for _, obj := range AttckJson.Objects {
		switch obj.Type {
		case "attack-pattern":
			temp, _ := Recast(obj, &models.Technique{})
			fmt.Printf("temp: %v", temp)

			// temp := *models.Technique(obj)
			// attackPatterns = append(attackPatterns, temp)
			//	fmt.Println(obj.Name)
			// case "course-of-action":
			// 	fmt.Println(obj.Name)
			// case "intrusion-set", "malware", "tool", "campaign":
			// 	fmt.Println(obj.Name)
			// case "x-mitre-tactic":
			// 	fmt.Println(obj.Name)
			// case "x-mitre-data-source":
			// 	fmt.Println(obj.Name)
			// case "x-mitre-data-component":
			// 	fmt.Println(obj.Name)
			// case "relationship":
			// 	relationships[obj.TargetRef] = append(relationships[obj.TargetRef], relationship{
			// 		id:               obj.ID,
			// 		description:      obj.Description,
			// 		relationshipType: obj.RelationshipType,
			// 		sourceRef:        obj.SourceRef,
			// 		targetRef:        obj.TargetRef,
			// 		references:       obj.ExternalReferences,
			// 	})
		}
	}
	return nil, nil
}
