package nachtsImStall

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestOneGame(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	won := Play()
	log.Printf("Won: %v\n", won)
}

func TestManyGames(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	won := 0
	for i := 0; i < 10000; i++ {
		if Play() {
			won++
		}
	}
	ratio := won / 100
	log.Printf("Win rate: %d%%\n", ratio)
	if ratio < 27 || ratio > 29 {
		t.Fatalf("Expected win ratio is [27..28]%%, got %d%%\n", ratio)
	}
}

func BenchmarkManyGames(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		Play()
	}
}
