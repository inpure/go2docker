package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"net"
)

var addr string

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "a server",
	Run: func(cmd *cobra.Command, args []string) {
		serv, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalln(err)
		}
		defer serv.Close()
	},
}

func init() {
	cmdServer.Flags().StringVar(&addr, "a", "127.0.0.1:8001", "listen address，default 127.0.0.1:8001")
}
