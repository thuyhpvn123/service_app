package router

import (
	// "encoding/json"

	"encoding/json"

	"github.com/jmoiron/sqlx"
)

//Hub struct
type Hub struct {
	connections map[string]*Client
	broadcast   chan *Broadcast
	register    chan *Client
	unregister  chan *Client
	service     *wsService
}

// NewHub return new Hub object.
func NewHub(db *sqlx.DB) *Hub {
	return &Hub{
		connections: make(map[string]*Client),
		broadcast:   make(chan *Broadcast),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		service:     newWSService(db),
	}
}

// Run Hub's main method.
func (h *Hub) Run() {
	for {
		select {
		// register new connection
		case conn := <-h.register:
			h.connections[conn.uid] = conn

			// get remote address
			wsRAddress := conn.ws.RemoteAddr().String()

			// add log
			h.service.addLog(conn.uid, wsRAddress, "Connected")

		// unregister connection
		case conn := <-h.unregister:
			if _, ok := h.connections[conn.uid]; ok {
				close(conn.sendChan)

				delete(h.connections, conn.uid)

				// unsubscribe from all topic
				h.service.unSubscribeAll(conn.uid)

				// get remote address
				wsRAddress := conn.ws.RemoteAddr().String()

				// add log
				h.service.addLog(conn.uid, wsRAddress, "Disconnected")
			}

		// read incoming message
		case b := <-h.broadcast:
			// add log
			h.service.addLog(b.uid, b.address, string(b.message))

			// unmarshal message
			var msg Message1
			err := json.Unmarshal(b.message, &msg)
			// err := b.connections[b.uid].ReadJSON(&msg)

			if err != nil {
				break
			}

			switch msg.Command {
			case "SUBSCRIBE":
				topic:= msg.Data.Data.(string)
				h.service.subscribe(topic, b.uid)
			case "UNSUBSCRIBE":
				topic:= msg.Data.Data.(string)
				h.service.unSubscribe(topic, b.uid)
			case "PUBLISH":
				// get all subscribers by topic name
				topic:= msg.Data.Data.(string)
				subscribers := h.service.getSubscribers(topic)
				// check if subscribers are greater then zero
				if len(subscribers) > 0 {
					// get subscriber from list
					for _, subscriberID := range subscribers {
						// check if subscriber is not me
						if subscriberID != b.uid {
							// get subscriber connection by id
							if conn, ok := h.connections[subscriberID]; ok {
								select {
								// send message
								case conn.sendChan <- msg:
								default:
									close(conn.sendChan)
									delete(h.connections, conn.uid)
								}
							}
						}
					}
				}
			}
		}
	}
}
