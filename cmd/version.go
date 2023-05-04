package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version 1.1",
	Short: "Print the version number of dk",
	Long:  `All software has versions. This is Gopher`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
