package main

import (
	"fmt"
	"net"
)

const (
	connHost = "*"
	connPort = ":8080"
	connType = "udp"
)

func handlePacket(buf []byte, rlen int, count int) {
	fmt.Println(string(buf[0:rlen]))
	fmt.Println(count)
}

func main() {
	addr, _ := net.ResolveUDPAddr(connType, connPort)
	fmt.Println(addr.Port)
	sock, _ := net.ListenUDP(connType, addr)

	i := 0
	for {
		i++
		buf := make([]byte, 1024)
		rlen, _, err := sock.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}
		go handlePacket(buf, rlen, i)
	}

}
