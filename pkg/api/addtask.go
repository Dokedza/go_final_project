package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go1f/pkg/db"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// обработка других методов будет добавлена на следующих шагах
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
func checkDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(DateFormat)
	}
	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return err
	}
	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if time.Now().After(t) {
		if task.Repeat == "" {
			task.Date = now.Format(DateFormat)
		} else {
			task.Date = next
		}
	}
	return nil
}
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		writeJson(w, map[string]string{"error": "ошибка десериализации JSON"})
		return
	}
	if task.Title == "" {
		writeJson(w, map[string]string{"error": "не указан заголовок задачи"})
		return
	}
	err = checkDate(&task)
	if err != nil {
		writeJson(w, map[string]string{"error": "дата представлена в формате, отличном от 20060102"})
		return
	}
	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, map[string]string{"error": "правило повторения указано в неправильном формате"})
		return
	}
	writeJson(w, map[string]string{"id": strconv.FormatInt(id, 10)})
}
func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}
