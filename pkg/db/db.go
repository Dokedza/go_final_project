package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL,
    comment TEXT,
    repeat VARCHAR(128)
);
CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
`

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if install {
		_, err := db.Exec(schema)
		if err != nil {
			return err
		}
	}
	defer db.Close()
	return nil
}
