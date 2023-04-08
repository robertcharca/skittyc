package setup

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SetupCmd = &cobra.Command{
	Use: "setup",	
	Short: "Welcome to 'setup'.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Section for setting up your principal config to the Kitty terminal.")
	},
}

func init() {
	//Funcionality for the 'setup' command.
}
