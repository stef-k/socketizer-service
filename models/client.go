package models

import (
	"github.com/gorilla/websocket"
	"time"
	"fmt"
	"github.com/jbrodriguez/mlog"
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
	// Pong check
	client.Connection.SetPongHandler(func(string) error {
		client.Connection.SetReadDeadline(time.Now().Add(pongWait)); return nil
	})

	go func() {
		// on exit remove Client and close connection
		defer func() {
			mlog.Info(fmt.Sprintf("Removing socket client for domain %v", client.Domain))
			client.Connection.Close()
			RemoveClient(client)
		}()
		for {
			// Connection check
			_, _, err := client.Connection.ReadMessage()
			if err != nil {
				mlog.Info(fmt.Sprintf("Sock error sending close control for domain %v, err msg: %v", client.Domain, err.Error()))
				break
			}
		}
	}()

	mlog.Info(fmt.Sprintf("Spawn new socket client for domain %v", domain))

	return client
}

func (c Client) SendMessage(msg Message) {
	c.Connection.WriteJSON(msg)
}

// Return total Domains and total Clients
func TotalCons() (int, int) {
	clientSum := 0
	d, _ := ListDomains()
	for _, domain := range DomainPool {
		clientSum += len(domain.ClientPool)
	}
	return d, clientSum
}
