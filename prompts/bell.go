package prompts

import (
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

var bellSetQuestion = []*survey.Question{
	{
		Name: "setBell",
		Prompt: &survey.Select{
			Message: "What would you like to set in your bell?",
			Options: []string{"enable audio bell", "visual bell color", "window alert on bell"},
		},
	},
}

var bellChangeQuestion = []*survey.Question{
	{
		Name: "changeBell",
		Prompt: &survey.Select{
			Message: "Choose an option: ",
			Options: []string{"bell path", "bell on tab", "visual bell duration"},
		},
	},
}

func HandleSetBell() (string, string) {
	answers := struct{ Option string `survey:"setBell"`}{}
	
	err := survey.Ask(bellSetQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var existingBell string

	bellConfirmationSelect := &survey.Select{
		Message: "Select your confirmation: ",
		Options: []string{"yes", "no"},
	}

	bellColorSelect := &survey.Select{
		Message: "Select your option: ",
		Options: []string{"hex color", "none"},
	}

	bellColorInput := &survey.Input{
		Message: answers.Option,
	}

	switch answers.Option {
	case "visual bell color":
		survey.AskOne(bellColorSelect, &existingBell)

		if existingBell != "none" {
			survey.AskOne(bellColorInput, &existingBell, survey.WithValidator(hexCodeValidation.Validate)) 
		}	
	default:
		survey.AskOne(bellConfirmationSelect, &existingBell)
	}  
	
	return answers.Option, existingBell
}

func HandleChangeBell() (string, string) {
	answers := struct{ Option string `survey:"changeBell"`}{}
	homePath, _ := os.UserHomeDir()

	err := survey.Ask(bellChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var existingBell string

	bellPathInput := &survey.Input{
		Message: "Paste a file path for your bell song:",
		Suggest: func(toComplete string) []string {
			var filteredFiles []string

			confFiles, _ := filepath.Glob(homePath + toComplete + "*.wav")
			textFiles, _ := filepath.Glob(homePath + toComplete + "*.oga")
			confFiles = append(confFiles, textFiles...)
			filteredFiles = append(filteredFiles, confFiles...)

			return filteredFiles
		},
	}

	bellDurationInput := &survey.Input{
		Message: answers.Option,
	}

	bellConfirmationSelect := &survey.Select{
		Message: "Select your confirmation: ",
		Options: []string{"yes", "no", "custom"},
	}

	bellOnTabInput := &survey.Input{
		Message: "Type a word or an emoji:",
	}
	
	switch answers.Option{
	case "bell path":
		survey.AskOne(bellPathInput, &existingBell, survey.WithValidator(survey.MinLength(6))) 
	case "visual bell duration":	
		survey.AskOne(bellDurationInput, &existingBell, survey.WithValidator(numberPositiveOnly.Validate)) 
	default:
		survey.AskOne(bellConfirmationSelect, &existingBell)

		if existingBell == "custom" {
			survey.AskOne(bellOnTabInput, &existingBell)
			return answers.Option, `"` + existingBell + ` "`
		}
	}

	return answers.Option, existingBell
}
