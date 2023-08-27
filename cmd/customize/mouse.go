package customize

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setM, changeM bool

var mouseCmd = &cobra.Command{
	Use: "mouse",
	Short: "'mouse' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setM == true { 
			prompt, res := prompts.HandleSetMouse()
			
			err := kfeatures.ChangingValues(prompt, res, "# Mouse")
			if err != nil {
				fmt.Println(err.Error())
			}	
		}

		if changeM == true { 
			prompt, res := prompts.HandleChangeMouse()
			err := kfeatures.ChangingValues(prompt, res, "# Mouse")
			if err != nil {
				fmt.Println(err.Error())
			}	
		}
	},
}

func init() {		
	CustomizeCmd.AddCommand(mouseCmd)
	
	mouseCmd.Flags().BoolVarP(&changeM, "change", "c", false, "Changing mouse configurations.")
	
	mouseCmd.Flags().BoolVarP(&setM, "set", "s", false, "Set mouse configurations.")	
}
