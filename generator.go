package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/Duke9289/go-dnd-dice/diceroller"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	class := os.Args[1]
	level, _ := strconv.Atoi(os.Args[2])

	database, err := sql.Open("sqlite3", "./stats.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	row, err := database.Query("SELECT hitdie FROM class WHERE name = '" + class + "' COLLATE NOCASE")
	if err != nil {
		panic(err)
	}
	var hitdie string
	for row.Next() {
		row.Scan(&hitdie)
	}

	hitPoints := diceroller.MaxRoll(fmt.Sprintf("%d%s", 1, hitdie))
	if level > 1 {
		levelPoints, _ := diceroller.ParseInputString(fmt.Sprintf("%d%s", level-1, hitdie))
		hitPoints = hitPoints + levelPoints
	}
	fmt.Println(hitPoints)

}
