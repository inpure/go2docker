package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{Use: "hello"}

func init() {
	rootCmd.AddCommand(cmdGreet)
	rootCmd.AddCommand(cmdServer)
}

func Execute() error {
	return rootCmd.Execute()
}
