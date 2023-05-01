package customize

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/robertcharca/skittyc/kittyc/kfeatures"
)

var (
	//Local flags
	setting string	
)

var fontCmd = &cobra.Command{
	Use: "fonts",
	Short: "'fonts' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'fonts' working")
		
		if setting == "default" {
			kfeatures.AddingDefaultFonts("appending file 2")	
		} else if setting == "customized" {
			fmt.Println("Customized fonts")
		} else {
			fmt.Println("Error")
		}

	},
}

func init() {	
	//Adding the command `fonts` to `customize`.
	CustomizeCmd.AddCommand(fontCmd)
	
	//Flags
	/*
		An argument is called using the flag variable and setting your command variable
		with the "Flags()" method. 
	*/
	fontCmd.Flags().StringVarP(&setting, "set", "s", "", "Setting font configuration.")
	fontCmd.MarkFlagRequired("set")
}
