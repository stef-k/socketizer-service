package controllers

import (
	"net/http"
	"github.com/gorilla/websocket"

	"projects.iccode.net/stef-k/socketizer-service/models"
	"fmt"
	"github.com/gorilla/mux"
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
		//http.Error(w, "Not a websocket", 400)
		return
	} else if err != nil {
		panic(err)
	}
	fmt.Println("got new client from " + host + " with IP: ", r.RemoteAddr)
	client := models.NewClient(ws, host)

	index, domain := models.FindDomain(host)
	if index == -1 {
		domain := models.NewDomain(host)
		domain.AddClient(client)
		models.AddDomain(domain)
	} else {
		domain.AddClient(client)
	}
	PoolInfo(w, r)
	msg := models.NewMessage(map[string]string{
		"id" : fmt.Sprintf("%p", client.Connection),
		"message": "socketizer connected",
	})

	client.SendMessage(msg)
}

