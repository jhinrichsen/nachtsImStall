package nachtsImStall

import (
	"log"
	"testing"
)

// Most games should be settled after this number of rolls
const mostRolls = 64

func TestOneGameStdlibs(t *testing.T) {
	won, rolls := Play(stdlibs)
	log.Printf("stdlibs: Won: %v after %d rolls\n", won, rolls)
}

func TestOneGameDevUrandom(t *testing.T) {
	f := devUrandom(mostRolls)
	won, rolls := Play(f)
	log.Printf("/dev/urandom: Won: %v after %d rolls\n", won, rolls)
}

// ration in %
func ratio(won, n int) int {
	return 100 * won / n
}

func TestManyGamesStdlib(t *testing.T) {
	n := 10000
	total := 0
	for i := 0; i < n; i++ {
		won, _ := Play(stdlibs)
		if won {
			total++
		}
	}
	log.Printf("Win rate(stdlibs): %d%%\n", ratio(total, n))
}

func TestManyGamesDevUrandom(t *testing.T) {
	n := 1000
	total := 0
	for i := 0; i < n; i++ {
		won, _ := Play(devUrandom(mostRolls))
		if won {
			total++
		}
	}
	log.Printf("Win rate(/dev/urandom): %d%%\n", ratio(total, n))
}

func BenchmarkGamesStdlib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Play(stdlibs)
	}
}

func BenchmarkGamesDevUrandom(b *testing.B) {
	f := devUrandom(b.N * mostRolls)
	for n := 0; n < b.N; n++ {
		Play(f)
	}
}

func TestDevUrandom(t *testing.T) {
	readDevUrandom(1)
}

func BenchmarkReadManySmallReads(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1e6; i++ {
			readDevUrandom(6)
		}
	}
}

func BenchmarkReadOneLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readDevUrandom(6e6)
	}
}

func TestRndStdlibs(t *testing.T) {
	rndTest(t, 10000, stdlibs)
}

func TestRndDevUrandom(t *testing.T) {
	f := devUrandom(10000 * mostRolls)
	rndTest(t, 10000, f)
}

func rndTest(t *testing.T, n int, f func() int) {
	is := make([]int, 6)
	for i := 0; i < n; i++ {
		r := f()
		if r < 0 || r > 5 {
			log.Fatalf("Expected [0..6) but got %d\n", r)
		}
		is[r]++
	}
	for i, v := range is {
		log.Printf("%d: %d\n", i, v)
	}
}
