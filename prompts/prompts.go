package prompts

import "github.com/AlecAivazis/survey/v2"

func selectSurveyOptions(message string, options []string) string {
	var value string

	prompt := &survey.Select{
		Message: message,
		Options: options,
	}

	survey.AskOne(prompt, &value)

	return value
}

func inputSurvey(message string, validation survey.Validator) string {
	var value string

	prompt := &survey.Input{
		Message: message,
	}

	survey.AskOne(prompt, &value, survey.WithValidator(validation))

	return value
}
