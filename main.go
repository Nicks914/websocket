package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := wsUpgrade.Upgrade(w, r, nil)

	if err != nil {
		log.Println("conn err :", err)
		return
	}

	for {

		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
			log.Println(err)
		}

		msg = []byte("pong")

		conn.WriteMessage(t, msg)

	}

}

func main() {

	r := gin.Default()

	r.LoadHTMLFiles("index.html")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {

		wsHandler(c.Writer, c.Request)

	})
	r.Run(":4000")
}
