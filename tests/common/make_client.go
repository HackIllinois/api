package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dghubble/sling"
)

func GetAdminClient() *sling.Sling {
	// First, get an admin authorization token by running `make setup`.
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	fmt.Println(filepath.Dir(filepath.Dir(path)))

	cmd := exec.Command("make", "setup")
	cmd.Dir = filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(path))))
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	out_lines := strings.Split(string(out[:]), "\n")
	admin_token := out_lines[len(out_lines)-3]

	return sling.New().Base("http://localhost:8000").Client(nil).Add("Authorization", admin_token)
}
