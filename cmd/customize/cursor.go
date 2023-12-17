package customize

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setCur, changeCur bool

var cursorCmd = &cobra.Command{
	Use:   "cursor",
	Short: "'cursor' subcommand",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setCur {
			prompt, res := prompts.HandleSetCursor()
			err := kfeatures.ChangingValues(prompt, res, "# Cursor")
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if changeCur {
			prompt, res := prompts.HandleChangeCursor()
			err := kfeatures.ChangingValues(prompt, res, "# Cursor")
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	},
}

func init() {
	CustomizeCmd.AddCommand(cursorCmd)

	cursorCmd.Flags().BoolVarP(&setCur, "set", "s", false, "Setting cursor styles")
	cursorCmd.Flags().BoolVarP(&changeCur, "change", "c", false, "Changing cursor styles")
}
