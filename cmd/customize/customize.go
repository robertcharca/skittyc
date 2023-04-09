package customize

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	//Global flags for the "customize" command.
	setting bool	
)

var CustomizeCmd = &cobra.Command{
	Use: "customize",
	Short: "Welcome to 'customize'.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Section for customizing your Kitty terminal.")
	},	
}

func init() {
	//PersistentFlags(): global flags for your command.
	CustomizeCmd.PersistentFlags().BoolVarP(&setting, "set", "s", false, "Setting a configuration")	
	
	CustomizeCmd.MarkPersistentFlagRequired("set")
}
