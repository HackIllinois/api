package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const MailFormat string = "JSON"

var MailRoutes = arbor.RouteCollection{
	arbor.Route{
		"SendMail",
		"POST",
		"/mail/send/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(SendMail).ServeHTTP,
	},
	arbor.Route{
		"SendMailList",
		"POST",
		"/mail/send/list/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(SendMailList).ServeHTTP,
	},
	arbor.Route{
		"GetAllMailLists",
		"GET",
		"/mail/list/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetAllMailLists).ServeHTTP,
	},
	arbor.Route{
		"CreateMailList",
		"POST",
		"/mail/list/create/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(CreateMailList).ServeHTTP,
	},
	arbor.Route{
		"AddToMailList",
		"POST",
		"/mail/list/add/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(AddToMailList).ServeHTTP,
	},
	arbor.Route{
		"RemoveFromMailList",
		"POST",
		"/mail/list/remove/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(RemoveFromMailList).ServeHTTP,
	},
	arbor.Route{
		"GetMailList",
		"GET",
		"/mail/list/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetMailList).ServeHTTP,
	},
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.MAIL_SERVICE+r.URL.String(), MailFormat, "", r)
}

func SendMailList(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.MAIL_SERVICE+r.URL.String(), MailFormat, "", r)
}

func CreateMailList(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.MAIL_SERVICE+r.URL.String(), MailFormat, "", r)
}

func AddToMailList(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.MAIL_SERVICE+r.URL.String(), MailFormat, "", r)
}

func RemoveFromMailList(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.MAIL_SERVICE+r.URL.String(), MailFormat, "", r)
}

func GetMailList(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.MAIL_SERVICE+r.URL.String(), MailFormat, "", r)
}

func GetAllMailLists(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.MAIL_SERVICE+r.URL.String(), MailFormat, "", r)
}
