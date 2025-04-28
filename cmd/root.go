package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/msadministrator/goattck"

	"github.com/msadministrator/goattck/internal/logger"
	"github.com/msadministrator/goattck/internal/menu"
)

var (
	slogger         = logger.NewLogger(logger.Debug)
	defaultAttckUrl = "https://raw.githubusercontent.com/mitre/cti/master/enterprise-attack/enterprise-attack.json"
	attck           goattck.Enterprise
)

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

var enterpriseCmd = &cobra.Command{
	Use:   "enterprise",
	Short: "Sets the scope of the CLI tool to the Enterprise MITRE ATT&CK Franework.",
	Long:  ReturnLogo() + "Sets the scope of the CLI tool to the Enterprise MITRE ATT&CK Franework.",
	Run: func(cmd *cobra.Command, args []string) {
		// Simple placeholder
	},
}

var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Displays an interactive UI for Mitre ATT&CK",
	Long:  ReturnLogo() + "This launches the cool (mid-90s) interactive menu using frameworks like bubbletea and lipgloss to make it look all fancy",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Parent().Name() == "enterprise" {

			enterprise, err := goattck.Enterprise{}.New(defaultAttckUrl)
			if err != nil {
				slogger.Fatal(fmt.Sprintf("Error loading MITRE ATT&CK data: %v\n", err))
			}
			if len(args) > 0 {
				val, err := strconv.ParseBool(args[0])
				if err != nil {
					slogger.Fatal("Unknown arguemnt value provided.")
				}
				if val {
					// if Force == true & we didn't just download it
					slogger.Debug(fmt.Sprintf("downloading and saving latest json data from %s", defaultAttckUrl))
					enterprise, err = enterprise.Load(val)
					if err != nil {
						slogger.Fatal(fmt.Sprintf("unable to download from url: %s %s", defaultAttckUrl, err))
					}
				}
			}

			attck, err = enterprise.Load(false)
			if err != nil {
				slogger.Fatal(fmt.Sprintf("Error loading MITRE ATT&CK data: %v\n", err))
			}

			menu.Load(attck)
		} else {
			slogger.Warning(fmt.Sprintf("unknown parent %s use 'enterprise' as parent instead", cmd.Parent().Name()))
		}
	},
}

func Execute() {
	var forceDownload bool
	rootCmd.PersistentFlags().BoolVar(&forceDownload, "force", false, "Force download latest Mitre ATT&CK JSON")
	enterpriseCmd.AddCommand(menuCmd)
	rootCmd.AddCommand(enterpriseCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func ReturnLogo() string {
	return fmt.Sprintf("\u200B %s", logo)
}
