package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"go1f/pkg/api"

	"github.com/dokedza/go_final_project/pkg/db"
)

func Run() error {
	err := db.Init("scheduler.db")
	if err != nil {
		return err
	}

	port := 7540

	envPort := os.Getenv("TODO_PORT")
	if envPort != "" {
		p, err := strconv.Atoi(envPort)
		if err == nil {
			port = p
		}
	}

	api.Init()

	http.Handle("/", http.FileServer(http.Dir("web")))
	fmt.Printf("Сервер запущен на порту: http://localhost:7540")
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
