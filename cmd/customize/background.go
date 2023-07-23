package customize

import (
	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var (
	setBF bool
	changeBF bool
)

var backforeCmd = &cobra.Command{
	Use: "backfore",
	Short: "'background' and 'foreground'subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if changeBF == true {
			prompt, res := prompts.HandleChangeBackfore()
			kfeatures.ChangingValues(prompt, res, "# Background and Foreground")
		}

		if setBF == true {
			prompt, res := prompts.HandleSetBackfore()
			kfeatures.ChangingValues(prompt, res, "# Background and Foreground")
		}
	},
}

func init() {	
	// Adding the command `fonts` to `customize`.
	CustomizeCmd.AddCommand(backforeCmd)
		
	// Flag for changing font values (size, bold, italic)
	backforeCmd.Flags().BoolVarP(&changeBF, "change", "c", false, "Changing background styles.")
	// Flag for setting a new font 
	backforeCmd.Flags().BoolVarP(&setBF, "set", "s", false, "Set backgroud styles.")	
}
