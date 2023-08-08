package prompts

import (
	"errors"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/robertcharca/skittyc/kittyc"
)

var urlValidation = &survey.Question{	
	Validate: func (url interface{}) error {
		if link, ok := url.(string); !ok || len(link) < 11 {
			return errors.New("This link cannot be less than 11 characters.")
		}
		return nil
	},
}

var hexCodeValidation = &survey.Question{
	Validate: func(val interface{}) error {
		code, ok := val.(string)
		codeList := kittyc.ConvertStringToList(code)
		if !ok || codeList[0] != "#" || len(codeList[1:]) > 6{
			return errors.New("Your code should start with '#' and be less than 7")
		} 

		return nil
	},
}

var numberZeroToOneValidator = &survey.Question{
	Validate: func(number interface{}) error {
		num, ok := number.(string)
		numConv, _ := strconv.ParseFloat(num, 8)
		if !ok || numConv < 0.0 || numConv > 1.0 {
			return errors.New("This number cannot be less than 0 or greater than 1.")
		}	

		return nil
	},
}
