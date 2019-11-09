package main

import (
	psql "github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/handlers"
	repo "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/repositories/postgresql"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/usecases"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const addr string = ":1540"

func main() {
	log.Info("Programmee started")

	db = psql.DB

	h := handlers.Controller{
		usecases.Controller{
			repo.Controller{
				db,
			},
		},
	}

	r := mux.NewRouter()
	r.HandleFunc("/employee/add", h.AddEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/add", h.AddEmployee).Methods(http.MethodPost)

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
