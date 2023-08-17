package customize

import (
	"fmt"
	"strings"

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
			
			fmt.Printf("prompt: %s, res: %s\n", prompt, res)	
			switch prompt {
			case "color scheme":
				handleColorPath(res)	
			default:
				err := kfeatures.ChangingValues(prompt, res, "# Colors")	
				if err != nil {
					fmt.Println(err.Error())
				}
			}	
		}
	},
}

func init() {
	CustomizeCmd.AddCommand(colorCmd)	

	colorCmd.Flags().BoolVarP(&setC, "set", "s", false, "Setting color styles.")
}

func handleColorPath(path string) {
	resultPath := strings.HasPrefix(path, "/")

	if !resultPath {
		colors := kfeatures.DownloadColors(path)
		kfeatures.SetColors(colors)	
	} else {
		kfeatures.SetColors(path)	
	}
}
