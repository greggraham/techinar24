package main

import (
	"fmt"
	combat1 "techinar/combat"
)

func combat() int {
	p := combat1.NewPlayer("Joe", 10, 10, 8)
	m := combat1.NewMonster("Orc", 10, 8, 10)
	pLives := true
	mLives := true

	for pLives && mLives {
		mLives = p.Attack(m)
		if !mLives {
			break
		}
		pLives = m.Attack(p)
	}
	if pLives {
		return 1
	} else {
		return 2
	}
}

func main() {
	playerCount := 0
	monsterCount := 0
	for i := 0; i < 1000; i++ {
		if combat() == 1 {
			playerCount++
		} else {
			monsterCount++
		}
	}
	fmt.Printf("Player won %d times.\n", playerCount)
	fmt.Printf("Monster won %d times.\n", monsterCount)
}
