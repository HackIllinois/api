package common

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dghubble/sling"
	"gopkg.in/mgo.v2"
)

func formatCommand(command string) *exec.Cmd {
	split_command := strings.Split(command, " ")

	return exec.Command(split_command[0], split_command[1:]...)
}

// Retrieves the root path of the project
func getProjectRootPath() string {
	cmd := formatCommand("git rev-parse --show-toplevel")

	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return strings.Trim(string(out), "\n")
}

// Returns a Sling client configured with the desired role
func GetSlingClient(role string) *sling.Sling {
	root_path := getProjectRootPath()

	// accountgen
	accountgen_cmd := formatCommand(fmt.Sprintf("bin/hackillinois-utility-accountgen -role %v", role))
	accountgen_cmd.Dir = root_path
	_, err := accountgen_cmd.Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	// tokengen
	tokengen_cmd := formatCommand(fmt.Sprintf("bin/hackillinois-utility-tokengen -role %v", role))
	tokengen_cmd.Dir = root_path
	out, err := tokengen_cmd.Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	out_lines := string(out)
	token := strings.Trim(strings.Split(out_lines, "Token:")[1], "\n")

	return sling.New().Base("http://localhost:8000").Client(nil).Add("Authorization", token)
}

func GetLocalMongoSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
	return session
}
