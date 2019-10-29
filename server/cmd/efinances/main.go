package main

import (
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const addr string = ":1540"

func main() {
	log.Info("Programmee started")

	r := mux.NewRouter()
	r.HandleFunc("/employee/add", handlers.AddEmployee).Methods(http.MethodPost)

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
