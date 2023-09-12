package kfeatures

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
)

func SavingKittyFileProfile(path, customName string) error {
	kittyThemeRepl, errTheme := filteringKittyTheme(path)
	if errTheme != nil {
		log.Fatalln(errTheme)
		return errTheme
	}

	newFileName := customName + ".conf"
	reduceSpace := strings.ReplaceAll(newFileName, " ", "") 
	emptyName := strings.HasPrefix(reduceSpace, "."); 
	
	if emptyName {
		return errors.New("A file with no name is not accepted")
	}

	fileKittyProfile := kittyProfile + reduceSpace 

	// Verifying if the "profile" directory exists. If not, then we create it
	err := os.MkdirAll(kittyProfile, 0750)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	// Creating a empty kitty file
	file, err := os.Create(fileKittyProfile)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	kittyc.DisplayStructure(file)

	if err := rewrittingKittyFile(fileKittyProfile, kittyThemeRepl); err != nil {
		return err
	}

	return nil
}
