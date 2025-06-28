package api

import (
	"encoding/json"
	"net/http"
	"time"

	"go1f/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	tasks, err := db.Tasks(50) // в параметре максимальное количество записей

	if err != nil {
		writeJson(w, map[string]string{"error:": err.Error()})
		return
	}
	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, map[string]string{"error:": "отсутствует id"})
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error:": "отсутствует задача"})
		return
	}
	writeJson(w, task)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJson(w, map[string]string{"error:": err.Error()})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, map[string]string{"error:": err.Error()})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJson(w, map[string]string{"error:": err.Error()})
		return
	}

	if task.Title == "" {
		writeJson(w, map[string]string{"error:": "Отсутствует заголовок задачи"})
		return
	}
	writeJson(w, map[string]string{})
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		writeJson(w, map[string]string{"error:": "Отсутствует id"})
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error:": err.Error()})
		return
	}

	writeJson(w, map[string]string{})
}

func doneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, map[string]string{"error:": "Отсутствует id"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error:": err.Error()})
		return
	}
	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeJson(w, map[string]string{"error:": err.Error()})
			return
		}
		writeJson(w, map[string]string{})
		return
	}
	now := time.Now()
	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		writeJson(w, map[string]string{"error:": err.Error()})
		return
	}
	err = db.UpdateDate(next, id)
	if err != nil {
		writeJson(w, map[string]string{"error:": err.Error()})
		return
	}
	writeJson(w, map[string]string{})
}
