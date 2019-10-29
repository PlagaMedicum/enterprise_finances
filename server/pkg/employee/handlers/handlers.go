package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	Usecases
}

func (c Controller) AddEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("POST /employee/add requested")
	w.WriteHeader(http.StatusOK)
}

func (c Controller) EditEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("POST /employee/{id} requested")
	w.WriteHeader(http.StatusOK)
}

func (c Controller) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("DELETE /employee/{id} requested")
	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetEmployeeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("GET /employee requested")
	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetEmployeePayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("GET / requested")
	w.WriteHeader(http.StatusOK)
}
