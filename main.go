package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("Hello World")
		time.Sleep(2 * time.Second)
	}
}
