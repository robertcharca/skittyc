package prompts

import (	
	"log"
	"os"
	"path/filepath"	

	"github.com/AlecAivazis/survey/v2"	
)
var backforeSetNewQuestion = []*survey.Question{
	{
		Name: "setBackfore",
		Prompt: &survey.Select{
			Message: "What would you like to set?",
			Options: []string{"background", "foreground", "background image", 
				"selection background", "selection foreground"},	
		},
	},
}

var backforeChangeQuestion = []*survey.Question{
	{
		Name: "changeBackfore",
		Prompt: &survey.Select{
			Message: "Choose an option: ",
			Options: []string{"background tint", "background opacity", "background image layout"},
			Help: "helping",
		},
	},
}

func HandleSetBackfore() (string, string) {
	homePath, _ := os.UserHomeDir()
	answers := struct{ Option string `survey:"setBackfore"`}{}

	err := survey.Ask(backforeSetNewQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var existingBackfore string
	var backforeValue string	

	backforeSetImage := &survey.Input{
		Message: "Type your png image path: ",
		Suggest: func (toComplete string) []string {
			pngFiles, _ := filepath.Glob(homePath + toComplete + "*.png")
			return pngFiles
		},
	}	

	switch answers.Option {
	case "background image":	
		survey.AskOne(backforeSetImage, &existingBackfore)	

		return answers.Option, existingBackfore
	default:
		backforeValue = inputSurvey(answers.Option, hexCodeValidation.Validate)
	}	
	
	return answers.Option, backforeValue 
}

func HandleChangeBackfore() (string, string) {	
	answers := struct{ Option string `survey:"changeBackfore"`}{}
	err := survey.Ask(backforeChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}
	
	var backforeValue string	

	switch answers.Option {
	case "background image layout":
		backforeValue = selectSurveyOptions("Select your image layout: ", []string{"tiled", "mirror-tiled", "scaled", "clamped", "centered", "cscaled"})
	default:
		backforeValue = inputSurvey(answers.Option, numberZeroToOneValidator.Validate)
	}	
	
	return answers.Option, backforeValue 
}
