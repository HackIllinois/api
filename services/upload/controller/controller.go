package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/services/upload/models"
	"github.com/HackIllinois/api/services/upload/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/resume/upload/", GetUpdateUserResume, "GET", router)
	metrics.RegisterHandler("/resume/", GetCurrentUserResume, "GET", router)
	metrics.RegisterHandler("/resume/{id}/", GetUserResume, "GET", router)

	metrics.RegisterHandler("/photo/upload/", GetUpdateUserPhoto, "GET", router)
	metrics.RegisterHandler("/photo/", GetCurrentUserPhoto, "GET", router)
	metrics.RegisterHandler("/photo/{id}/", GetUserPhoto, "GET", router)

	metrics.RegisterHandler("/blobstore/", CreateBlob, "POST", router)
	metrics.RegisterHandler("/blobstore/", UpdateBlob, "PUT", router)
	metrics.RegisterHandler("/blobstore/{id}/", GetBlob, "GET", router)
	metrics.RegisterHandler("/blobstore/{id}/", DeleteBlob, "DELETE", router)
}

/*
	Endpoint to get a specified user's resume
*/
func GetUserResume(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	resume, err := service.GetUserResumeLink(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "(S3) Cannot fetch user resume link."))
		return
	}

	json.NewEncoder(w).Encode(resume)
}

/*
	Endpoint to get the current user's resume
*/
func GetCurrentUserResume(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	resume, err := service.GetUserResumeLink(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "(S3) Cannot fetch user resume link."))
		return
	}

	json.NewEncoder(w).Encode(resume)
}

/*
	Endpoint to update the specified user's resume
*/
func GetUpdateUserResume(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	resume, err := service.GetUpdateUserResumeLink(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "(S3) Cannot get/update user's resume."))
		return
	}

	json.NewEncoder(w).Encode(resume)
}

/*
	Endpoint to get a specified user's photo
*/
func GetUserPhoto(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	photo, err := service.GetUserPhotoLink(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "(S3) Cannot fetch user photo link."))
		return
	}

	json.NewEncoder(w).Encode(photo)
}

/*
	Endpoint to get the current user's photo
*/
func GetCurrentUserPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	photo, err := service.GetUserPhotoLink(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "(S3) Cannot fetch user photo link."))
		return
	}

	json.NewEncoder(w).Encode(photo)
}

/*
	Endpoint to update the specified user's photo
*/
func GetUpdateUserPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	photo, err := service.GetUpdateUserPhotoLink(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "(S3) Cannot get/update user's photo."))
		return
	}

	json.NewEncoder(w).Encode(photo)
}

/*
	Endpoint to get a blob
*/
func GetBlob(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	blob, err := service.GetBlob(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Unable to retrieve blob."))
		return
	}

	json.NewEncoder(w).Encode(blob)
}

/*
	Endpoint to create and store a blob
*/
func CreateBlob(w http.ResponseWriter, r *http.Request) {
	var blob models.Blob
	json.NewDecoder(r.Body).Decode(&blob)

	if blob.ID == "" {
		errors.WriteError(w, r, errors.InternalError("Must set an id for the blob.", "Must set an id for the blob."))
		return
	}

	err := service.CreateBlob(blob)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Unable to create blob."))
		return
	}

	stored_blob, err := service.GetBlob(blob.ID)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Unable to retrieve blob."))
		return
	}

	json.NewEncoder(w).Encode(stored_blob)
}

/*
	Endpoint to update a blob
*/
func UpdateBlob(w http.ResponseWriter, r *http.Request) {
	var blob models.Blob
	json.NewDecoder(r.Body).Decode(&blob)

	if blob.ID == "" {
		errors.WriteError(w, r, errors.InternalError("Must set an id for the blob.", "Must set an id for the blob."))
		return
	}

	err := service.UpdateBlob(blob)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Unable to update blob."))
		return
	}

	stored_blob, err := service.GetBlob(blob.ID)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Unable to retrieve blob."))
		return
	}

	json.NewEncoder(w).Encode(stored_blob)
}

/*
	Endpoint to delete a blob
*/
func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	blob, err := service.DeleteBlob(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Unable to delete blob."))
		return
	}

	json.NewEncoder(w).Encode(blob)
}
