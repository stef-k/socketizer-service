package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"projects.iccode.net/stef-k/socketizer-service/models"
	"fmt"
	"encoding/json"
)

type Request struct {
	Host         string `json:"host"`
	SecretKey    string `json:"secretKey"`
	PostUrl      string `json:"postUrl"`
	PostId       string `json:"postId"`
	PageForPosts string `json:"pageForPosts`
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
}

// Refresh all clients for a domain for a specified post
func ClientRefreshPost(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Empty request", 400)
		return
	}
	var request Request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w, "Bad request", 400)
	}

	index, domain := models.FindDomain(request.Host)

	if index != -1 {
		domain.DomainBroadast(models.NewMessage(map[string]string{
			"cmd": "refreshPost",
			"postUrl": request.PostUrl,
			"postId": request.PostId,
			"host": request.Host,
			"pageForPosts": request.PageForPosts,
		}))
	}
}
