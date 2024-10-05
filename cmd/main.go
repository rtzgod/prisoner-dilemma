package main

import (
	"github.com/rtzgod/prisoner-dilemma/internal/behaviors"
	"github.com/rtzgod/prisoner-dilemma/internal/game"
)

func main() {
	g := game.New(10)

	p1 := behaviors.NewCooperator()
	p2 := behaviors.NewCheater()

	g.Play(p1, p2)
}
