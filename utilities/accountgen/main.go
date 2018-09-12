package main

import (
	"flag"
	"fmt"
	"os"
)

type Roles []string

func (roles *Roles) String() string {
	formatted_roles := ""
	for _, role := range *roles {
		formatted_roles = formatted_roles + role
	}
	return formatted_roles
}

func (roles *Roles) Set(role string) error {
	*roles = append(*roles, role)
	return nil
}

func main() {
	var id string
	flag.StringVar(&id, "id", "localadmin", "The user's id")

	var username string
	flag.StringVar(&username, "username", "localadmin", "The user's username")

	var firstName string
	flag.StringVar(&firstName, "firstname", "local", "The user's first name")

	var lastName string
	flag.StringVar(&lastName, "lastname", "admin", "The user's last name")

	var email string
	flag.StringVar(&email, "email", "localadmin@local.local", "The user's email")

	var roles Roles
	flag.Var(&roles, "role", "The user's role, this flag may be specified multiple times")

	flag.Parse()

	if len(roles) == 0 {
		roles = []string{"Admin", "User"}
	}

	err := CreateAccount(id, roles, username, firstName, lastName, email)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create account\nError: %v\n", err.Error())
		os.Exit(1)
	}
}
