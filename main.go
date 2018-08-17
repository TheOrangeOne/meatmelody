package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/namsral/flag"
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "static/index.html")
}

func roomHandler(w http.ResponseWriter, r *http.Request, hotel *Hotel) {
	http.ServeFile(w, r, "static/melody.html")
}

func main() {
	var host, port string
	flag.StringVar(&host, "host", "", "http host")
	flag.StringVar(&port, "port", "8080", "http port")
	flag.Parse()
	r := mux.NewRouter()
	hotel := newHotel()

	r.HandleFunc("/", serveHome).Methods("GET")

	r.HandleFunc("/r/{room}", func(w http.ResponseWriter, r *http.Request) {
		roomHandler(w, r, hotel)
	}).Methods("GET")

	r.HandleFunc("/ws/{room}", hotel.serveHotel)
	http.Handle("/", r)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("listening on %s:%s\n", host, port)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
