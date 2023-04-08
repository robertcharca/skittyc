package setup

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shortcutsCmd = &cobra.Command{
	Use: "shortcuts",	
	Short: "'shortcuts' command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'shortcuts' working")
	},
}

func init() {
	SetupCmd.AddCommand(shortcutsCmd)
}
