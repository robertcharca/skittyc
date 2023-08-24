package prompts

import (
	"log"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
)

var tabBarSetQuestion = []*survey.Question{
	{
		Name: "setTabBar",
		Prompt: &survey.Select{
			Message: "Choose a configuration group to set in your tab bar: ",
			Options: []string{"tab bar", "active tab", "inactive tab"},
		},
	},
}

var tabBarChangeQuestion = []*survey.Question{
	{
		Name: "changeTabBar",
		Prompt: &survey.Select{
			Message: "Choose an option: ",
			Options: []string{"tab bar", "tab fade", "tab separator", "tab activity symbol", 
				"tab powerline style", "tab switch strategy", "tab title max length"},
		},
	},
}

func HandleSetTabBar() (string, string) {
	tabBarSelect := []*survey.Question{
		{
			Name: "tabBarSelect",
			Prompt: &survey.Select{
				Message: "Choose an option: ",
				Options: []string{"tab bar align", "tab bar edge", "tab bar style", "tab bar background", "tab bar margin color"},
			},
		},
	}

	tabBarActive := []*survey.Question{
		{
			Name: "tabBarActive",
			Prompt: &survey.Select{
				Message: "Choose an option: ",
				Options: []string{"active tab background", "active tab foreground", "active tab font style"},
			},
		},
	}

	tabBarInactive := []*survey.Question{
		{
			Name: "tabBarInactive",
			Prompt: &survey.Select{
				Message: "Choose an option: ",
				Options: []string{"inactive tab background", "inactive tab foreground", "inactive tab font style"},
			},
		},
	}

	answers := struct{	
		Option string `survey:"setTabBar"`
		TbsOption string `survey:"tabBarSelect"`
		TbaOption string `survey:"tabBarActive"`
		TbiOption string `survey:"tabBarInactive"`
	}{}
	
	err := survey.Ask(tabBarSetQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var existingTabBar string
	var tbsValue string

	selectFontText := "Select your font style: "

	switch answers.Option {
	case "tab bar":
		survey.Ask(tabBarSelect, &answers)
		tbsValue := handleTabBarSetOptions(answers.TbsOption)

		if tbsValue == "hex color" {
			optionInput := &survey.Input{
				Message: answers.TbsOption,
			}

			survey.AskOne(optionInput, &existingTabBar, survey.WithValidator(hexCodeValidation.Validate))

			return answers.TbsOption, existingTabBar
		}

		return answers.TbsOption, tbsValue
	case "active tab":
		survey.Ask(tabBarActive, &answers)

		if answers.TbaOption != "active tab font style" {
			tbaValue := inputSurvey(answers.TbaOption, hexCodeValidation.Validate)

			return answers.TbaOption, tbaValue
		}

		tbsValue = selectSurveyOptions(selectFontText, []string{"bold", "normal", "italic", "bold-italic"})

		return answers.TbaOption, tbsValue
	case "inactive tab":
		survey.Ask(tabBarInactive, &answers)

		if answers.TbiOption != "inactive tab font style" {
			tbiValue := inputSurvey(answers.TbiOption, hexCodeValidation.Validate)

			return answers.TbiOption, tbiValue
		}

		tbsValue = selectSurveyOptions(selectFontText, []string{"bold", "normal", "italic", "bold-italic"})	

		return answers.TbiOption, tbsValue
	}	

	return answers.Option, tbsValue 
}

func HandleChangeTabBar() (string, string) {
	tabBarSelect := []*survey.Question{
		{
			Name: "tabBarSelectChange",
			Prompt: &survey.Select{
				Message: "Choose an option: ",
				Options: []string{"tab bar min tabs", "tab bar margin width", "tab bar margin height"},
			},
		},
	}

	answers := struct{ 
		Option string `survey:"changeTabBar"`
		TbcOption string `survey:"tabBarSelectChange"`
	}{}
	
	err := survey.Ask(tabBarChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}
	
	var tbcValue string

	switch answers.Option {
	case "tab bar":
		survey.Ask(tabBarSelect, &answers)

		tbcValue = handleTabBarSetOptions(answers.TbcOption)

		return answers.TbcOption, tbcValue
	case "tab separator", "tab activity symbol":
		tbcValue = inputSurvey(answers.Option, survey.MinLength(1))

		if verify := regexp.MustCompile(`\s`).MatchString(tbcValue); verify {
			return answers.Option, `"` + tbcValue + `"`
		}
	case "tab powerline style":
		tbcValue = selectSurveyOptions("Select your powerline style: ", []string{"angled", "round", "slanted"})	
	case "tab title max length":
		tbcValue = inputSurvey(answers.Option, numberPositiveOnly.Validate)
	case "tab switch strategy":
		tbcValue = selectSurveyOptions("Select your switch strategy: ", []string{"left", "right", "last", "previous"})
	case "tab fade":
		tbcValue = inputSurvey(answers.Option, multiplePositiveNumbers.Validate)
	}

	return answers.Option, tbcValue 
}

func handleTabBarSetOptions(option string) string {
	var tbSelect string

	switch option {
	case "tab bar align":
		tbSelect = selectSurveyOptions("Select your alignment: ", []string{"left", "center", "right"})	
	case "tab bar edge":
		tbSelect = selectSurveyOptions("Select your edge: ", []string{"top", "bottom"})	
	case "tab bar style":
		tbSelect = selectSurveyOptions("Select your tab bar style: ", []string{"fade", "slant", "hidden", "separator", "powerline"})	
	case "tab bar background", "tab bar margin color":
		tbSelect = selectSurveyOptions("Seleect your option: ", []string{"hex color", "none"})	
	case "tab bar margin height":
		tbSelect = inputSurvey(option, multiplePositiveNumbers.Validate)
	case "tab bar min tabs", "tab bar margin width":
		tbSelect = inputSurvey(option, numberPositiveOnly.Validate)
	}

	return tbSelect
}

