package prompts

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

var colorSetQuestion = []*survey.Question{
	{
		Name: "setColor",
		Prompt: &survey.Select{
			Message: "How would you like to set your terminal colors?",
			Options: []string{"color scheme", "select color"},	
		},
	},
}

var colorScheme = []string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}

func HandleSetColor() (string, string) {
	answers := struct{ Option string `survey:"setColor"`}{}

	err := survey.Ask(colorSetQuestion, &answers)
	if err != nil {
		log.Fatalln(err)	
	}

	var existingColor string	
	
	colorSchemeUrlDownload := &survey.Input{
		Message: "Paste a url for implementing your color scheme:",
	}
	
	colorSchemeListMultiple := &survey.Select{
		Message: "Select your color:",
		Options: colorScheme,
	}

	switch answers.Option {
	case "color scheme":
		survey.AskOne(colorSchemeUrlDownload, &existingColor, survey.WithValidator(urlValidation.Validate))
	case "select color":
		survey.AskOne(colorSchemeListMultiple, &existingColor)

		colorName := handleColorCode(existingColor)

		colorInput := &survey.Input{
			Message: colorName,
		}
		
		survey.AskOne(colorInput, &existingColor, survey.WithValidator(hexCodeValidation.Validate))
		
		return colorName, existingColor
	default:
		fmt.Println("Select a valid option")
	}

	return answers.Option, ""
}

func handleColorCode(color string) string {
	var colorCodeList []string
	
	number := 0
	name := "color"

	for i := 0; i < len(colorScheme); i++ {
		number = i

		if color == colorScheme[i] {
			firstNum := strconv.Itoa(number)
			secondNum := strconv.Itoa(number + 8)
			colorCodeList = append(colorCodeList, name + firstNum, name + secondNum)

			break
		}

	}	

	selectedColor := []*survey.Question{
		{
			Name: "colorCode",
			Prompt: &survey.Select{
				Message: "Select one of the color codes:",
				Options: colorCodeList,
			},
		},
	}

	colorAnswer := struct{ Option string `survey:"colorCode"`}{}

	err := survey.Ask(selectedColor, &colorAnswer)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return colorAnswer.Option
}
