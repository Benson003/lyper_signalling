package handlers

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/gorilla/websocket"													
)

type Handler struct{}

func (h *Handler) HeartBeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","time":"%s"}`, time.Now().Format(time.RFC3339))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// In production, you’d want to restrict this properly
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) HandleSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to websocket", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		// simple echo server
		fmt.Printf("received: %s\n", msg)
		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("write error:", err)
			break
		}
	}
}