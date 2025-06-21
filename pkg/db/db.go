package db

import (
	"database/sql"
	"os"
)

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(256) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);
CREATE INDEX date_index ON scheduler (date)
`

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	check := os.IsNotExist(err)

	Db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	if check {
		if _, err = Db.Exec(schema); err != nil {
			return err
		}
	}

	defer Db.Close()
	return nil
}
