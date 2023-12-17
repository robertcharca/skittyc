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

	fileKittyProfile, err := EmptyKittyProfile(customName)
	if err != nil {
		return err
	}

	if err := rewrittingKittyFile(fileKittyProfile, kittyThemeRepl); err != nil {
		return err
	}

	return nil
}

func EmptyKittyProfile(customName string) (string, error) {
	newFileName := customName + ".conf"
	reduceSpace := strings.ReplaceAll(newFileName, " ", "")
	emptyName := strings.HasPrefix(reduceSpace, ".")

	if emptyName {
		return "", errors.New("a file with no name is not accepted")
	}

	emptyProfile := kittyProfile + reduceSpace

	// Verifying if the "profile" directory exists. If not, then we create it
	err := os.MkdirAll(kittyProfile, 0750)
	if err != nil {
		return "", err
	}

	// Creating a empty kitty file
	file, err := os.Create(emptyProfile)
	if err != nil {
		return "", err
	}

	kittyc.DisplayStructure(file)

	return emptyProfile, nil
}

func SavingKittyProfileChanges(path string) error {
	kittyConfChanges, err := filteringKittyTheme(kittyPath)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if err := rewrittingKittyFile(path, kittyConfChanges); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
