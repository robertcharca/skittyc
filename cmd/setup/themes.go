package setup

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setTheme bool

var themesCmd = &cobra.Command{
	Use: "themes",	
	Short: "'themes' command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setTheme {
			res := prompts.HandleSetTheme()
			fmt.Printf("res: %s\n", res)

			handleKittyConfTheme(res)
		}
	},
}

func init() {
	SetupCmd.AddCommand(themesCmd)
	themesCmd.Flags().BoolVarP(&setTheme, "set", "s", false, "Setting themes.")
}

func handleKittyConfTheme(path string) {
	option, exists := prompts.ConfirmKittyConfExistence()
	fmt.Printf("option: %s, exists: %t\n", option, exists)	

	if !exists{
		// Function to create an empty file and replace it with the theme.
		// Return a boolean value.
	}

	kfeatures.ReplacingKittyFile(path)

	switch option {
	case "save it as a profile":
		// Function to create a new kitty conf file that will have the theme.
	case "replace it":
		// Function to empty the kitty file and rewrite it with the theme.
	default:
		// Function for quitting the theme implementation process. 
	}
	// Function to replace an existing file with the theme.
	// Return a boolean value.
}
