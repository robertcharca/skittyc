package prompts

import (
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

func HandleSetTheme() string {
	homePath, _ := os.UserHomeDir()

	setThemeQuestion := &survey.Input{
		Message: "Paste a url or .conf/.txt file path of your theme:",
		Suggest: func(toComplete string) []string {
			var filteredFiles []string

		confFiles, _ := filepath.Glob(homePath + toComplete + "*.conf")
		textFiles, _ := filepath.Glob(homePath + toComplete + "*.txt")
		confFiles = append(confFiles, textFiles...)
		filteredFiles = append(filteredFiles, confFiles...)

		return filteredFiles
		},
	}	

	var themeValue string	

	if err := survey.AskOne(setThemeQuestion, &themeValue); err != nil {
		log.Fatalln(err)
	}

	return themeValue
}
