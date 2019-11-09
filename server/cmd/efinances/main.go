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

func main() {
	db := psql.DB{
		User:     "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     5432,
		Database: "efinances",
	}
	db.Connect()

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
		Addr:    ":1540",
		Handler: r,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Unexpected http server error: " + err.Error())
	}
}
