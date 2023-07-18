/*
Copyright © 2023 Josh Rickard @MSAdministrator
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/msadministrator/goattck/internal/models"
)

// actorsCmd represents the actors command
var actorsCmd = &cobra.Command{
	Use:   "actors",
	Short: "Actors or Groups that have been identified as threat actors or adversaries.",
	Long: ReturnLogo() + `By default this will return all actors or groups that have been identified as threat actors or adversaries.

	You can use the '--id' flag to return a specific actor or group by providing the MitreAttckId of the actor or group.
	Additionally, data is enriched by external data from pyattck-data.`,
	Aliases: []string{"groups", "adversaries"},
	SuggestFor: []string{"group", "adversary"},
	ValidArgs: []string{"id", "name", "type"},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Figure a good way to parse the data base on flags
		pretty, _ := cmd.Flags().GetBool("pretty")
		name, _ := cmd.Flags().GetString("name")
		id, _ := cmd.Flags().GetString("id")
		var headers table.Row
		headers = table.Row{"MitreAttckId", "Name", "Type", "Aliases", "Revoked"}
		var rows []table.Row
		var actors []models.ActorObject
		for _, actor := range EnterpriseAttck.Actors {
			if name == "" && id == "" {
				actors = append(actors, actor)
				if pretty {
					rows = append(rows, table.Row{actor.MitreAttckId, actor.Name, actor.Type, actor.Aliases, actor.Revoked})
				}
			} else {
				if name != "" {
					if name == actor.Name {
						actors = append(actors, actor)
						if pretty {
							rows = append(rows, table.Row{actor.MitreAttckId, actor.Name, actor.Type, actor.Aliases, actor.Revoked})
						}
					}
				}
				if id != "" {
					if id == actor.MitreAttckId {
						actors = append(actors, actor)
						if pretty {
							rows = append(rows, table.Row{actor.MitreAttckId, actor.Name, actor.Type, actor.Aliases, actor.Revoked})
						}
					}
				}
			}
		}
		if pretty {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(headers)
			t.AppendRows(rows)
			t.Render()
		}
	},
}

func init() {
	enterpriseCmd.AddCommand(actorsCmd)
	// Local flags
	actorsCmd.Flags().StringP("id", "i", "", "The MitreAttckId of the actor or group")
	actorsCmd.Flags().StringP("name", "n", "", "The name of the actor or group")
}
