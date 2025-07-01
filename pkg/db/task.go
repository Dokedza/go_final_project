package db

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		fmt.Println(err)

	}
	defer db.Close()
	// определите запрос
	res, err := db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (date, title, comment, repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func GetTask(id string) (*Task, error) {

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		fmt.Println(err)

	}
	defer db.Close()

	row := db.QueryRow("SELECT date, title, comment, repeat FROM scheduler WHERE id = : id")
	task := &Task{}
	err = row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTask(task *Task) error {

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		fmt.Println(err)

	}
	defer db.Close()

	res, err := db.Exec("UPDATE scheduler SET date = : date AND title = : title AND comment = : comment AND repeat = : repeat WHERE id = : id",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))
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

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		fmt.Println(err)

	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM scheduler WHERE id = : id ", sql.Named("id", id))
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

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		fmt.Println(err)

	}
	defer db.Close()

	rows, err := db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASK limit", sql.Named("limit", limit))
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

	return task, nil
}

func UpdateDate(next string, id string) error {

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		fmt.Println(err)

	}
	defer db.Close()

	res, err := db.Exec("UPDATE scheduler SET date = : date WHERE id = : id", sql.Named("date", next), sql.Named("id", id))
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
