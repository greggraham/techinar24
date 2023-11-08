package main

import (
	"fmt"
	"techinar/player"
)

func combat() int {
	p1 := player.New("Joe", 10, 10, 8)
	p2 := player.New("Bob", 10, 8, 10)
	p1Lives := true
	p2Lives := true

	for p1Lives && p2Lives {
		p2Lives = p1.Attack(p2)
		if !p2Lives {
			break
		}
		p1Lives = p2.Attack(p1)
	}
	if p1Lives {
		return 1
	} else {
		return 2
	}
}

func main() {
	p1Count := 0
	p2Count := 0
	for i := 0; i < 1000; i++ {
		if combat() == 1 {
			p1Count++
		} else {
			p2Count++
		}
	}
	fmt.Printf("Player 1 won %d times.\n", p1Count)
	fmt.Printf("Player 2 won %d times.\n", p2Count)
}
