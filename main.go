package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var pause int
	var name string

	//flag.IntVar(&pause, "p", 1, "pause time，default 1s")
	//flag.IntVar(&times, "t", 1, "output times，default 1")
	//flag.StringVar(&name, "n", "World", "name，default \"World\"")
	//flag.Parse() //parse the command line arguments

	var cmdGreet = &cobra.Command{
		Use:   "greet",
		Short: "name",
		Long:  "greet name",
		//Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				for {
					log.Println("Hello " + name)
					time.Sleep(time.Duration(pause) * time.Second)
				}
			}()
		},
	}

	cmdGreet.Flags().IntVar(&pause, "p", 1, "pause time，default 1s")
	cmdGreet.Flags().StringVar(&name, "n", "World", "name，default \"World\"")

	var rootCmd = &cobra.Command{Use: "hello"}
	rootCmd.AddCommand(cmdGreet)
	rootCmd.Execute()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Bye!")

}
