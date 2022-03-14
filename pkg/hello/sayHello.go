package hello

import (
	"log"
	"time"
)

func LoopSay(name string, pause int) {
	for {
		log.Println("Hello " + name)
		time.Sleep(time.Duration(pause) * time.Second)
	}
}
