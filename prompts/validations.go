package prompts

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

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
		num, _ := number.(string)
		numConv, numOk := strconv.ParseFloat(num, 8)
		if numOk != nil || numConv < 0.0 || numConv > 1.0 {
			return errors.New("This number cannot be less than 0 or greater than 1.")
		}	

		return nil
	},
}

var numberPositiveOnly = &survey.Question{
	Validate: func(number interface{}) error {
		num, _ := number.(string)	
		numConv, okNum := strconv.ParseFloat(num, 8)
		if okNum != nil || numConv > 99.99 {
			return errors.New("This number cannot be less than 0")
		}

		return nil
	},
}

var numberPositiveLarge = &survey.Question{
	Validate: func(number interface{}) error {
		num, _ := number.(string)	
		numConv, okNum := strconv.ParseFloat(num, 8)
		if okNum != nil || numConv < 0.0 {
			return errors.New("This number cannot be less than 0")
		}

		return nil
	},
}

var numberAllRanges = &survey.Question{
	Validate: func(number interface{}) error {
		num, _ := number.(string)
		numConv, okNum := strconv.ParseFloat(num, 8)
		if okNum != nil || numConv < -99.99 || numConv > 99.99 {
			return errors.New("This number cannot be less than -99.99 and greater than 99.99")
		}

		return nil
	},
}

var multiplePositiveNumbers = &survey.Question{
	Validate: func(number interface{}) error {
		num, _ := number.(string)
		spaces := regexp.MustCompile(`\s`).MatchString(num)
		 
		if !spaces {
			numConv, okNum := strconv.ParseFloat(num, 8)
			if okNum != nil || numConv < -99.99 || numConv > 99.99 {
				return errors.New("This number cannot be less than -99.99 and greater than 99.99")
			}
		}

		listNums := strings.Fields(num)

		for ln := 0; ln < len(listNums); ln++ {	
			numConv, okNum := strconv.ParseFloat(listNums[ln], 8)
			if okNum != nil || numConv < -99.99 || numConv > 99.99 {
				return errors.New("These numbers cannot be less than -99.99 and greater than 99.99")
			}
		}

		return nil 
	},
}
