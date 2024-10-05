package game

import (
	"fmt"
	"github.com/rtzgod/prisoner-dilemma/internal"
)

type Player interface {
	Move() int
	UpdateScore(int)
	GetScore() int
	BehaviorName() string
}

type Game struct {
	rounds int
}

func New(rounds int) *Game {
	return &Game{
		rounds: rounds,
	}
}

func (g *Game) Play(p1, p2 Player) {
	for i := 0; i < g.rounds; i++ {
		g.Match(p1, p2)
	}
	fmt.Printf("Player 1 score: %d. Behavior: %s\n", p1.GetScore(), p1.BehaviorName())
	fmt.Printf("Player 2 score: %d. Behavior: %s\n", p2.GetScore(), p2.BehaviorName())
}

func (g *Game) Match(p1, p2 Player) {
	p1Move := p1.Move()
	p2Move := p2.Move()
	if p1Move == internal.COOPERATE && p2Move == internal.COOPERATE {
		p1.UpdateScore(internal.COOPERATED)
		p2.UpdateScore(internal.COOPERATED)
	} else if p1Move == internal.COOPERATE && p2Move == internal.CHEAT {
		p1.UpdateScore(internal.GET_CHEATED)
		p2.UpdateScore(internal.CHEATED)
	} else if p1Move == internal.CHEAT && p2Move == internal.COOPERATE {
		p1.UpdateScore(internal.CHEATED)
		p2.UpdateScore(internal.GET_CHEATED)
	}
}
