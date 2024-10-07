package main

import (
	"github.com/rtzgod/prisoner-dilemma/internal/game"
)

func main() {
	g := game.New(10)

	//p1 := behaviors.NewDetective()
	//p2 := behaviors.NewGrudger()
	// g.Play(p1, p2)

	g.Sandbox()
}
