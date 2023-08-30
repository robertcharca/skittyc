package customize

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setBF, changeBF bool

var backforeCmd = &cobra.Command{
	Use: "backfore",
	Short: "'background' and 'foreground'subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if changeBF {
			prompt, res := prompts.HandleChangeBackfore()
			
			err := kfeatures.ChangingValues(prompt, res, "# Background and Foreground")
			if err != nil {
				fmt.Println(err.Error())
			} 
		}

		if setBF {
			prompt, res := prompts.HandleSetBackfore()
			
			err := kfeatures.ChangingValues(prompt, res, "# Background and Foreground")
			if err != nil {
				fmt.Println(err.Error())
			} 
		}
	},
}

func init() {		
	CustomizeCmd.AddCommand(backforeCmd)
	
	backforeCmd.Flags().BoolVarP(&changeBF, "change", "c", false, "Changing background styles.")
	
	backforeCmd.Flags().BoolVarP(&setBF, "set", "s", false, "Set backgroud styles.")	
}
