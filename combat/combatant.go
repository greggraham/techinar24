package combat

type Combatant interface {
	Name() string
	AttackStr() int
	DefenseStr() int
	Health() int
	ApplyDamage(amount int)
	Attack(op Combatant) bool
}
