/*
Copyright © 2023 Josh Rickard @MSAdministrator
*/
package cmd

import (
	"os"
	"fmt"
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/msadministrator/goattck/internal/logger"
)

var slogger = logger.NewLogger(logger.Info, true)

const logo = `                          
                      __    __          __    
   ____   _________ _/  |__/  |_  ____ |  | __
  / ___\ /  _ \__  \\   __\   __\/ ___\|  |/ /
 / /_/  >  <_> ) __ \|  |  |  | \  \___|    < 
 \___  / \____(____  /__|  |__|  \___  >__|_ \
/_____/            \/                \/     \/
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goattck",
	Short: "A golang CLI for the MITRE ATT&CK Framework",
	Long: ReturnLogo() + `This is a golang CLI for the MITRE ATT&CK Framework.
It utilizes work from pyattck and the MITRE ATT&CK API`,
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
	// Persistent Flags
	rootCmd.PersistentFlags().BoolP("pretty", "p", false, "Pretty print the output")
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func ReturnLogo() string {
	return fmt.Sprintf("\u200B %s", logo)
}