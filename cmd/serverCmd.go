package cmd

import (
	"github.com/inpure/go2docker/pkg/funs"
	"github.com/spf13/cobra"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "a server",
	Run: func(cmd *cobra.Command, args []string) {
		go funs.Server(addr)
	},
}

func init() {
	cmdServer.Flags().StringVar(&addr, "a", "127.0.0.1:8001", "listen address，default 127.0.0.1:8001")
}
