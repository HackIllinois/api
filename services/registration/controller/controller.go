package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HackIllinois/api/common/datastore"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"github.com/HackIllinois/api/services/registration/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/", GetAllCurrentRegistrations, "GET", router)

	metrics.RegisterHandler("/attendee/", GetCurrentUserRegistration, "GET", router)
	metrics.RegisterHandler("/attendee/", CreateCurrentUserRegistration, "POST", router)
	metrics.RegisterHandler("/attendee/", UpdateCurrentUserRegistration, "PUT", router)

	metrics.RegisterHandler("/attendee/list/", GetFilteredUserRegistrations, "GET", router)

	metrics.RegisterHandler("/mentor/", GetFilteredUserRegistrations, "GET", router)
	metrics.RegisterHandler("/mentor/", CreateCurrentMentorRegistration, "POST", router)
	metrics.RegisterHandler("/mentor/", UpdateCurrentMentorRegistration, "PUT", router)

	metrics.RegisterHandler("/mentor/list/", GetFilteredMentorRegistrations, "GET", router)

	metrics.RegisterHandler("/{id}/", GetAllRegistrations, "GET", router)
	metrics.RegisterHandler("/attendee/{id}/", GetUserRegistration, "GET", router)
	metrics.RegisterHandler("/mentor/{id}/", GetMentorRegistration, "GET", router)

	metrics.RegisterHandler("/internal/stats/", GetStats, "GET", router)
}

/*
	Endpoint to get all registrations (attendee, mentor) for the current user.
	If registrations could not be found for either attendee or mentor, that field is set to nil/null.
*/
func GetAllCurrentRegistrations(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	user_registration, _ := service.GetUserRegistration(id)

	mentor_registration, _ := service.GetMentorRegistration(id)

	var all_registration = models.AllRegistration{
		Attendee: user_registration,
		Mentor:   mentor_registration,
	}

	json.NewEncoder(w).Encode(&all_registration)
}

/*
	Endpoint to get all registrations (attendee, mentor) for the specified user.
	If registrations could not be found for either attendee or mentor, that field is set to nil.
*/
func GetAllRegistrations(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_registration, _ := service.GetUserRegistration(id)

	mentor_registration, _ := service.GetMentorRegistration(id)

	var all_registration = models.AllRegistration{
		Attendee: user_registration,
		Mentor:   mentor_registration,
	}

	json.NewEncoder(w).Encode(&all_registration)
}

/*
	Endpoint to get the registration for the current user
*/
func GetCurrentUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	user_registration, err := service.GetUserRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's registration."))
		return
	}

	json.NewEncoder(w).Encode(user_registration)
}

/*
	Endpoint to create the registration for the current user.
	On successful creation, adds user to a "registered" mailing list, and sends the user a confirmation mail.
*/
func CreateCurrentUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	user_registration := datastore.NewDataStore(config.REGISTRATION_DEFINITION)
	err := json.NewDecoder(r.Body).Decode(&user_registration)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not decode user registration information. Possible failure in JSON validation, or invalid registration format."))
		return
	}

	user_registration.Data["id"] = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not get user info."))
		return
	}

	user_registration.Data["github"] = user_info.Username

	user_registration.Data["createdAt"] = time.Now().Unix()
	user_registration.Data["updatedAt"] = time.Now().Unix()

	err = service.CreateUserRegistration(id, user_registration)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not create user registration."))
		return
	}

	err = service.AddApplicantRole(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not add applicant role."))
		return
	}

	err = service.AddInitialDecision(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not add initial decision."))
		return
	}

	updated_registration, err := service.GetUserRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user registration."))
		return
	}

	// Add user to mailing list
	mail_list := "registered_users"
	err = service.AddUserToMailList(id, mail_list)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not add user to registered users mailing list."))
		return
	}

	// Send confirmation mail
	mail_template := "registration_confirmation"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not send registration confirmation email."))
		return
	}

	// Update user's name (if needed) using registration info
	is_user_nameless := user_info.FirstName == "" || user_info.LastName == ""

	if is_user_nameless {
		first_name, first_ok := user_registration.Data["firstName"].(string)
		last_name, last_ok := user_registration.Data["lastName"].(string)

		if first_ok && last_ok {
			user_info.FirstName = first_name
			user_info.LastName = last_name
			err = service.SetUserInfo(user_info)

			if err != nil {
				errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not set user's name."))
				return
			}
		}
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to update the registration for the current user.
	On successful update, sends the user a confirmation mail.
*/
func UpdateCurrentUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	user_registration := datastore.NewDataStore(config.REGISTRATION_DEFINITION)
	err := json.NewDecoder(r.Body).Decode(&user_registration)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not decode user registration information. Possible failure in JSON validation, or invalid registration format."))
		return
	}

	user_registration.Data["id"] = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not get user info."))
		return
	}

	original_registration, err := service.GetUserRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user's original registration."))
		return
	}

	user_registration.Data["github"] = user_info.Username

	user_registration.Data["createdAt"] = original_registration.Data["createdAt"]
	user_registration.Data["updatedAt"] = time.Now().Unix()

	err = service.UpdateUserRegistration(id, user_registration)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not update user's registration."))
		return
	}

	updated_registration, err := service.GetUserRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch user's updated registration."))
		return
	}

	mail_template := "registration_update"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not send registration update email."))
		return
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to get user registrations based on filters
*/
func GetFilteredUserRegistrations(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	user_registrations, err := service.GetFilteredUserRegistrations(parameters)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get filtered user registrations."))
		return
	}

	json.NewEncoder(w).Encode(user_registrations)
}

