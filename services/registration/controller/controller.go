package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/registration/models"
	"github.com/HackIllinois/api/services/registration/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(GetAllCurrentRegistrations)).Methods("GET")
	
	router.Handle("/attendee/", alice.New().ThenFunc(GetCurrentUserRegistration)).Methods("GET")
	router.Handle("/attendee/", alice.New().ThenFunc(CreateCurrentUserRegistration)).Methods("POST")
	router.Handle("/attendee/", alice.New().ThenFunc(UpdateCurrentUserRegistration)).Methods("PUT")
	router.Handle("/filter/", alice.New().ThenFunc(GetFilteredUserRegistrations)).Methods("GET")
	
	router.Handle("/mentor/", alice.New().ThenFunc(GetCurrentMentorRegistration)).Methods("GET")
	router.Handle("/mentor/", alice.New().ThenFunc(CreateCurrentMentorRegistration)).Methods("POST")
	router.Handle("/mentor/", alice.New().ThenFunc(UpdateCurrentMentorRegistration)).Methods("PUT")
	
	router.Handle("/{id}/", alice.New().ThenFunc(GetAllRegistrations)).Methods("GET")
	router.Handle("/attendee/{id}/", alice.New().ThenFunc(GetUserRegistration)).Methods("GET")
	router.Handle("/mentor/{id}", alice.New().ThenFunc(GetMentorRegistration)).Methods("GET")
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
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(user_registration)
}

/*
	Endpoint to create the registration for the current user.
+	On successful creation, sends the user a confirmation mail.
*/
func CreateCurrentUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.UnprocessableError("Must provide id"))
	}

	var user_registration models.UserRegistration
	json.NewDecoder(r.Body).Decode(&user_registration)

	user_registration.ID = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	user_registration.GitHub = user_info.Username
	user_registration.Email = user_info.Email
	user_registration.FirstName = user_info.FirstName
	user_registration.LastName = user_info.LastName

	err = service.CreateUserRegistration(id, user_registration)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	err = service.AddApplicantRole(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	err = service.AddInitialDecision(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_registration, err := service.GetUserRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	mail_template := "registration_confirmation"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to update the registration for the current user.
+	On successful update, sends the user a confirmation mail.
*/
func UpdateCurrentUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.UnprocessableError("Must provide id"))
	}

	var user_registration models.UserRegistration
	json.NewDecoder(r.Body).Decode(&user_registration)

	user_registration.ID = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	user_registration.GitHub = user_info.Username
	user_registration.Email = user_info.Email
	user_registration.FirstName = user_info.FirstName
	user_registration.LastName = user_info.LastName

	err = service.UpdateUserRegistration(id, user_registration)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_registration, err := service.GetUserRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	mail_template := "registration_update"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to get registrations based on filters
*/
func GetFilteredUserRegistrations(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	user_registrations, err := service.GetFilteredUserRegistrations(parameters)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(mentor_registration)
}

/*
	Endpoint to create the registration for the current mentor
*/
func CreateCurrentMentorRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.UnprocessableError("Must provide id"))
	}

	var mentor_registration models.MentorRegistration
	json.NewDecoder(r.Body).Decode(&mentor_registration)

	mentor_registration.ID = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	mentor_registration.GitHub = user_info.Username
	mentor_registration.Email = user_info.Email
	mentor_registration.FirstName = user_info.FirstName
	mentor_registration.LastName = user_info.LastName

	err = service.CreateMentorRegistration(id, mentor_registration)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	err = service.AddMentorRole(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_registration, err := service.GetMentorRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to update the registration for the current mentor
*/
func UpdateCurrentMentorRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.UnprocessableError("Must provide id"))
	}

	var mentor_registration models.MentorRegistration
	json.NewDecoder(r.Body).Decode(&mentor_registration)

	mentor_registration.ID = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	mentor_registration.GitHub = user_info.Username
	mentor_registration.Email = user_info.Email
	mentor_registration.FirstName = user_info.FirstName
	mentor_registration.LastName = user_info.LastName

	err = service.UpdateMentorRegistration(id, mentor_registration)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_registration, err := service.GetMentorRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to get the registration for a specified user
*/
func GetUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_registration, err := service.GetUserRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(mentor_registration)
}
