package character

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/Duke9289/go-dnd-dice/diceroller"
	"github.com/Duke9289/go-dnd-dice/statrolls"
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

func (c *Character) Print() {
	fmt.Printf("Class: %s\n", c.Class)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("Hit Points: %d\n", c.HitPoints)
	fmt.Printf("Strength: %d, Constitution: %d, Dexterity: %d, Intelligence: %d, Wisdom: %d, Charisma: %d\n",
		c.Str, c.Con, c.Dex, c.Int, c.Wis, c.Cha)
}

func (c *Character) RollStats(preferredAttr string) {

	classFields := reflect.ValueOf(c)

	preferredAttrSplit := strings.Split(preferredAttr, ",")

	stats := statrolls.ThreeDSix()
	sort.Sort(sort.Reverse(sort.IntSlice(stats)))

	fields := classFields.Elem()
	fields.FieldByName(string(preferredAttrSplit[0])).SetInt(int64(stats[0]))
	fields.FieldByName(string(preferredAttrSplit[1])).SetInt(int64(stats[1]))
	fields.FieldByName(string(preferredAttrSplit[2])).SetInt(int64(stats[2]))
	fields.FieldByName(string(preferredAttrSplit[3])).SetInt(int64(stats[3]))
	fields.FieldByName(string(preferredAttrSplit[4])).SetInt(int64(stats[4]))
	fields.FieldByName(string(preferredAttrSplit[5])).SetInt(int64(stats[5]))
}

func (c *Character) RollHitpoints(hitDie string, level int) {
	hitPoints := diceroller.MaxRoll(fmt.Sprintf("%d%s", 1, hitDie))
	if level > 1 {
		levelPoints, _ := diceroller.ParseInputString(fmt.Sprintf("%d%s", level-1, hitDie))
		hitPoints = hitPoints + levelPoints
	}
	c.HitPoints = hitPoints
}
