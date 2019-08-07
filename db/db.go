package db

import (
	"database/sql"
	"errors"

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

func GetRace(race string) (string, error) {

	database, err := sql.Open("sqlite3", "db/chargen.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	var row *sql.Row
	if race != "random" {
		row = database.QueryRow("SELECT name FROM races WHERE name = '" + race + "' COLLATE NOCASE")
	} else {
		row = database.QueryRow("SELECT name FROM races ORDER BY RANDOM() LIMIT 1")
	}
	var name string
	switch err := row.Scan(&name); err {
	case sql.ErrNoRows:
		return name, errors.New("Class entered is not valid")
	case nil:
		return name, nil
	default:
		panic(err)
	}

}
