package kittyc

import (
	"errors"
	"fmt"	
	"os"
)

var homePath, errHomePath = os.UserHomeDir()

var path string = homePath + "/.config/kitty/kitty.conf" 

func KittyConfigExistence() (error, bool) {
	//Checking if $HOME variable is set 
	if errHomePath != nil {
		return errHomePath, false
	}

	//Declaring and initializing the path for 'kitty,conf' file
	/*
		os.Stat() only recognizes a path that doesn't have any type of variable,
		like $HOME, $PATH and similar types.
	*/
	_, err := os.Stat(path) 

	//Checking if the file 'kitty.conf' does exist.
	if errors.Is(err, os.ErrNotExist) {
		return err, false
	}

	return nil, true
}

func CreateKittyConf(){

	fmt.Println("Verifying if there's a 'kitty.conf' file created...")

	_, kittyConfExistence := KittyConfigExistence()	

	//kitty.conf file path
	var kittyConfPath string = path
	
	if kittyConfExistence == false {
		file, err := os.Create(kittyConfPath)	

		//Handling errors when it's creating a 'kitty.conf' file.
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer file.Close()
		
		displayStructure(file)

		fmt.Println("kitty.conf has been created!")
	} else {
		fmt.Println("Ops! There's a file created")
	}
}

func Kitty() {	
	fmt.Println("I'm working")
}
