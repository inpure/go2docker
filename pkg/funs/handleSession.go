package funs

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

func handleSession(conn net.Conn) {
	log.Println("Session started:")

	reader := bufio.NewReader(conn)

	//Receive data, handle with sayHello
	str, err := reader.ReadString('\n')
	if err == nil {
		str = strings.TrimSpace(str)
		strArray := strings.Fields(str)     // split string with space
		n := strArray[0]                    // name
		i, err := strconv.Atoi(strArray[1]) // interval
		if err == nil {
			LoopSay(n, i) //say hello
		}
	}
}
