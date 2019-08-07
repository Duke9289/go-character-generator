package main

import (
	"flag"

	"github.com/Duke9289/go-character-generator/character"
	"github.com/Duke9289/go-character-generator/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	const (
		classDefault = "random"
		classUsage   = "The Character's class"
		levelDefault = 1
		levelUsage   = "The character's level"
		raceDefault  = "random"
		raceUsage    = "The character's race"
	)

	var class string
	var level int
	var race string

	flag.StringVar(&class, "class", classDefault, classUsage)
	flag.StringVar(&class, "c", classDefault, classUsage)

	flag.IntVar(&level, "level", levelDefault, levelUsage)
	flag.IntVar(&level, "l", levelDefault, levelUsage)

	flag.StringVar(&race, "race", raceDefault, raceUsage)
	flag.StringVar(&race, "r", raceDefault, raceUsage)

	flag.Parse()

	class, hitDie, preferredAttr := db.GetClass(class)

	race = db.GetRace(race)

	generatedCharacter := character.Character{
		Class: class,
		Level: level,
		Race:  race,
	}

	generatedCharacter.RollStats(preferredAttr)

	generatedCharacter.RollHitpoints(hitDie, level)

	generatedCharacter.Print()

}
