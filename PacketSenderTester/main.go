package packetsendertester

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1")
	go StartSending()
	fmt.Fprintf(conn, "Test")
}

func StartSending() {

}
