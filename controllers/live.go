package controllers

import (
	"net/http"
	"github.com/gorilla/websocket"

	"projects.iccode.net/stef-k/socketizer-service/models"
	"github.com/gorilla/mux"
	"fmt"
	"projects.iccode.net/stef-k/socketizer-service/site"
)

func Live(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	host := parameters["host"]

	var upgrader = websocket.Upgrader{
		ReadBufferSize:     1024,
		WriteBufferSize:    1024,
		// Do not check for origin we accept them all
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket", 400)
		return
	} else if err != nil {
		panic(err)
	}
	//fmt.Println("got new client from " + host + " with IP: ", r.RemoteAddr)
	// Check if domain is active and has empty slots
	clientDomain := site.FindDomainByName(host)
	fmt.Println(clientDomain)
	if clientDomain.IsActive() {
		// if is in domain pool check current connections
		index, domain := models.FindDomain(host)
		// check if domain's current connections exceeded max concurrent connections
		if domain.ClientCount() < clientDomain.MaxConcurrentConnections {
			fmt.Println("connecting client")
			client := models.NewClient(ws, host)
			if index == -1 {
				domain := models.NewDomain(host)
				domain.AddClient(client)
				models.AddDomain(domain)
			} else {
				domain.AddClient(client)
			}
			//PoolInfo(w, r)
			msg := models.NewMessage(map[string]string{
				"id" : fmt.Sprintf("%p", client.Connection),
				"message": "socketizer connected",
			})

			client.SendMessage(msg)
		} else {
			fmt.Println("max client count reached")
		}
	} else {
		// client not found or is not active
		fmt.Println("client not found or is not active")
	}
}

