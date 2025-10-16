package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new user",
	Long:  `register a new user`,
	RunE:  Registration,
}

func Registration(cmd *cobra.Command, args []string) error {
	username := args[0]
	password := args[1]
	email := args[2]

	if username == "" || password == "" || email == "" {
		return fmt.Errorf("username or password or email is empty")
	}

	host := "http://localhost:8080/user/register"

	jsn := map[string]string{"username": username, "password": password, "email": email}

	body, _ := json.Marshal(jsn)
	resp, err := http.Post(host, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return fmt.Errorf("register failed: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	return nil

}
