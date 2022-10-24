package udpreceiver

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/jaypey/GoPlant/pkg/config"
	"github.com/jaypey/GoPlant/pkg/dal"
	"github.com/jaypey/GoPlant/pkg/models"
)

const (
	connHost = "*"
	connPort = ":8080"
	connType = "udp"
)

var SensorCache sync.Map

func handlePacket(buf []byte, addr *net.UDPAddr, rlen int, count int) {
	reception := string(buf[0:rlen])
	receptionParts := strings.Split(reception, ";")

	for i := 0; i < len(receptionParts); i++ {
		sensorPart := strings.Split(receptionParts[i], ":")

		sensorid, found := SensorCache.Load(sensorPart[0])
		if !found {
			sensor := models.Sensor{Name: sensorPart[0], IP: addr.IP.String()}
			if _, err := dal.AddSensor(&sensor); err != nil {
				fmt.Println("Error adding sensor")
			}
			sensorid = sensor.ID
			SensorCache.Store(sensor.Name, sensor.ID)
		}
		if f, err := strconv.ParseFloat(sensorPart[1], 32); err == nil {
			sensorValue := models.SensorValue{Value: f, SensorID: sensorid.(uint)}
			if _, err := dal.AddSensorValue(&sensorValue); err != nil {
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

	//Retrieve existing sensors
	var sensors []models.Sensor
	result := config.GetDB().Select("id", "name").Find(&sensors)
	if result.Error != nil {
		println("[-] Error retrieving sensors")
	} else {
		for _, sensor := range sensors {
			SensorCache.Store(sensor.Name, sensor.ID)
		}
	}

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
