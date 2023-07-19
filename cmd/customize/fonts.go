package customize

import (
	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var (	
	setting bool	
	changing bool
)

var fontCmd = &cobra.Command{
	Use: "fonts",
	Short: "'fonts' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if changing == true {
			prompt, res := prompts.HandleFontChangeValues()	
			kfeatures.ChangingFontValues(prompt, res)
		}

		if setting == true {
			promt, res := prompts.HandleNewFont()		

			switch promt {
			case "automatic":
				unknownFont, download := prompts.ConfirmFontExistence(res)
				
				if unknownFont == false && download == false {	
					kfeatures.SetNewFont(res)
				} else if unknownFont == true && download == true{	
					font := kfeatures.DownloadNewFont(res)
					kfeatures.SetFontComparing(font)
				} 
			case "url":	
				font := kfeatures.DownloadNewFont(res)	
				kfeatures.SetFontComparing(font)
			case "select from system":
				kfeatures.SetNewFont(res)
			}
		}			
	},
}

func init() {	
	// Adding the command `fonts` to `customize`.
	CustomizeCmd.AddCommand(fontCmd)
		
	// Flag for changing font values (size, bold, italic)
	fontCmd.Flags().BoolVarP(&changing, "change", "c", false, "Changing font values.")
	// Flag for setting a new font 
	fontCmd.Flags().BoolVarP(&setting, "set", "s", false, "Setting a new font.")	
}
