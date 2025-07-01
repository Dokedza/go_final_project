package api

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/tasks", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", doneHandler)
	http.HandleFunc("/api/task/delete", deleteTaskHandler)
}
