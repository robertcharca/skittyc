package customize

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fontCmd = &cobra.Command{
	Use: "font",
	Short: "'font' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'font' working")
	},
}

func init() {
	//
	CustomizeCmd.AddCommand(fontCmd)
}
