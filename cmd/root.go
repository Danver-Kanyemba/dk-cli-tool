package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{Use: "dk    CLi 1.0"}
	rootCmd.AddCommand(mkpkgCmd, mkBase, mkPage, versionCmd)
	rootCmd.Execute()
}
