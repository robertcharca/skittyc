package customize

import (
	"fmt"
	"strings"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setC bool

var colorCmd = &cobra.Command{
	Use: "colors",
	Short: "'colors' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setC {
			prompt, res := prompts.HandleSetColor()
	
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
		colors := kfeatures.DownloadKittyFiles(path, ".conf")
		kfeatures.SetColors(colors)	
	} else {
		kfeatures.SetColors(path)	
	}
}
