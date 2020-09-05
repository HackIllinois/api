package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/recognition/models"
	"github.com/HackIllinois/api/services/recognition/service"
	"github.com/gorilla/mux"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.HandleFunc("/", CreateRecognition).Methods("POST")
	router.HandleFunc("/", GetAllRecognitions).Methods("GET")

	router.HandleFunc("/filter/", GetFilteredRecognitions).Methods("GET")
	router.HandleFunc("/{id}/", DeleteRecognition).Methods("DELETE")
}

/*
	Endpoint to delete an recognition with the specified id.
	It removes the recognition from the recognition trackers, and every user's tracker.
	On successful deletion, it returns the recognition that was deleted.
*/
func DeleteRecognition(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	recognition, err := service.DeleteRecognition(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not delete either the recognition"))
		return
	}

	json.NewEncoder(w).Encode(recognition)
}

/*
	Endpoint to get all recognitions
*/
func GetAllRecognitions(w http.ResponseWriter, r *http.Request) {
	recognition_list, err := service.GetAllRecognitions()

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get all recognitions."))
		return
	}

	json.NewEncoder(w).Encode(recognition_list)
}

/*
	Endpoint to create an recognition
*/
func CreateRecognition(w http.ResponseWriter, r *http.Request) {
	var recognition models.Recognition
	json.NewDecoder(r.Body).Decode(&recognition)

	recognition.ID = utils.GenerateUniqueID()

	err := service.CreateRecognition(recognition.ID, recognition)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not create new recognition."))
		return
	}

	updated_recognition, err := service.GetRecognition(recognition.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated recognition."))
		return
	}

	json.NewEncoder(w).Encode(updated_recognition)
}


/*
	Endpoint to get recognitions based on filters
*/
func GetFilteredRecognitions(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	recognition, err := service.GetFilteredRecognitions(parameters)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch filtered list of recognitions."))
		return
	}

	json.NewEncoder(w).Encode(recognition)
}
