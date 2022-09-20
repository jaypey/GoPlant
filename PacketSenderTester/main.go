package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1")

	if err != nil {
		for {
			fmt.Fprintf(conn, "Test")
			time.Sleep(5 * time.Second)
		}
	}
}
