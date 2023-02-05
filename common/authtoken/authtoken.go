package authtoken

import (
	"net/http"

	"github.com/HackIllinois/api/common/config"
	"github.com/HackIllinois/api/common/utils"
)

type Role = string

const (
	AdminRole     = "Admin"
	StaffRole     = "Staff"
	MentorRole    = "Mentor"
	ApplicantRole = "Applicant"
	AttendeeRole  = "Attendee"
	UserRole      = "User"
	BlobstoreRole = "Blobstore"
)

func IsAuthorized(secret string, token string, authorized_roles []Role) (bool, error) {
	for _, role := range authorized_roles {
		is_authorized, err := HasRole(secret, token, role)
		if err != nil {
			return false, err
		}

		if is_authorized {
			return true, nil
		}
	}

	return false, nil
}

func HasRole(secret string, token string, required_role Role) (bool, error) {
	roles, err := utils.ExtractFieldFromJWT(secret, token, "roles")
	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role == required_role {
			return true, nil
		}
	}

	return false, nil
}

func IsRequestFromStaffOrHigher(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	is_at_least_staff, err := IsAuthorized(
		config.TOKEN_SECRET,
		token,
		[]Role{StaffRole, AdminRole},
	)

	return err == nil && is_at_least_staff
}
