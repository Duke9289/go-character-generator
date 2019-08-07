package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetClass(class string) (dbClass string, hitdie string, preferredAttr string) {
	database, err := sql.Open("sqlite3", "db/chargen.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	var row *sql.Row
	if class != "random" {
		row = database.QueryRow("SELECT name, hitdie, preferred_attr FROM class WHERE name = '" + class + "' COLLATE NOCASE")
	} else {
		row = database.QueryRow("SELECT name, hitdie, preferred_attr FROM class ORDER BY RANDOM() LIMIT 1")
	}
	row.Scan(&dbClass, &hitdie, &preferredAttr)
	return
}
