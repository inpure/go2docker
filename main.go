package main

import (
	"fmt"
	"github.com/inpure/go2docker/cmd"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("process exited with errors: \n%s\n", err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Bye!")
}
