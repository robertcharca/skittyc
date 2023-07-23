package prompts

import (	
	"log"
	"os"
	"path/filepath"	

	"github.com/AlecAivazis/survey/v2"	
)

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

var backforeSetNewQuestion = []*survey.Question{
	{
		Name: "setBackfore",
		Prompt: &survey.Select{
			Message: "What would you like to set?",
			Options: []string{"background", "foreground", "background image", "selection background", "selection foreground"},	
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

	backforeSetColor := &survey.Input{
		Message: answers.Option,	 
	}

	backforeSetImage := &survey.Input{
		Message: "Type your png image path: ",
		Suggest: func (toComplete string) []string {
			pngFiles, _ := filepath.Glob(homePath + toComplete + "*.png")
			return pngFiles
		},
	}	

	if answers.Option != "background image" {
		bfInput := survey.AskOne(backforeSetColor, &existingBackfore, survey.WithValidator(hexCodeValidation.Validate), survey.WithValidator(survey.MaxLength(7)))
		if bfInput != nil {
			log.Fatalln(bfInput)
			return " ", " "
		}

		return answers.Option, existingBackfore
	}

	bfInputImage := survey.AskOne(backforeSetImage, &existingBackfore)
	if bfInputImage != nil {
		log.Fatalln(bfInputImage)
		return " ", " "
	}
	
	return answers.Option, existingBackfore
}

func HandleChangeBackfore() (string, string) {	
	answers := struct{ Option string `survey:"changeBackfore"`}{}
	err := survey.Ask(backforeChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var existingBackfore string

	backforeChangeBackground := &survey.Input{
		Message: answers.Option,
	}

	backforeChangeImage := &survey.Select{
		Message: "Select your image layout: ",
		Options: []string{"tiled", "mirror-tiled", "scaled", "clamped", "centered", "cscaled"},
	}

	if answers.Option != "background image layout" {
		bfInput := survey.AskOne(backforeChangeBackground, &existingBackfore, survey.WithValidator(numberZeroToOneValidator.Validate))
		if bfInput != nil {
			log.Fatalln(bfInput)
			return " ", " "
		}

		return answers.Option, existingBackfore
	}

	bfInputImage := survey.AskOne(backforeChangeImage, &existingBackfore)
	if bfInputImage != nil {
		log.Fatalln(bfInputImage)
		return "", ""
	}
	
	return answers.Option, existingBackfore
}
