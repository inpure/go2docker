package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{Use: "hello"}

func init() {
	rootCmd.AddCommand(cmdGreet)
}

func Execute() error {
	return rootCmd.Execute()
}
