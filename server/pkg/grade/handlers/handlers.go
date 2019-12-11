package handlers

import (
	"encoding/json"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Controller struct {
	Usecases
}

func handleError(err error, w http.ResponseWriter, status int) {
	log.Error(err, " status: ", status)

	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(err.Error())
	if err != nil {
		log.Error("Error encoding error in response.")
	}
}

func parseID(r *http.Request) (uint64, error) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	return id, err
}

// AddInfo ...
func (c Controller) AddInfo(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, " ", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, content-type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodOptions {
		return
	}

	g := grade.Grade{}
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		handleError(errors.Errorf("Error decoding json: %s", err), w, http.StatusInternalServerError)
		return
	}
	err = c.Usecases.AddInfo(r.Context(), g)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// EditInfo ...
func (c Controller) EditInfo(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, " ", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, content-type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodOptions {
		return
	}

	g := grade.Grade{}
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		handleError(errors.Errorf("Error decoding json: %s", err), w, http.StatusInternalServerError)
		return
	}
	g.ID, err = parseID(r)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}
	err = c.Usecases.EditInfo(r.Context(), g)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
}

// DeleteInfo ...
func (c Controller) DeleteInfo(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, " ", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, content-type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodOptions {
		return
	}

	id, err := parseID(r)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}
	err = c.Usecases.DeleteInfo(r.Context(), id)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
}

// GetGradeList ...
func (c Controller) GetGradeList(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, " ", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, content-type")
	w.Header().Set("Content-Type", "application/json")

	glist, err := c.Usecases.GetGradeList(r.Context())
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(glist)
	if err != nil {
		handleError(errors.Errorf("Error encoding json: %s", err), w, http.StatusInternalServerError)
		return
	}
}
