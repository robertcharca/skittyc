package customize

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	//Local flags
	customFonts bool
	defaultFonts bool
)

var fontCmd = &cobra.Command{
	Use: "fonts",
	Short: "'fonts' subcommand",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'fonts' working")
	},
}

func init() {	
	//Adding the command `fonts` to `customize`.
	CustomizeCmd.AddCommand(fontCmd)
	
	//These flags are for setting a default or custom configuration in `fonts`.
	fontCmd.Flags().BoolVarP(&customFonts, "custom", "c", false, "Custom setting")
	fontCmd.Flags().BoolVarP(&defaultFonts, "default", "d", false, "Default setting")

	//XOR flag execution.
	fontCmd.MarkFlagsMutuallyExclusive("custom", "default")
	
}
