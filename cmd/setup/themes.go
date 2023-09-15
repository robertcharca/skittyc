package setup

import (
	"fmt"
	"log"

	"github.com/robertcharca/skittyc/kittyc"
	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setTheme bool

var themesCmd = &cobra.Command{
	Use: "themes",	
	Short: "'themes' command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setTheme {
			res := prompts.InputSetTheme()
			fmt.Printf("res: %s\n", res)

			handleKittyConfTheme(res)
		}
	},
}

func init() {
	SetupCmd.AddCommand(themesCmd)
	themesCmd.Flags().BoolVarP(&setTheme, "set", "s", false, "Setting themes.")
}

func handleKittyConfTheme(path string) {
	option, exists := prompts.ConfirmKittyConfExistence()	

	if !exists{
		kittyc.CreateKittyConf()
		filePath := handleSetupDownloads(path)
		if err := kfeatures.ReplacingKittyConf(filePath); err != nil {	
			log.Fatalln(err)
		}
	}

	switch option {
	case "save it as a profile":	
		profileName := prompts.ProfileNameInput()
		filePath := handleSetupDownloads(path)
		if err := kfeatures.SavingKittyFileProfile(filePath, profileName); err != nil {
			log.Fatalln(err)
		} 
	case "replace it":
		filePath := handleSetupDownloads(path)
		if err := kfeatures.ReplacingKittyConf(filePath); err != nil {	
			log.Fatalln(err)
		}	
	default:
		fmt.Println("Quitting...")
	}	
}
