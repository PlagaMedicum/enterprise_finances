package main

import (
	psql "github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
	ehandlers "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/handlers"
	erepo "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/repositories/postgresql"
	eusecases "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/usecases"
	ghandlers "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/handlers"
	grepo "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/repositories/postgresql"
	gusecases "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/usecases"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	db := psql.DB{
		User:         "postgres",
		Password:     "postgres",
		Host:         "localhost",
		Port:         5432,
		DatabaseName: "efinances",
		SSLMode:      "disable",
	}
	db.Connect()
	db.MigrateDown()
	db.MigrateUp()

	eh := ehandlers.Controller{
		Usecases: eusecases.Controller{
			Repository: erepo.Controller{
				DB: db,
			},
		},
	}
	gh := ghandlers.Controller{
		Usecases: gusecases.Controller{
			Repository: grepo.Controller{
				DB: db,
			},
		},
	}

	r := mux.NewRouter()
	r.HandleFunc("/employee/add", eh.AddEmployee).Methods(http.MethodPost, http.MethodOptions) // TODO: Create another POST request, changing date for id
	r.HandleFunc("/employee/{id}", eh.EditEmployee).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/employee/{id}/delete", eh.DeleteEmployee).Methods(http.MethodDelete)
	r.HandleFunc("/employee", eh.GetEmployeeList).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", eh.GetEmployeePayments).Methods(http.MethodGet)
	r.HandleFunc("/grade/add", gh.AddInfo).Methods(http.MethodPost, http.MethodOptions) // TODO: Create another POST request for adding info, changing date for id
	r.HandleFunc("/grade/{id}", gh.EditInfo).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/grade/{id}/delete", gh.DeleteInfo).Methods(http.MethodDelete)
	r.HandleFunc("/grade", gh.GetGradeList).Methods(http.MethodGet)
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
