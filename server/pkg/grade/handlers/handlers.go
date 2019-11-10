package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	Usecases
}

func (c Controller) AddInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) EditInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) DeleteInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetGradeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info(r.Method + r.URL.Path)
	w.WriteHeader(http.StatusOK)
}
