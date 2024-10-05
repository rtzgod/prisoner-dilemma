package behaviors

import (
	"github.com/rtzgod/prisoner-dilemma/internal"
)

type Info struct {
	name  string
	score int
}

func (i *Info) BehaviorName() string {
	return i.name
}
func (i *Info) GetScore() int {
	return i.score
}

type Cooperator struct {
	*Info
}

func NewCooperator() *Cooperator {
	return &Cooperator{
		Info: &Info{"Cooperator", 0},
	}
}
func (c *Cooperator) Move() int {
	return internal.COOPERATE
}
func (c *Cooperator) UpdateScore(result int) {
	if result == internal.COOPERATED {
		c.score += 2
	} else if result == internal.GET_CHEATED {
		c.score -= 1
	}
}

type Copycat struct {
	*Info
	opponentMove int
}

func NewCopycat() *Copycat {
	return &Copycat{
		Info:         &Info{"Copycat", 0},
		opponentMove: 0,
	}
}
func (cc *Copycat) Move() int {
	return cc.opponentMove
}
func (cc *Copycat) UpdateScore(result int) {
	if result == internal.COOPERATED {
		cc.score += 2
		cc.opponentMove = internal.COOPERATE
	} else if result == internal.GET_CHEATED {
		cc.score -= 1
		cc.opponentMove = internal.CHEAT
	} else if result == internal.CHEATED {
		cc.score += 3
		cc.opponentMove = internal.COOPERATE
	}
}

type Cheater struct {
	*Info
}

func NewCheater() *Cheater {
	return &Cheater{
		Info: &Info{"Cheater", 0},
	}
}
func (c *Cheater) Move() int {
	return internal.CHEAT
}
func (c *Cheater) UpdateScore(result int) {
	if result == internal.CHEATED {
		c.score += 3
	}
}
