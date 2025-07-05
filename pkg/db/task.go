package db

import (
	"fmt"
)

type Task struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Title   string `json:"title,omitempty"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
}

func AddTask(task *Task) (int64, error) {

	// определите запрос
	res, err := DB.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (?,?,?,?)", task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetTask(id string) (*Task, error) {

	row := DB.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?", id)
	task := &Task{}
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTask(task *Task) error {

	res, err := DB.Exec("UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?", task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("Неизвестная задача")
	}
	return nil
}

func DeleteTask(id string) error {

	res, err := DB.Exec("DELETE FROM scheduler WHERE id = ? ", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("Неизвестный ID")
	}
	return nil
}

func Tasks(limit int) ([]*Task, error) {

	rows, err := DB.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASС limit ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task []*Task
	for rows.Next() {
		taskValue := new(Task)
		err := rows.Scan(&taskValue.ID, &taskValue.Date, &taskValue.Title, &taskValue.Comment, &taskValue.Repeat)
		if err != nil {
			return nil, err
		}
		task = append(task, taskValue)
	}
	err = rows.Err()

	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateDate(next string, id string) error {

	res, err := DB.Exec("UPDATE scheduler SET date = ? WHERE id = ?", next, id)
	if err != nil {
		return err
	}
	val, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if val == 0 {
		return fmt.Errorf("ошибка")
	}
	return nil
}
