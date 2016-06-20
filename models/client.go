package models

import (
	"github.com/gorilla/websocket"
	"time"
	"fmt"
)


// Client represents an entity connected using a websocket
// Domain for the client
// Connection the websocket connection
type Client struct {
	Id         string
	Domain     string
	Connection *websocket.Conn
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// NewClient creates a new Client object
func NewClient(ws *websocket.Conn, domain string) *Client {
	client := new(Client)

	client.Id = fmt.Sprintf("%p", ws)
	client.Domain = domain
	client.Connection = ws
	client.Connection.SetReadLimit(maxMessageSize)
	client.Connection.SetReadDeadline(time.Now().Add(pongWait))

	removeCLient := make(chan bool)
	// Ping handler, set periodic pings to check if client is still connected
	ticker := time.NewTicker(pongWait)
	go func() {
		for range ticker.C {
			select {
			case <-removeCLient:
				ticker.Stop()
				return
			default:
				//fmt.Println(time.Now(), " pinging client: ", client.Id)
				err := client.Connection.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait))
				if err != nil {
					//fmt.Println(time.Now(), " client disconect, will remove from domain pool ",  client.Id)
					RemoveClient(client)
					ticker.Stop()
					fmt.Println("Exiting routine")
					removeCLient <- true
					return
				}
			}
		}
	}()
	client.Connection.SetPongHandler(func(string) error {
		client.Connection.SetReadDeadline(time.Now().Add(pongWait)); return nil
	})
	//go func() {
	//	for {
	//		select {
	//		case <-removeCLient:
	//			ticker.Stop()
	//			return
	//		default:
	//			fmt.Println("reading from client ", client.Id)
	//			if _, _, err := client.Connection.NextReader(); err != nil {
	//				if err != nil {
	//					fmt.Println("%v client disconect, setting it inactive ", time.Now(), client.Id)
	//					RemoveClient(client)
	//					ticker.Stop()
	//					fmt.Println("Exiting routine")
	//					removeCLient <- true
	//				}
	//			}
	//		}
	//	}
	//}()
	return client
}

func (c Client)  SendMessage(msg Message) {
	c.Connection.WriteJSON(msg)
}
