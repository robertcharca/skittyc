package customize

import (
	"fmt"
	
	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setTB, changeTB bool

var tabBarCmd = &cobra.Command{
	Use: "tabbar",
	Short: "'tabbar' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setTB == true {
			prompt, res := prompts.HandleSetTabBar()
			
			err := kfeatures.ChangingValues(prompt, res, "# Tab bar")
			if err != nil {
				fmt.Println(err.Error())
			}	
		}

		if changeTB == true {
			prompt, res := prompts.HandleChangeTabBar()

			err := kfeatures.ChangingValues(prompt, res, "# Tab bar")
			if err != nil {
				fmt.Println(err.Error())
			}	
		}
	},
}

func init() {
	CustomizeCmd.AddCommand(tabBarCmd)

	tabBarCmd.Flags().BoolVarP(&setTB, "set", "s", false, "Setting tab bar styles")
	tabBarCmd.Flags().BoolVarP(&changeTB, "change", "c", false, "Changing tab bat styles")
}
