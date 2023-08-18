package customize

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setCur, changeCur bool 

var cursorCmd = &cobra.Command{
	Use: "cursor",
	Short: "'cursor' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setCur == true {
			prompt, res := prompts.HandleSetCursor()
			kfeatures.ChangingValues(prompt, res, "# Cursor")
			fmt.Printf("prompt: %s, res: %s\n", prompt, res)
		}

		if changeCur == true {
			prompt, res := prompts.HandleChangeCursor()
			kfeatures.ChangingValues(prompt, res, "# Cursor")
		}
	},
}

func init() {
	CustomizeCmd.AddCommand(cursorCmd)

	cursorCmd.Flags().BoolVarP(&setCur, "set", "s", false, "Setting cursor styles")
	cursorCmd.Flags().BoolVarP(&changeCur, "change", "c", false, "Changing cursor styles")
}
