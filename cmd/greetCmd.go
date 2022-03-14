package cmd

import (
	"github.com/inpure/go2docker/pkg/funs"
	"github.com/spf13/cobra"
)

var cmdGreet = &cobra.Command{
	Use:   "greet",
	Short: "name",
	Long:  "greet name",
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		go funs.LoopSay(name, interval)
	},
}

func init() {
	cmdGreet.Flags().IntVar(&interval, "i", 1, "interval time，default 1s")
	cmdGreet.Flags().StringVar(&name, "n", "World", "name，default \"World\"")
}
