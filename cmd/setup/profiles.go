package setup

import (
	"log"	

	"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/prompts"
	"github.com/spf13/cobra"
)

var setProfile, changeProfile bool

var profilesCmd = &cobra.Command{
	Use: "profiles",	
	Short: "'profiles' command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setProfile {
			path, name := prompts.HandleSetProfile()

			if path != "" {
				filePath := handleSetupDownloads(path)
				if err := kfeatures.SavingKittyFileProfile(filePath, name); err != nil {
					log.Fatalln(err)
				}
			} else {
				kfeatures.EmptyKittyProfile(name)
			}
		}

		if changeProfile {
			prompt, profile := prompts.HandleChangeProfile()
			
			switch prompt {
			case "save changes":	
				kfeatures.SavingKittyProfileChanges(profile)	
			case "change profile":
				kfeatures.ReplacingKittyConf(profile)
			}
		} 
	},
}

func init() {
	SetupCmd.AddCommand(profilesCmd)
	profilesCmd.Flags().BoolVarP(&setProfile, "set", "s", false, "Setting profiles.")
	profilesCmd.Flags().BoolVarP(&changeProfile, "change", "c", false, "Change profiles.")
}
