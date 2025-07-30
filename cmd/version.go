package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version will be set by the linker during build
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Display the current version of glm-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("glm-cli version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
