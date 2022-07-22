package udpreceiver

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/jaypey/GoPlant/pkg/dal"
	"github.com/jaypey/GoPlant/pkg/models"
	"gorm.io/gorm"
)

const (
	connHost = "*"
	connPort = ":8080"
	connType = "udp"
)

func handlePacket(buf []byte, addr *net.UDPAddr, rlen int, count int) {
	reception := string(buf[0:rlen])
	receptionParts := strings.Split(reception, ";")

	for i := 0; i < len(receptionParts); i++ {
		sensorPart := strings.Split(receptionParts[i], ":")

		//TODO: Add in-memory list of to prevent querying too much
		sensor, err := dal.GetSensorByNameAndIP(sensorPart[0], addr.IP.String())
		if err == gorm.ErrRecordNotFound {
			sensor = models.Sensor{Name: sensorPart[0], IP: addr.IP.String()}
			if _, err := dal.AddSensor(&sensor); err == nil {
				fmt.Println("Error adding sensor")
			}
		}
		if f, err := strconv.ParseFloat(sensorPart[1], 32); err == nil {
			sensorValue := models.SensorValue{Value: f, SensorID: sensor.ID}
			if _, err := dal.AddSensorValue(&sensorValue); err == nil {
				fmt.Println("Error adding sensor value")
			}
		} else {
			fmt.Println("Error parsing value")
		}
	}
	fmt.Println(string(buf[0:rlen]))
	fmt.Println(count)
}

func ListenPacket() {
	fmt.Println("Started listening for udp packets")
	addr, _ := net.ResolveUDPAddr(connType, connPort)
	fmt.Println(addr.Port)
	sock, _ := net.ListenUDP(connType, addr)

	i := 0
	for {
		i++
		buf := make([]byte, 1024)
		rlen, addr, err := sock.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}
		go handlePacket(buf, addr, rlen, i)
	}
}
