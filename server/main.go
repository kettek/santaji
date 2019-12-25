// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var STATIC_DIR = "../static"

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	/*if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}*/
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "../index.html")
	} else {
		http.ServeFile(w, r, "../index.html")
	}
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	router := mux.NewRouter().StrictSlash(true)
	//router.PathPrefix(STATIC_DIR).Handler(http.FileServer(http.Dir("."+STATIC_DIR)))
	//router.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))
	router.HandleFunc("/", serveHome)
	//router.Handle("/static/", http.FileServer(http.Dir(STATIC_DIR)))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
