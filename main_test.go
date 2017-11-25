package main

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestOneGame(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	g := NewGame()
	won := g.Play()
	log.Printf("Won: %v\n", won)
}

func TestManyGames(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	won := 0
	for i := 0; i < 1000; i++ {
		g := NewGame()
		if g.Play() {
			won++
		}
	}
	log.Printf("Win rate: %d%%\n", won/10)
}

func BenchmarkManyGames(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		NewGame().Play()
	}
}
