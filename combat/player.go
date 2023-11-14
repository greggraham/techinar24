package combat

type Player struct {
	Character
	inventory []string
}

func NewPlayer(name string, health int, attackStr int, defenseStr int) *Player {
	var p Player
	p.name = name
	p.health = health
	p.attackStr = attackStr
	p.defenseStr = defenseStr
	p.inventory = []string{"lantern", "sword", "rations"}
	return &p
}
