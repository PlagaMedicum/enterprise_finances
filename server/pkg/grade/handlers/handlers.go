package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AddInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("POST /grade/add requested")
	w.WriteHeader(http.StatusOK)
}

func EditInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("POST /grade/{id} requested")
	w.WriteHeader(http.StatusOK)
}

func DeleteInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("DELETE /grade/{id} requested")
	w.WriteHeader(http.StatusOK)
}

func GetGradeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Info("GET /grade requested")
	w.WriteHeader(http.StatusOK)
}
