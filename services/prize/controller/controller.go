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

	router.HandleFunc("/points/award/", AwardShopPoints).Methods("POST")
	router.HandleFunc("/points/redeem/", RedeemShopPoints).Methods("POST")
}

/*
	GetPrize is the endpoint to get a prize from the given prize id.
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

	err := service.UpdatePrize(prize.ID, prize)

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

/*
	AwardPoints gives the specified number of points to the current user.
*/
func AwardShopPoints(w http.ResponseWriter, r *http.Request) {
	var request models.AwardPointsRequest
	json.NewDecoder(r.Body).Decode(&request)

	id := request.ID
	add_points := request.ShopPoints

	user_points, err := service.AwardPoints(add_points, id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update the user's points"))
		return
	}

	json.NewEncoder(w).Encode(user_points)
}

/*
	RedeemPoints attempts to give the specified item to the user for points.
*/
func RedeemShopPoints(w http.ResponseWriter, r *http.Request) {
	var request models.RedeemPointsRequest
	json.NewDecoder(r.Body).Decode(&request)

	id := request.ID
	item_id := request.ShopItemID

	err := service.RedeemPrize(item_id, id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get prize of id \""+item_id+"\" for user \""+id+"\"."))
		return
	}

	// Could have return UserPoints struct instead
	res := models.RedeemPointsResponse{
		Status: "success",
	}

	json.NewEncoder(w).Encode(res)
}
