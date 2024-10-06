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

type Grudger struct {
	*Info
	getCheatedOnce bool
}

func NewGrudger() *Grudger {
	return &Grudger{
		Info:           &Info{"Grudger", 0},
		getCheatedOnce: false,
	}
}
func (g *Grudger) Move() int {
	if g.getCheatedOnce {
		return internal.CHEAT
	}
	return internal.COOPERATE
}
func (g *Grudger) UpdateScore(result int) {
	if result == internal.COOPERATED {
		g.score += 2
	} else if result == internal.CHEATED {
		g.score += 3
	} else if result == internal.GET_CHEATED {
		g.score -= 1
		g.getCheatedOnce = true
	}
}

type Detective struct {
	*Info
	movesPreset      []int
	movesPointer     int
	opponentRevenged bool
	copycatBehavior  *Copycat
	cheaterBehavior  *Cheater
}

func NewDetective() *Detective {
	return &Detective{
		Info:             &Info{"Detective", 0},
		movesPreset:      []int{internal.COOPERATE, internal.CHEAT, internal.COOPERATE, internal.COOPERATE},
		movesPointer:     0,
		opponentRevenged: false,
		copycatBehavior:  NewCopycat(),
		cheaterBehavior:  NewCheater(),
	}
}
func (d *Detective) Move() int {
	if d.opponentRevenged && d.movesPointer >= len(d.movesPreset) {
		return d.copycatBehavior.Move()
	}
	if d.movesPointer < len(d.movesPreset) {
		return d.movesPreset[d.movesPointer]
	}
	return d.cheaterBehavior.Move()
}
func (d *Detective) UpdateScore(result int) {
	if result == internal.COOPERATED {
		d.score += 2
		d.copycatBehavior.opponentMove = internal.COOPERATE
	} else if result == internal.GET_CHEATED {
		d.score -= 1
		d.copycatBehavior.opponentMove = internal.CHEAT
	} else if result == internal.CHEATED {
		d.score += 3
		d.copycatBehavior.opponentMove = internal.COOPERATE
	}
	if result == internal.GET_CHEATED && d.movesPointer == 2 {
		d.opponentRevenged = true
	}
	d.movesPointer++
}
