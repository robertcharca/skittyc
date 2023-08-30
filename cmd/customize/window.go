package customize

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setW, changeW bool

var windowCmd = &cobra.Command{
	Use: "window",
	Short: "'window' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {	
		if setW {
			prompt, res := prompts.HandleSetWindow()
		
			err := kfeatures.ChangingValues(prompt, res, "# Window layout")
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if changeW {
			prompt, res := prompts.HandleChangeWindow()

			err := kfeatures.ChangingValues(prompt, res, "# Window layout")
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	},
}

func init() {
	CustomizeCmd.AddCommand(windowCmd)

	windowCmd.Flags().BoolVarP(&setW, "set", "s", false, "Setting window styles")
	windowCmd.Flags().BoolVarP(&changeW, "change", "c", false, "Changing window styles")
}
