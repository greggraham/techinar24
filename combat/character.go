package combat

import (
	"fmt"
	"math/rand"
)

type Character struct {
	name       string
	health     int
	attackStr  int
	defenseStr int
}

func (p *Character) Name() string {
	return p.name
}

func (p *Character) AttackStr() int {
	return p.attackStr
}

func (p *Character) DefenseStr() int {
	return p.defenseStr
}

func (p *Character) Health() int {
	return p.health
}

func (p *Character) ApplyDamage(amount int) {
	p.health -= amount
}

func (p *Character) Attack(op Combatant) bool {
	toHit := op.DefenseStr() - p.attackStr + 5
	attackRoll := rand.Intn(10) + 1
	if attackRoll >= toHit {
		damage := rand.Intn(4) + 1
		op.ApplyDamage(damage)
		fmt.Printf("%s hits %s and does %d damage. ", p.name, op.Name(), damage)
		if op.Health() > 0 {
			fmt.Printf("Remaining health for %s is %d.\n", op.Name(), op.Health())
			return true
		} else {
			fmt.Printf("%s died.\n\n", op.Name())
			return false
		}
	} else {
		fmt.Printf("%s misses %s.\n", p.name, op.Name())
		return true
	}
}
