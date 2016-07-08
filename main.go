package main

import (
	"log"
	"net/http"
	"projects.iccode.net/stef-k/socketizer-service/router"
	"github.com/jbrodriguez/mlog"
)

func main() {
	// enable the logger and close it on exit
	mlog.StartEx(mlog.LevelInfo, "logs/app.log", 5*1024*1024, 30)
	mlog.DefaultFlags = log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile
	defer  mlog.Stop()
	// the router
	router := router.NewRouter()
	// static files for test and own html
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("static/"))))
	// static files for clients each in its directory
	// example URL: example.com/service/static/wordpress/socketizer.js maps to directory
	// project/static/js/service/wordpres/socketizer.js
	router.PathPrefix("/service/static/").Handler(http.StripPrefix("/service/static/",
		http.FileServer(http.Dir("static/js/service/"))))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

