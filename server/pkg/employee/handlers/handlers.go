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

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) EditEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetEmployeeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetEmployeePayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}
