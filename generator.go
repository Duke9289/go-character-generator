package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/Duke9289/go-dnd-dice/diceroller"
	_ "github.com/mattn/go-sqlite3"
)

type Character struct {
	//Race  string
	Class     string
	Level     int
	HitPoints int
	Str       int
	Con       int
	Dex       int
	Int       int
	Wis       int
	Cha       int
}

func (c *Character) print() {
	fmt.Printf("Class: %s\n", c.Class)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("Hit Points: %d\n", c.HitPoints)
	fmt.Printf("Strength: %d, Constitution: %d, Dexterity: %d, Intelligence: %d, Wisdom: %d, Charisma: %d\n",
		c.Str, c.Con, c.Dex, c.Int, c.Wis, c.Cha)
}

func main() {

	const (
		classDefault = "random"
		classUsage   = "The Character's class"
		levelDefault = 1
		levelUsage   = "The character's level"
	)

	var class string
	var level int

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
		row = database.QueryRow("SELECT name, hitdie, preferred_attr FROM class WHERE name = '" + class + "' COLLATE NOCASE")
	} else {
		row = database.QueryRow("SELECT name, hitdie, preferred_attr FROM class ORDER BY RANDOM() LIMIT 1")
	}
	var hitdie string
	var preferredAttr string
	row.Scan(&class, &hitdie, &preferredAttr)

	hitPoints := diceroller.MaxRoll(fmt.Sprintf("%d%s", 1, hitdie))
	if level > 1 {
		levelPoints, _ := diceroller.ParseInputString(fmt.Sprintf("%d%s", level-1, hitdie))
		hitPoints = hitPoints + levelPoints
	}

	generatedCharacter := &Character{
		Class:     class,
		Level:     level,
		HitPoints: hitPoints,
	}

	generatedCharacter.print()

}
