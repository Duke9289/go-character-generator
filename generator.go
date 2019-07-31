package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/Duke9289/go-dnd-dice/diceroller"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var class string
	var level int
	const (
		classDefault = "random"
		classUsage   = "The Character's class"
		levelDefault = 1
		levelUsage   = "The character's level"
	)

	flag.StringVar(&class, "class", classDefault, classUsage)
	flag.StringVar(&class, "c", classDefault, classUsage)

	flag.IntVar(&level, "level", levelDefault, levelUsage)
	flag.IntVar(&level, "l", levelDefault, levelUsage)

	flag.Parse()

	database, err := sql.Open("sqlite3", "./stats.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	var row *sql.Row
	if class != "random" {
		row = database.QueryRow("SELECT name, hitdie FROM class WHERE name = '" + class + "' COLLATE NOCASE")
	} else {
		row = database.QueryRow("SELECT name, hitdie FROM class ORDER BY RANDOM() LIMIT 1")
	}
	var hitdie string
	row.Scan(&class, &hitdie)

	hitPoints := diceroller.MaxRoll(fmt.Sprintf("%d%s", 1, hitdie))
	if level > 1 {
		levelPoints, _ := diceroller.ParseInputString(fmt.Sprintf("%d%s", level-1, hitdie))
		hitPoints = hitPoints + levelPoints
	}
	fmt.Println(class)
	fmt.Println(level)
	fmt.Println(hitPoints)

}
