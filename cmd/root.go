package cmd

import (
	"fmt"
	"os"

	"github.com/robertcharca/skittyc/cmd/customize"
	"github.com/robertcharca/skittyc/kittyc"
	"github.com/robertcharca/skittyc/cmd/setup"	
	"github.com/spf13/cobra"
)

var (
	createKittyConf bool
)

//Variables that stores the root command in the CLI.
var rootCmd = &cobra.Command{
	Use: "skittyc",
	Short: "Introduction for a Kitty customizer",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to skitty")
		
		//Creating kitty.conf file inside of .config/kitty
		if createKittyConf == true {
			kittyc.CreateKittyConf()
		}
	},
}

//Execute(): executes all commands starting from the "rootCmd" variable.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

//init(): works for initializing commands or flags (of any kind).
func init() {	
	//AddCommand(): method for adding principal subcommands for skittyc.
	rootCmd.AddCommand(customize.CustomizeCmd)
	rootCmd.AddCommand(setup.SetupCmd)

	//Local flags
	rootCmd.Flags().BoolVar(&createKittyConf, "create", false, "create a kitty.conf file")

}
