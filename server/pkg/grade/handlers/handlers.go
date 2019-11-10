package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	Usecases
}

func (c Controller) AddInfo(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

func (c Controller) EditInfo(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

func (c Controller) DeleteInfo(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetGradeList(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
