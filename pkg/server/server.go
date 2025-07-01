package server

import (
	"fmt"
	"net/http"

	"go1f/pkg/api"

	"github.com/dokedza/go_final_project/pkg/db"
)

func Run() error {
	err := db.Init("sheduler.db")
	if err != nil {
		return err
	}
	api.Init()
	port := 7540
	http.Handle("/", http.FileServer(http.Dir("web")))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
