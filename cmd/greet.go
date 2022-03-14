package cmd

import (
	"github.com/inpure/go2docker/pkg/hello"
	"github.com/spf13/cobra"
)

var pause int
var name string

var cmdGreet = &cobra.Command{
	Use:   "greet",
	Short: "name",
	Long:  "greet name",
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		go hello.LoopSay(name, pause)
	},
}

func init() {
	cmdGreet.Flags().IntVar(&pause, "p", 1, "pause time，default 1s")
	cmdGreet.Flags().StringVar(&name, "n", "World", "name，default \"World\"")
}
