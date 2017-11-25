package main

import (
	"math/rand"
)

// There are 3 cows, sheeps, cats, and pigs each
const nTotalAnimals = 3 * 4

// 5 portions of hay
const nTotalHay = 5

// State represents a game
type State struct {
	nAwake, nHay int
}

// NewGame starts a fresh game
func NewGame() State {
	return State{nTotalAnimals, nTotalHay}
}

func (a State) isWon() bool {
	return a.nAwake == 0
}

func (a State) isLost() bool {
	return a.nHay < 0
}

func (a *State) cock() {
	a.nAwake = nTotalAnimals
}

func (a *State) hay() {
	a.nHay--
}

func (a *State) moon() {
	a.nAwake--
}

func (a *State) step() {
	var fns = []func(a *State){
		(*State).hay,
		(*State).cock,
		(*State).moon, (*State).moon, (*State).moon, (*State).moon,
	}
	fn := fns[rand.Intn(len(fns))]
	fn(a)
	// log.Printf("State: %+v\n", a)
}

// Play plays a full game and returns true if game is won.
// You need to rand.Seed() yourself if you want random behaviour
func (a State) Play() bool {
	for {
		if a.isWon() {
			return true
		}
		if a.isLost() {
			return false
		}
		a.step()
	}
}
