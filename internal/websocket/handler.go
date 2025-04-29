package websocket

import (
	"log"

	"github.com/gofiber/fiber/v2"
	ws "github.com/gofiber/websocket/v2"
)

func Register(app *fiber.App) {
	hub := NewHub()
	go hub.Run()

	app.Use("/ws", PublicHandler(hub))
	app.Use("/ws/room/:roomId", RoomPrivate(hub))
	app.Use("/ws/:auth_id", Private(hub))
}

func RoomPrivate(hub *Hub) fiber.Handler {
	return ws.New(func(c *ws.Conn) {
		client := &Client{
			ID:   "1",
			Conn: c,
			Send: make(chan []byte, 256),
			Hub:  hub,
		}
		go client.WritePump()
		client.ReadPump()
	})
}
func PublicHandler(hub *Hub) fiber.Handler {
	return ws.New(func(c *ws.Conn) {
		client := &Client{
			ID:   "1",
			Conn: c,
			Send: make(chan []byte, 256),
			Hub:  hub,
		}
		go client.WritePump()
		client.ReadPump()
	})
}
func Private(hub *Hub) fiber.Handler {
	return ws.New(func(c *ws.Conn) {
		authID := c.Params("auth_id")
		roomID := c.Query("room_id")

		log.Printf("[ROOM]: %s", roomID)
		if authID == "" {
			log.Printf("AuthID [%s] tidak ditemukan. menutup koneksi websocket", authID)
			c.Close()
			return
		}

		client := &Client{
			ID:   authID,
			Conn: c,
			Send: make(chan []byte, 256),
			Hub:  hub,
		}

		client.Hub.Register <- client

		// Verifikasi jika klien terautentikasi
		if !client.IsAuthenticated() {
			log.Println("User tidak terauthentikasi, menutup koneksi")
			c.Close()
			return
		}
		go client.WritePump()
		client.ReadPump()

	})
}
