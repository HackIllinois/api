package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const ProfileFormat string = "JSON"

var ProfileRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUserProfile",
		"GET",
		"/profile/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole})).ThenFunc(GetProfile).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentUserProfile",
		"POST",
		"/profile/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole})).ThenFunc(CreateProfile).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentUserProfile",
		"PUT",
		"/profile/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateProfile).ServeHTTP,
	},
	arbor.Route{
		"DeleteCurrentUserProfile",
		"DELETE",
		"/profile/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(DeleteProfile).ServeHTTP,
	},
	arbor.Route{
		"GetAllProfiles",
		"GET",
		"/profile/list/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetFilteredProfiles).ServeHTTP,
	},
	arbor.Route{
		"GetProfileLeaderboard",
		"GET",
		"/profile/leaderboard/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetProfileLeaderboard).ServeHTTP,
	},
	arbor.Route{
		"GetValidFilteredProfiles",
		"GET",
		"/profile/search/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole}), middleware.IdentificationMiddleware).ThenFunc(GetValidFilteredProfiles).ServeHTTP,
	},
	arbor.Route{
		"RedeemEvent",
		"POST",
		"/profile/event/checkin/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(RedeemEvent).ServeHTTP,
	},
	arbor.Route{
		"AwardPoints",
		"POST",
		"/profile/points/award/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(AwardPoints).ServeHTTP,
	},
	arbor.Route{
		"GetProfileFavorites",
		"GET",
		"/profile/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole}), middleware.IdentificationMiddleware).ThenFunc(GetProfileFavorites).ServeHTTP,
	},
	arbor.Route{
		"AddProfileFavorite",
		"POST",
		"/profile/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole}), middleware.IdentificationMiddleware).ThenFunc(AddProfileFavorite).ServeHTTP,
	},
	arbor.Route{
		"RemoveProfileFavorite",
		"DELETE",
		"/profile/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole}), middleware.IdentificationMiddleware).ThenFunc(RemoveProfileFavorite).ServeHTTP,
	},
	arbor.Route{
		"GetTierThresholds",
		"GET",
		"/profile/tier/threshold/",
		http.HandlerFunc(GetTierThresholds).ServeHTTP,
	},
	// This needs to be the last route in order to prevent endpoints like "search", "leaderboard" from accidentally being routed as the {id} variable.
	arbor.Route{
		"GetUserProfileById",
		"GET",
		"/profile/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.AttendeeRole, authtoken.ApplicantRole, authtoken.StaffRole, authtoken.MentorRole}), middleware.IdentificationMiddleware).ThenFunc(GetProfileById).ServeHTTP,
	},
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func GetProfileById(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func GetValidFilteredProfiles(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func GetFilteredProfiles(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func GetProfileLeaderboard(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func RedeemEvent(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func AwardPoints(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func GetProfileFavorites(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func AddProfileFavorite(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func RemoveProfileFavorite(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}

func GetTierThresholds(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROFILE_SERVICE+r.URL.String(), ProfileFormat, "", r)
}
