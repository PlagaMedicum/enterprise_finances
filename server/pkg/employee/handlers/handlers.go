package handlers

import (
	"encoding/json"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
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

// AddEmployee ...
func (c Controller) AddEmployee(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, " ", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, content-type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodOptions {
		return
	}

	var e employee.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		handleError(errors.Errorf("Error decoding json: %s", err), w, http.StatusBadRequest)
		return
	}
	e.ID, err = c.Usecases.AddEmployee(r.Context(), e)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	resp := struct {
		ID uint64 `json:"id"`
	}{ID: e.ID}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// EditEmployee ...
func (c Controller) EditEmployee(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, " ", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, content-type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodOptions {
		return
	}

	var e employee.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		handleError(errors.Errorf("Error decoding json: %s", err), w, http.StatusBadRequest)
		return
	}
	e.ID, err = parseID(r)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}
	err = c.Usecases.EditEmployee(r.Context(), e)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
}

// DeleteEmployee ...
func (c Controller) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
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
	err = c.Usecases.DeleteEmployee(r.Context(), id)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetEmployeeList ...
func (c Controller) GetEmployeeList(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, " ", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, content-type")
	w.Header().Set("Content-Type", "application/json")

	elist, err := c.Usecases.GetEmployeeList(r.Context())
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(elist)
	if err != nil {
		handleError(errors.Errorf("Error encoding json: %s", err), w, http.StatusInternalServerError)
		return
	}
}

// GetEmployeePayments ...
func (c Controller) GetEmployeePayments(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, " ", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, content-type")
	w.Header().Set("Content-Type", "application/json")

	id, err := parseID(r)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}
	e, err := c.Usecases.GetEmployee(r.Context(), id)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(e)
	if err != nil {
		handleError(errors.Errorf("Error encoding json: %s", err), w, http.StatusInternalServerError)
		return
	}
}
