package kittyc

import (
	"errors"
	"fmt"
	"os"
)

func kittyConfigExistence() (error, bool) {
	//Checking if $HOME variable is set 
	homePath, errHomePath := os.UserHomeDir()
	
	if errHomePath != nil {
		return errHomePath, false
	}

	//Declaring and initializing the path for 'kitty,conf' file
	/*
		os.Stat() only recognizes a path that doesn't have any type of variable,
		like $HOME, $PATH and similar types.
	*/
	var path string = homePath + "/.config/kitty/kitty.conf" 

	_, err := os.Stat(path) 

	//Checking if the file 'kitty.conf' does exist.
	if errors.Is(err, os.ErrNotExist) {
		return err, false
	}

	return nil, true
}

func CreateKittyConf(){

	fmt.Println("Verifying if there's a 'kitty.conf' file created...")

	_, kittyConfExistence := kittyConfigExistence()	
	
	if kittyConfExistence == false {
		file, err := os.Create("kitty.conf")	

		//Handling errors when it's creating a 'kitty.conf' file.
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer file.Close()
	} else {
		fmt.Println("Ops! There's a file created")
	}
}

func Kitty() {
	// - Create a function to create kitty.conf file. -checked
	// - Verify if kitty.conf file exists for creating automatically. -checked
	// - Create a function for writing and verifying if a config line exists.
	// - Watch how you can structure kitty configuration files through functions or structs.
	// Idea: verify if it's possible to add a flag that contains an argument like --set=default
	fmt.Println("I'm working")

}
