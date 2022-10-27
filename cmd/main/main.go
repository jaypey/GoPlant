package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jaypey/GoPlant/pkg/config"
	"github.com/jaypey/GoPlant/pkg/controllers"
	"github.com/jaypey/GoPlant/pkg/udpreceiver"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func main() {
	config.Connect()

	router := gin.Default()
	router.GET("/sensors", controllers.GetSensors)
	router.POST("/sensor", controllers.AddSensor)
	router.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	go udpreceiver.ListenPacket()
	router.Run("localhost:8080")
}
