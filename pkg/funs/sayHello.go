package funs

import (
	"log"
	"time"
)

func LoopSay(name string, interval int) {
	for {
		log.Println("Hello " + name)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
