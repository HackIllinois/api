package common

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		os.Exit(1)
	}

	return strings.Trim(string(out), "\n")
}

// Returns a Sling client configured with the desired role
func GetSlingClient(role string) *sling.Sling {

	token := "FAKE_TOKEN" // Unauthenticated by default

	if role != "Unauthenticated" {
		root_path := getProjectRootPath()

		// accountgen
		accountgen_cmd := formatCommand(fmt.Sprintf("bin/hackillinois-utility-accountgen -role %v", role))
		accountgen_cmd.Dir = root_path
		_, err := accountgen_cmd.Output()
		if err != nil {
			exit_error := &exec.ExitError{}
			if errors.As(err, &exit_error) {
				// Error code 2 is explicitly allowed, as we don't specify all the args,
				// this removes random spew from the output.
				// I am not liable for the inevitable git blame, this is a good enough hack.
				// Blame Jareth.
				if exit_error.ExitCode() != 2 {
					fmt.Printf("ERROR: %v\n", exit_error)
					os.Exit(1)
				}
			} else {
				fmt.Printf("ERROR: %v\n", err)
				os.Exit(1)
			}
		}

		// tokengen
		tokengen_cmd := formatCommand(fmt.Sprintf("bin/hackillinois-utility-tokengen -role %v", role))
		tokengen_cmd.Dir = root_path
		out, err := tokengen_cmd.Output()
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}

		out_lines := string(out)
		token = strings.Trim(strings.Split(out_lines, "Token:")[1], "\n")
	}

	return sling.New().Base("http://localhost:8000").Client(nil).Add("Authorization", token)
}

func GetLocalMongoSession() *mongo.Client {
	client_options := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), client_options)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
	return client
}
