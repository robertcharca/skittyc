package customize

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	//"github.com/robertcharca/skittyc/kittyc/kfeatures"
	"github.com/robertcharca/skittyc/internal/tui"
)

var (
	//Local flags
	setting bool	
	changing bool
)

var fontCmd = &cobra.Command{
	Use: "fonts",
	Short: "'fonts' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'fonts' working")
		
		if changing == true {
			fmt.Println("Change")
		}

		if setting == true {
			fmt.Println("Set")
		}

		if err := tea.NewProgram(tui.InitialModel()).Start(); err != nil {
			log.Fatal(err)	
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
	// Flag for changing font values (size, bold, italic)
	fontCmd.Flags().BoolVarP(&changing, "change", "c", false, "Changing font values.")
	// Flag for setting a new font 
	fontCmd.Flags().BoolVarP(&setting, "set", "s", false, "Setting a new font.")	
}
