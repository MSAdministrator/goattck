/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/msadministrator/goattck/internal/logger"
	"github.com/spf13/cobra"
)

var slogger = logger.NewLogger(logger.Info, true)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goattck",
	Short: "A golang CLI for the MITRE ATT&CK Framework",
	Long: `This is a golang CLI for the MITRE ATT&CK Framework.
	It utilizes work from pyattck and the MITRE ATT&CK API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goattck.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// type MitreFramework interface {
// 	Actors() ([]Actor, error)
// 	Campaigns() ([]Campaign, error)
// 	Controls() ([]Control, error)
// 	DataComponents() ([]DataComponent, error)
// 	DataSources() ([]DataSource, error)
// 	Malwares() ([]Malware, error)
// 	Mitigations() ([]Mitigation, error)
// 	Relationships() ([]Relationship, error)
// 	Tactics() ([]Tactic, error)
// 	Techniques() ([]Technique, error)
// 	Tools() ([]Tool, error)
// }

// type MitreFramework interface {
// 	Actors() ([]Actor, error)
// 	Campaigns() ([]Campaign, error)
// 	Controls() ([]Control, error)
// 	DataComponents() ([]DataComponent, error)
// 	DataSources() ([]DataSource, error)
// 	Malwares() ([]Malware, error)
// 	Mitigations() ([]Mitigation, error)
// 	Relationships() ([]Relationship, error)
// 	Tactics() ([]Tactic, error)
// 	Techniques() ([]Technique, error)
// 	Tools() ([]Tool, error)
// }
