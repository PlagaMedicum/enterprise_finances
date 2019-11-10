package handlers

import (
	"encoding/json"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"math/big"
	"net/http"
)

type Controller struct {
	Usecases
}

func handleError(err error, w http.ResponseWriter, status int) {
	log.Error(err)

	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(err.Error())
	if err != nil {
		log.Error("Error encoding error in response.")
	}
}

func (c Controller) AddEmployee(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var e employee.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}
	e.ID, err = c.Usecases.AddEmployee(r.Context(), e)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	type resp struct {
		ID big.Int `json:"id"`
	}
	err = json.NewEncoder(w).Encode(resp{ID: e.ID})
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c Controller) EditEmployee(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	id, ok := big.Int{}.SetString(mux.Vars(r)["id"], 10)
	if !ok {
		log.Error("Error parsing id to big.Int.")
	}
	err := c.Usecases.EditEmployee(r.Context(), id)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c Controller) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetEmployeeList(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetEmployeePayments(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
