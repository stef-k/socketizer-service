package controllers

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/stef-k/socketizer-service/models"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/stef-k/socketizer-service/site"
	"github.com/jbrodriguez/mlog"
)

func Live(w http.ResponseWriter, r *http.Request) {

	settings, e  := site.GetSettings()
	// domain found on database
	domainExists := false
	if e != nil {
		mlog.Info("panicking from live could not read settings ", e)
		panic(e)
	}

	if !settings.ServiceIsActive {
		mlog.Info("Service is inactive, aborting connection")
		return
	}

	parameters := mux.Vars(r)
	host := parameters["host"]
	var upgrader = websocket.Upgrader{
		ReadBufferSize:     1024,
		WriteBufferSize:    1024,
		// Do not check for origin we accept them all
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	if _, ok := err.(websocket.HandshakeError); ok {
		panic(err)
		return
	}
	mlog.Info("got new client from %s with IP: %s", host , r.RemoteAddr)
	// Check if domain is active and has empty slots
	clientDomain, er := site.FindDomainByName(host)
	if er == nil {
		domainExists = true
	}

	// server wide connection limits
	connectionLimit := models.GetAllClients()

	// do not connect the client if we have reached server's limits
	if connectionLimit >= settings.MaxConnection {
		mlog.Info("Server reached its limit, aborting connection")
		return
	}

	if clientDomain.IsActive() || settings.FreeKeys {
		// if is in domain pool check current connections
		index, domain := models.FindDomain(host)
		// max connections how many connections has this domain
		connections := 0
		if settings.FreeKeys {
			connections = settings.MaxConcurrentConnections
		} else {
			connections = clientDomain.MaxConcurrentConnections
		}
		// check if domain's current connections exceeded max concurrent connections
		// and domain exists in database
		if domain.ClientCount() < connections && domainExists {

			client := models.NewClient(ws, host)
			if index == -1 {
				domain := models.NewDomain(host)
				domain.AddClient(client)
				models.AddDomain(domain)
			} else {
				domain.AddClient(client)
			}

			d, c := models.TotalCons()
			mlog.Info(fmt.Sprintf("Domains: %v, Clients: %v", d, c))
			// Update stats
			site.IncreaseTotalClientsBy(1)
			site.UpdateMaxConcurrentClients(c)
			site.UpdateMaxConcurrentDomains(d)
			msg := models.NewMessage(map[string]string{
				"id" : fmt.Sprintf("%p", client.Connection),
				"message": "socketizer connected",
			})

			client.SendMessage(msg)
		} else {
			mlog.Info("max client count reached")
			return
		}
	} else {
		// client not found or is not active
		mlog.Info("client not found or is not active")
	}
}

