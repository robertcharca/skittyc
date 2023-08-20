package customize

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setB, changeB bool

var bellCmd = &cobra.Command{
	Use: "bell",
	Short: "'bell' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setB == true {
			prompt, res := prompts.HandleSetBell()
			err := kfeatures.ChangingValues(prompt, res, "# Terminal bell")
			if err != nil {
				fmt.Println(err.Error())
			}	
		}

		if changeB == true {
			prompt, res := prompts.HandleChangeBell()	
			err := kfeatures.ChangingValues(prompt, res, "# Terminal bell")
			if err != nil {
				fmt.Println(err.Error())
			}	
		}
	},
}

func init() {
	CustomizeCmd.AddCommand(bellCmd)

	bellCmd.Flags().BoolVarP(&setB, "set", "s", false, "Setting terminal bell configurations")
	bellCmd.Flags().BoolVarP(&changeB, "change", "c", false, "Changing terminal bell configurations")
}
