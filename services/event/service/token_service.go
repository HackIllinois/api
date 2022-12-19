package service

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/services/event/config"
)

func IsRequestFromStaffOrHigher(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	is_at_least_staff, err := authtoken.IsAuthorized(config.TOKEN_SECRET, token, []authtoken.Role{authtoken.StaffRole, authtoken.AdminRole})

	return err == nil && is_at_least_staff
}
