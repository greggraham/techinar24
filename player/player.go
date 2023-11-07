package player

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
