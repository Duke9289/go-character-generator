package main

import (
	"database/sql"
	"flag"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/Duke9289/go-dnd-dice/diceroller"
	"github.com/Duke9289/go-dnd-dice/statrolls"
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

	preferredAttrSplit := strings.Split(preferredAttr, ",")

	hitPoints := diceroller.MaxRoll(fmt.Sprintf("%d%s", 1, hitdie))
	if level > 1 {
		levelPoints, _ := diceroller.ParseInputString(fmt.Sprintf("%d%s", level-1, hitdie))
		hitPoints = hitPoints + levelPoints
	}

	stats := statrolls.ThreeDSix()
	sort.Sort(sort.Reverse(sort.IntSlice(stats)))

	generatedCharacter := Character{
		Class:     class,
		Level:     level,
		HitPoints: hitPoints,
	}

	classFields := reflect.ValueOf(&generatedCharacter)

	fields := classFields.Elem()
	fields.FieldByName(string(preferredAttrSplit[0])).SetInt(int64(stats[0]))
	fields.FieldByName(string(preferredAttrSplit[1])).SetInt(int64(stats[1]))
	fields.FieldByName(string(preferredAttrSplit[2])).SetInt(int64(stats[2]))
	fields.FieldByName(string(preferredAttrSplit[3])).SetInt(int64(stats[3]))
	fields.FieldByName(string(preferredAttrSplit[4])).SetInt(int64(stats[4]))
	fields.FieldByName(string(preferredAttrSplit[5])).SetInt(int64(stats[5]))

	generatedCharacter.print()

}
