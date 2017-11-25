package nachtsImStall

import (
	"log"
	"testing"
)

func TestOneGame(t *testing.T) {
	won := Play(stdlibs)
	log.Printf("Won: %v\n", won)
}

// ration in %
func ratio(won, n int) int {
	return 100 * won / n
}

func TestManyGamesStdlib(t *testing.T) {
	n := 10000
	won := 0
	for i := 0; i < n; i++ {
		if Play(stdlibs) {
			won++
		}
	}
	log.Printf("Win rate(stdlibs): %d%%\n", ratio(won, n))
}

func TestManyGamesDevUrandom(t *testing.T) {
	won := 0
	n := 1000
	for i := 0; i < n; i++ {
		if Play(devUrandom(n)) {
			won++
		}
	}
	log.Printf("Win rate(/dev/urandom): %d%%\n", ratio(won, n))
}

func BenchmarkGamesStdlib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Play(stdlibs)
	}
}

func BenchmarkGamesDevUrandom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f := devUrandom(b.N)
		Play(f)
	}
}

func TestDevUrandom(t *testing.T) {
	readDevUrandom(1)
}

func BenchmarkManySmallReads(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1e6; i++ {
			readDevUrandom(6)
		}
	}
}

func BenchmarkOneLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readDevUrandom(6e6)
	}
}

func TestRndStdlibs(t *testing.T) {
	rndTest(t, 10000, stdlibs)
}

func TestRndDevUrandom(t *testing.T) {
	f := devUrandom(10000)
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
