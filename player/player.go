package player

import (
	"fmt"
	"math/rand"
)

type Player struct {
	name       string
	health     int
	attackStr  int
	defenseStr int
}

func New(name string, health int, attackStr int, defenseStr int) *Player {
	var p Player
	p.name = name
	p.health = health
	p.attackStr = attackStr
	p.defenseStr = defenseStr
	return &p
}

func (p *Player) Attack(op *Player) bool {
	toHit := op.defenseStr - p.attackStr + 5
	attackRoll := rand.Intn(10) + 1
	if attackRoll >= toHit {
		damage := rand.Intn(4) + 1
		op.health -= damage
		fmt.Printf("%s hits %s and does %d damage. ", p.name, op.name, damage)
		if op.health > 0 {
			fmt.Printf("Remaining health for %s is %d.\n", op.name, op.health)
			return true
		} else {
			fmt.Printf("%s died.\n\n", op.name)
			return false
		}
	} else {
		fmt.Printf("%s misses %s.\n", p.name, op.name)
		return true
	}
}
