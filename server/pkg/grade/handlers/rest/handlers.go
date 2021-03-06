package rest

import (
	"encoding/json"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type Controller struct {
	handlers.Usecases
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
	w.WriteHeader(http.StatusCreated)

	g := grade.Grade{}
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		handleError(errors.Wrap(err, "Error decoding json"), w, http.StatusInternalServerError)
		return
	}

	err = c.Usecases.AddInfo(r.Context(), g)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
}

// EditInfo ...
func (c Controller) EditInfo(w http.ResponseWriter, r *http.Request) {
	g := grade.Grade{}
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		handleError(errors.Wrap(err,"Error decoding json"), w, http.StatusInternalServerError)
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
	var d time.Time
	err := d.UnmarshalText([]byte(r.URL.Query().Get("date")))
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
	}

	glist, err := c.Usecases.GetGradeList(r.Context(), d)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(glist)
	if err != nil {
		handleError(errors.Wrap(err, "Error encoding json"), w, http.StatusInternalServerError)
		return
	}
}

// GetGrade ...
func (c Controller) GetGrade(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	g, err := c.Usecases.GetGrade(r.Context(), id)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(g)
	if err != nil {
		handleError(errors.Wrap(err, "Error encoding json"), w, http.StatusInternalServerError)
		return
	}
}
