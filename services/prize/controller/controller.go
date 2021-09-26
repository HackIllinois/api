package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/prize/models"
	"github.com/HackIllinois/api/services/prize/service"
	"github.com/gorilla/mux"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.HandleFunc("/", GetPrize).Methods("GET")
	router.HandleFunc("/", CreatePrize).Methods("POST")
	router.HandleFunc("/", UpdatePrize).Methods("PUT")
	router.HandleFunc("/", DeletePrize).Methods("DELETE")
}

/*
	GetProfile is the endpoint to get the profile for the current user
*/
func GetPrize(w http.ResponseWriter, r *http.Request) {
	var request models.GetPrizeRequest
	json.NewDecoder(r.Body).Decode(&request)

	prize, err := service.GetPrize(request.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get prize id"))
		return
	}

	json.NewEncoder(w).Encode(prize)
}

func CreatePrize(w http.ResponseWriter, r *http.Request) {
	var prize models.Prize
	json.NewDecoder(r.Body).Decode(&prize)

	err := service.CreatePrize(prize)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not create a new prize"))
		return
	}

	json.NewEncoder(w).Encode(prize)
}

func UpdatePrize(w http.ResponseWriter, r *http.Request) {
	var prize models.Prize
	json.NewDecoder(r.Body).Decode(&prize)

	err := service.UpdatePrize(prize)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update a new prize"))
		return
	}

	json.NewEncoder(w).Encode(prize)
}

func DeletePrize(w http.ResponseWriter, r *http.Request) {
	var request models.DeletePrizeRequest
	json.NewDecoder(r.Body).Decode(&request)

	prize, err := service.GetPrize(request.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Prize does not exist"))
		return
	}

	err = service.DeletePrize(request.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not delete prize"))
		return
	}

	json.NewEncoder(w).Encode(prize)
}
