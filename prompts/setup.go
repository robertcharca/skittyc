package prompts

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func InputSetTheme() string {
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

func listAvailableProfiles() []string {
	homePath, _ := os.UserHomeDir()
	profilePath := homePath + "/.config/kitty/profiles/"

	var profiles []string

	profs, _ := ioutil.ReadDir(profilePath)	

	for _, p := range profs {
		noFormat := strings.ReplaceAll(p.Name(), ".conf", "")
		profiles = append(profiles, noFormat)
	}

	return profiles
}

func ProfileNameInput() string {
	newFileName := inputSurvey("Type your profile name: ", survey.MinLength(1))

	return newFileName
}

func HandleSetProfile() (string, string) {
	var path string
	var fileName string

	profileSet := selectSurveyOptions("How would you like to set your new profile? ", []string{"empty file", "implement a theme"})

	switch profileSet {
	case "empty file":
		fileName = ProfileNameInput() 
	case "implement a theme":
		path = InputSetTheme()
		fileName = ProfileNameInput()
	}

	return path, fileName 
}

func HandleChangeProfile() (string, string) {
	var selection string

	homePath, _ := os.UserHomeDir()
	profilePath := homePath + "/.config/kitty/profiles/"

	profiles := listAvailableProfiles()

	profileChange := []*survey.Question{
		{
			Name: "changeProfile",
			Prompt: &survey.Select{
				Message: "Choose your option",
				Options: []string{"save changes", "change profile"},
			},
		},
	}

	answers := struct { Option string `survey:"changeProfile"` }{}

	err := survey.Ask(profileChange, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	switch answers.Option {
	case "save changes":
		profiles = append(profiles, "new profile")
		selection = selectSurveyOptions("In which profile? ", profiles)
	case "change profile":
		selection = selectSurveyOptions("Select your profile: ", profiles)
	}

	return answers.Option, profilePath + selection + ".conf"
}
