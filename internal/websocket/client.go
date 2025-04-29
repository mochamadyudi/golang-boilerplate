package websocket

import (
	"log"

	ws "github.com/gofiber/websocket/v2"
)

type Client struct {
	ID   string
	Hub  *Hub
	Conn *ws.Conn
	Send chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		log.Println(c.Hub.Clients)
		if err != nil {
			break
		}
		c.Hub.Broadcast <- msg
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				log.Println("[DISCONNECT] Channel Closed")
				return
			}
			log.Printf("[SENT] %s", string(msg))
			if err := c.Conn.WriteMessage(ws.TextMessage, msg); err != nil {
				log.Printf("[ERROR SEND]: %v", err)
				return
			}
		}
	}
}

func (c *Client) IsAuthenticated() bool {
	return c.ID == "AAA"
}
