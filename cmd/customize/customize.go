package customize

import (
	"fmt"

	"github.com/spf13/cobra"
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
	//	
}
