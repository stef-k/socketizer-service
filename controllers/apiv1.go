package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"projects.iccode.net/stef-k/socketizer-service/models"
	"fmt"
	"encoding/json"
	"projects.iccode.net/stef-k/socketizer-service/site"
	"github.com/jbrodriguez/mlog"
	"errors"
)

type Request struct {
	Host         string `json:"host"`
	ApiKey       string `json:"apiKey"`
	PostUrl      string `json:"postUrl"`
	PostId       string `json:"postId"`
	PageForPosts string `json:"pageForPosts"`
	What         string `json:"what"`
	CommentUrl   string `json:"commentUrl"`
	CommentId    string `json:"commentId"`
}

type ServiceMessage struct {
	ServiceKey string `json:"serviceKey"`
}

// BroadcastPool broadcasts a message to the DomainPool
func BroadcastPool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	msg := models.NewMessage(map[string]string{
		"message": params["msg"],
	})

	models.PoolBroadcast(msg)
}

// BroadcastDomain broadcasts a message to a specified domain
func BroadcastDomain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	host := params["host"]
	message := params["msg"]

	index, domain := models.FindDomain(host)

	if index != -1 {
		domain.DomainBroadast(models.NewMessage(map[string]string{
			"message": message,
		}))
	}
}

// PoolInfo get information about the Pool
func PoolInfo(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		mlog.Warning("HTTP400, Empty request")
		panic(errors.New("HTTP400, Empty request"))
		return
	}
	var service ServiceMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&service)

	if err != nil {
		mlog.Warning("HTTP400, Bad request")
		panic(errors.New("HTTP400, Bad request, please provide the Service Key"))
	}

	settings, e := site.GetSettings()
	if e != nil {
		mlog.Warning("could not read settings from database, %s", e)
		panic(errors.New("HTTP500, Could not read settings from database"))
	} else {
		if settings.ServiceKey == service.ServiceKey {
			//if true {
			clientSum := 0
			i, d := models.ListDomains()
			for _, domain := range models.DomainPool {
				clientSum += len(domain.ClientPool)
			}
			data := struct {
				DomainCount string
				DomainList  []string
				ClientSub   string
			}{
				fmt.Sprintf("%v", i),
				d,
				fmt.Sprintf("%v", clientSum),
			}
			js, _ := json.Marshal(data)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		} else {
			mlog.Warning("HTTP403, Forbidden ")
			panic(errors.New("HTTP403, Forbidden"))
		}
	}

}

// Refresh all clients for a domain for a specified post
func ClientRefreshPost(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		mlog.Warning("HTTP400, Empty request")
		http.Error(w, "Empty request", 400)
		return
	}
	var request Request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)

	if err != nil {
		mlog.Warning("HTTP400, Bad request")
		http.Error(w, "Bad request", 400)
	}

	// find client in database - check API key, days left or if is free key
	clientDomain, er := site.FindDomainByApiKey(request.ApiKey)
	if er != nil {
		mlog.Warning("could not read settings from database, %s", er)
	}
	settings, e := site.GetSettings()
	if e != nil {
		mlog.Warning("could not read settings from database, %s", e)
	}
	if clientDomain.IsActive() || settings.FreeKeys {
		index, domain := models.FindDomain(request.Host)

		if index != -1 {
			mlog.Info("Client found/is active/has subscription, broadcasting message to sockets")
			domain.DomainBroadast(models.NewMessage(map[string]string{
				"cmd": "refreshPost",
				"postUrl": request.PostUrl,
				"postId": request.PostId,
				"host": request.Host,
				"pageForPosts": request.PageForPosts,
				"what": request.What,
				"commentUrl" : request.CommentUrl,
				"commentId": request.CommentId,
			}))
			site.UpdateTotalMessagesBroadcasted()
		}
	} else {
		mlog.Info("Client not found/not active/not with subscription")
	}

}