/*
	Endpoint to get the registration for the current mentor
*/
func GetCurrentMentorRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	mentor_registration, err := service.GetMentorRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get mentor registration."))
		return
	}

	json.NewEncoder(w).Encode(mentor_registration)
}

/*
	Endpoint to create the registration for the current mentor
*/
func CreateCurrentMentorRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	mentor_registration := datastore.NewDataStore(config.MENTOR_REGISTRATION_DEFINITION)
	err := json.NewDecoder(r.Body).Decode(&mentor_registration)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not decode mentor registration information. Possible failure in JSON validation, or invalid registration format."))
		return
	}

	mentor_registration.Data["id"] = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get mentor's user info."))
		return
	}

	mentor_registration.Data["github"] = user_info.Username

	mentor_registration.Data["createdAt"] = time.Now().Unix()
	mentor_registration.Data["updatedAt"] = time.Now().Unix()

	err = service.CreateMentorRegistration(id, mentor_registration)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not create mentor registration."))
		return
	}

	err = service.AddMentorRole(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not add mentor role."))
		return
	}

	updated_registration, err := service.GetMentorRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated mentor registration."))
		return
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to update the registration for the current mentor
*/
func UpdateCurrentMentorRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	mentor_registration := datastore.NewDataStore(config.MENTOR_REGISTRATION_DEFINITION)
	err := json.NewDecoder(r.Body).Decode(&mentor_registration)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not decode mentor registration information. Possible failure in JSON validation, or invalid registration format."))
		return
	}

	mentor_registration.Data["id"] = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get mentor's user info."))
		return
	}

	original_registration, err := service.GetMentorRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get mentor registration."))
		return
	}

	mentor_registration.Data["github"] = user_info.Username

	mentor_registration.Data["createdAt"] = original_registration.Data["createdAt"]
	mentor_registration.Data["updatedAt"] = time.Now().Unix()

	err = service.UpdateMentorRegistration(id, mentor_registration)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not update mentor registration."))
		return
	}

	updated_registration, err := service.GetMentorRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated mentor registration."))
		return
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to get mentor registrations based on filters
*/
func GetFilteredMentorRegistrations(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	mentor_registrations, err := service.GetFilteredMentorRegistrations(parameters)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get filtered mentor registrations."))
		return
	}

	json.NewEncoder(w).Encode(mentor_registrations)
}

/*
	Endpoint to get the registration for a specified user
*/
func GetUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_registration, err := service.GetUserRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user registration."))
		return
	}

	json.NewEncoder(w).Encode(user_registration)
}

/*
	Endpoint to get the registration for a specified mentor
*/
func GetMentorRegistration(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	mentor_registration, err := service.GetMentorRegistration(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get mentor registration."))
		return
	}

	json.NewEncoder(w).Encode(mentor_registration)
}

/*
	Endpoint to get registration stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not fetch registration service statistics."))
		return
	}

	json.NewEncoder(w).Encode(stats)
}
