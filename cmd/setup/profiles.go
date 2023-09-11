package setup

import (
	"fmt"

	"github.com/spf13/cobra"
)

var profilesCmd = &cobra.Command{
	Use: "profiles",	
	Short: "'profiles' command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'profiles' working")
	},
}

func init() {
	SetupCmd.AddCommand(profilesCmd)
}
