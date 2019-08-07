package main

import (
	"flag"
	"fmt"

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
	)

	var class string
	var level int

	flag.StringVar(&class, "class", classDefault, classUsage)
	flag.StringVar(&class, "c", classDefault, classUsage)

	flag.IntVar(&level, "level", levelDefault, levelUsage)
	flag.IntVar(&level, "l", levelDefault, levelUsage)

	flag.Parse()

	class, hitDie, preferredAttr := db.GetClass(class)

	generatedCharacter := character.Character{
		Class: class,
		Level: level,
	}

	fmt.Println(class)
	fmt.Println(preferredAttr)
	generatedCharacter.RollStats(preferredAttr)

	generatedCharacter.RollHitpoints(hitDie, level)

	generatedCharacter.Print()

}
