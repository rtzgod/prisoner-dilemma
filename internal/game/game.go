package game

import (
	"fmt"
	"github.com/rtzgod/prisoner-dilemma/internal"
	"github.com/rtzgod/prisoner-dilemma/internal/behaviors"
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

func (g *Game) Play(p1, p2 Player) (Player, Player) {
	for round := 1; round <= g.rounds; round++ {
		g.Match(p1, p2)
		g.MatchResult(round, p1, p2)
	}
	fmt.Printf("Player 1 score: %d. Behavior: %s\n", p1.GetScore(), p1.BehaviorName())
	fmt.Printf("Player 2 score: %d. Behavior: %s\n", p2.GetScore(), p2.BehaviorName())
	return p1, p2
}

func (g *Game) Sandbox() {
	matchCounter := 1
	behaviorsScore := map[string]int{
		"Cooperator": 0,
		"Copycat":    0,
		"Detective":  0,
		"Grudger":    0,
		"Cheater":    0,
	}
	allBehaviors := []Player{behaviors.NewCooperator(), behaviors.NewCopycat(), behaviors.NewDetective(), behaviors.NewGrudger(), behaviors.NewCheater(), behaviors.NewRandom()}
	for i := 0; i < len(allBehaviors)-1; i++ {
		for j := i + 1; j < len(allBehaviors); j++ {
			fmt.Println("Match ", matchCounter)

			p1, p2 := g.Play(allBehaviors[i], allBehaviors[j])

			behaviorsScore[p1.BehaviorName()] += p1.GetScore()
			behaviorsScore[p2.BehaviorName()] += p2.GetScore()

			allBehaviors = []Player{behaviors.NewCooperator(), behaviors.NewCopycat(), behaviors.NewDetective(), behaviors.NewGrudger(), behaviors.NewCheater(), behaviors.NewRandom()}

			matchCounter++
		}
	}
	fmt.Println("Final scores:", behaviorsScore)
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
	} else {
		p1.UpdateScore(internal.BOTH_CHEATED)
		p2.UpdateScore(internal.BOTH_CHEATED)
	}
}
func (g *Game) MatchResult(round int, p1, p2 Player) {
	fmt.Printf("Round %d. Player 1 %s: %d | Player 2 %s: %d\n", round, p1.BehaviorName(), p1.GetScore(), p2.BehaviorName(), p2.GetScore())
}
