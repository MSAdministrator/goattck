/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// actorsCmd represents the actors command
var actorsCmd = &cobra.Command{
	Use:   "actors",
	Short: "Actors or Groups that have been identified as threat actors or adversaries.",
	Long: `By default this will return all actors or groups that have been identified as threat actors or adversaries.
	
	You can use the '--id' flag to return a specific actor or group by providing the ID of the actor or group.
	Additionally, data is enriched by external data from pyattck-data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("actors called")
	},
}

func init() {
	enterpriseCmd.AddCommand(actorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
