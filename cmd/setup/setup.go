package setup

import (
	"fmt"
	"strings"

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/spf13/cobra"
)

var SetupCmd = &cobra.Command{
	Use: "setup",	
	Short: "Welcome to 'setup'.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Section for setting up your principal config to the Kitty terminal.")
	},
}

func init() {
	//Funcionality for the 'setup' command.
}

func handleSetupDownloads(path string) string {
	var newPath string

	resultPath := strings.HasPrefix(path, "/")

	if !resultPath {
		newPath = kfeatures.DownloadKittyFiles(path, ".conf")		
	} else {
		newPath = path
	}

	return newPath
}
