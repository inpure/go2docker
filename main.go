package main

import (
	"github.com/spf13/cobra"
	"log"
	"time"
)

func main() {
	var pause int
	var times int
	var name string

	//flag.IntVar(&pause, "p", 1, "pause time，default 1s")
	//flag.IntVar(&times, "t", 1, "output times，default 1")
	//flag.StringVar(&name, "n", "World", "name，default \"World\"")
	//flag.Parse() //parse the command line arguments

	var cmdGreet = &cobra.Command{
		Use:   "greet",
		Short: "name",
		Long:  "greet name",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < times; i++ {
				log.Println("Hello " + name)
				if i == times-1 {
					break
				}
				time.Sleep(time.Duration(pause) * time.Second)
			}
		},
	}

	cmdGreet.Flags().IntVar(&pause, "p", 1, "pause time，default 1s")
	cmdGreet.Flags().IntVar(&times, "t", 1, "output times，default 1")
	cmdGreet.Flags().StringVar(&name, "n", "World", "name，default \"World\"")

	var rootCmd = &cobra.Command{Use: "hello"}
	rootCmd.AddCommand(cmdGreet)
	rootCmd.Execute()
}
