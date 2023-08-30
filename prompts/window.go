package prompts

import (
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

var windowSetNewQuestion = []*survey.Question{
	{
		Name: "setWindow",
		Prompt: &survey.Select{
			Message: "What would you like to set?",
				Options: []string{"window", "bell border color", "inactive text alpha", 
				"active border color", "inactive border color"},
		},
	},
}

var windowChangeQuestion = []*survey.Question{
	{
		Name: "changeWindow",
		Prompt: &survey.Select{
			Message: "Choose an option: ",
			Options: []string{"window", "enabled layouts", "placement strategy", "draw minimal borders"},	
		},
	},
}

func HandleSetWindow() (string, string) {
	windowSelectQuestion := []*survey.Question{
		{
			Name: "selectWindowSet",
			Prompt: &survey.Select{
				Message: "Choose an option: ",
				Options: []string{"window border width", "window margin width", "initial window width", "initial window height", 
					"window padding width", "single window margin width"},
			},
		},
	}

	answers := struct{	
		Option string `survey:"setWindow"`
		WSetOption string `survey:"selectWindowSet"`
	}{}

	err := survey.Ask(windowSetNewQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var windowValue string

	switch answers.Option {
	case "window":
		survey.Ask(windowSelectQuestion, &answers)
		
		windowValue = handleWindowOptions(answers.WSetOption)
		
		return answers.WSetOption, windowValue
	case "inactive text alpha":
		windowValue = inputSurvey(answers.Option, numberZeroToOneValidator.Validate)
	case "bell border color", "active border color", "inactive border color":
		windowValue = inputSurvey(answers.Option, hexCodeValidation.Validate)
	}

	return answers.Option, windowValue 
}

func HandleChangeWindow() (string, string) {
	windowSelectQuestion := []*survey.Question{
		{
			Name: "selectWindowChange",
			Prompt: &survey.Select{
				Message: "Choose an option: ",
				Options: []string{"window logo alpha", "window logo path", "window logo position", 
					"remember window size", "hide window decorations", "confirm os window close"},
			},
		},
	}

	answers := struct{	
		Option string `survey:"changeWindow"`
		WChangeOption string `survey:"selectWindowChange"`
	}{}

	err := survey.Ask(windowChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var windowValue string

	switch answers.Option {
	case "window":
		survey.Ask(windowSelectQuestion, &answers)

		windowValue = handleWindowOptions(answers.WChangeOption)
		
		return answers.WChangeOption, windowValue
	case "enabled layouts":
		windowValue = selectSurveyOptions("Select an enabled layout: ", []string{"all", "fat", "tall", "grid", "stack", "splits", "vertical", "horizontal"})	
	case "placement strategy":
		windowValue = selectSurveyOptions("Select a placement strategy: ", []string{"center", "top-left"})
	case "draw minimal borders":
		windowValue = selectSurveyOptions("Select an option: ", []string{"yes", "no"})
	}

	return answers.Option, windowValue
}

func handleWindowOptions(option string) string {
	var windowValue string

	homePath, _ := os.UserHomeDir()

	windowSetImage := &survey.Input{
		Message: "Type your png image path: ",
		Suggest: func (toComplete string) []string {
			pngFiles, _ := filepath.Glob(homePath + toComplete + "*.png")
			return pngFiles
		},
	}
	
	switch option {
	case "window margin width", "window padding width", "single window margin width":
		windowValue = inputSurvey(option, multiplePositiveNumbers.Validate)
	case "window border width", "initial window width", "initial window height":
		windowValue = inputSurvey(option, numberPositiveLarge.Validate)
	case "hide window decorations", "remember window size":
		windowValue = selectSurveyOptions("Select an option: ", []string{"yes", "no"})
	case "confirm os window close":
		windowValue = inputSurvey(option, numberAllRanges.Validate)
	case "window logo position":
		windowValue = selectSurveyOptions("Select an logo position: ", []string{"top", "left", "right", "center", "bottom", 
			"top-left", "top-right", "bottom-left", "bottom-right"})
	case "window logo alpha":
		windowValue = inputSurvey(option, numberZeroToOneValidator.Validate)
	case "window logo path":
		survey.AskOne(windowSetImage, &windowValue)
	}

	return windowValue
}
