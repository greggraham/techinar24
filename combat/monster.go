package combat

import "math/rand"

type Monster struct {
	Character
	gold int
}

func NewMonster(name string, health int, attackStr int, defenseStr int) *Monster {
	var m Monster
	m.name = name
	m.health = health
	m.attackStr = attackStr
	m.defenseStr = defenseStr
	m.gold = int(rand.Int31n(100))
	return &m
}
