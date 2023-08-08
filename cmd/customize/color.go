package customize

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var (
	setC bool
	changeC bool
)

var colorCmd = &cobra.Command{
	Use: "colors",
	Short: "'colors' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setC == true {
			prompt, res := prompts.HandleSetColor()
			
			if prompt != "color scheme" {
				kfeatures.ChangingValues(prompt, res, "# Colors")
			} else {
				fmt.Printf("prompt: %s, res: %s\n", prompt, res)	
			}
		}
	},
}

func init() {
	CustomizeCmd.AddCommand(colorCmd)	

	colorCmd.Flags().BoolVarP(&setC, "set", "s", false, "Setting color styles.")
}
