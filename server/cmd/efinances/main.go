package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const addr string = ":1540"

func main() {
	log.Info("Programmee started")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("GET / requested")
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	http.Handle("/", r)
	s := http.Server{
		Addr:    addr,
		Handler: r,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Unexpected http server error: " + err.Error())
	}
}
