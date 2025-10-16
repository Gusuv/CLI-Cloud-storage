package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a user",
	RunE:  UserLogin,
}

func UserLogin(cmd *cobra.Command, args []string) error {
	var username string = args[0]
	var password string = args[1]

	if username == "" || password == "" {
		return fmt.Errorf("username or password is empty")
	}
	body, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		log.Fatal("Error marshalling body:")
	}

	resp, err := http.Post("http://localhost:8080/user/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("Error while logging in: %v", "Bibi")
	}
	defer resp.Body.Close()

	return nil
}

// Дописать нормальные ответы пользователю
// Проверка что ответил сервер
// Добавить отдельные cobra флаги под username & password
// И в идеале замарочиться с ssl
