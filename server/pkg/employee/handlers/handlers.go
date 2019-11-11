package handlers

import (
	"encoding/json"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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

func parseID(r *http.Request) (big.Int, error) {
	id, ok := big.Int{}.SetString(mux.Vars(r)["id"], 10)
	if !ok {
		err := errors.Errorf("Error parsing id to big.Int.")
		return big.Int{}, err
	}
	return *id, nil
}

func (c Controller) AddEmployee(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

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
	resp := struct {
		ID big.Int `json:"id"`
	}{ID: e.ID}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
}

func (c Controller) EditEmployee(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var e employee.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
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

func (c Controller) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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
}

func (c Controller) GetEmployeeList(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	elist, err := c.Usecases.GetEmployeeList(r.Context())
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(elist)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
}

func (c Controller) GetEmployeePayments(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method + r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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
		handleError(err, w, http.StatusInternalServerError)
		return
	}
}
