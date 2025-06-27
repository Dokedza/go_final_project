package api

import (
	"net/http"
	//nextdate "go1f/pkg/api"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
}
