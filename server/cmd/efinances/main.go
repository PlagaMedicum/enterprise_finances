package main

import (
    psql "github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
    erest "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/handlers/rest"
    erepo "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/repositories/postgresql"
    eusecases "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/usecases"
    grest "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/handlers/rest"
    grepo "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/repositories/postgresql"
    gusecases "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/usecases"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
)

var httpHeaders = map[string]string {
	"Access-Control-Allow-Origin": "*",
	"Access-Control-Allow-Headers": "access-control-allow-origin, content-type",
	"Content-Type": "application/json",
}

func httpMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r * http.Request) {
		log.Info(r.Method, " ", r.URL.Path)

		for k, v := range httpHeaders {
			w.Header().Set(k, v)
		}

		if r.Method == http.MethodOptions {
        	return
		}

		next(w, r)
	}
}

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

	eh := erest.Controller{
		Usecases: eusecases.Controller{
			Repository: erepo.Controller{
				DB: db,
			},
		},
	}
	gh := grest.Controller{
		Usecases: gusecases.Controller{
			Repository: grepo.Controller{
				DB: db,
			},
		},
	}

	mw := httpMiddleware
	r := mux.NewRouter()
	r.HandleFunc("/employee/add", mw(eh.AddEmployee)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/employee/{id}", mw(eh.EditEmployee)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/employee/{id}/delete", mw(eh.DeleteEmployee)).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/employee", mw(eh.GetEmployeeList)).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", mw(eh.GetEmployee)).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}/payments", mw(eh.GetEmployeePayments)).Methods(http.MethodGet)
	r.HandleFunc("/grade/add", mw(gh.AddInfo)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/grade/{id}", mw(gh.EditInfo)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/grade/{id}/delete", mw(gh.DeleteInfo)).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/grade", mw(gh.GetGradeList)).Methods(http.MethodGet)
	r.HandleFunc("/grade/{id}", mw(gh.GetGrade)).Methods(http.MethodGet)
	http.Handle("/", r)

	s := http.Server{
		Addr:    ":1540",
		Handler: r,
	}

	log.Info("Listening to localhost" + s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Unexpected http server error: " + err.Error())
	}
}
