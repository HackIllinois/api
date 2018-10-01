package main

import (
	"github.com/HackIllinois/api/utilities/tokengen/models"
	"flag"
	"fmt"
	"os"
)

type Roles []models.Role

func (roles *Roles) String() string {
	formatted_roles := ""
	for _, role := range *roles {
		formatted_roles = formatted_roles + string(role)
	}
	return formatted_roles
}

func (roles *Roles) Set(role string) error {
	*roles = append(*roles, models.Role(role))
	return nil
}

func main() {
	var id string
	flag.StringVar(&id, "id", "localadmin", "The user's id")

	var exp int64
	flag.Int64Var(&exp, "exp", 2524608000, "The Unix timestamp of expiration")

	var email string
	flag.StringVar(&email, "email", "localadmin@local.local", "The user's email")

	var secret string
	flag.StringVar(&secret, "secret", "secret_string", "The secret to sign the token with")

	var roles Roles
	flag.Var(&roles, "role", "The user's role, this flag may be specified multiple times")

	flag.Parse()

	if len(roles) == 0 {
		roles = []models.Role{models.Admin, models.User}
	}

	token, err := MakeToken(id, exp, email, roles, []byte(secret))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not generate token\nError: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Token:\n%v\n", token)
}
