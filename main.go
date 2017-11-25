package nachtsImStall

import (
	"math/rand"
)

// Play plays a full game and returns true if game is won.
// You need to rand.Seed() yourself if you want random behaviour
func Play() bool {
	// All animals are awake
	awake := 12
	hay := 5
	for {
		// throw a regular 6-sided dice
		switch rand.Intn(6) {

		// Hay: decrease number of hays
		case 0:
			hay--

		// Cock awakes all animals
		case 1:
			awake = 12

		// Moon makes one animal go asleep
		default:
			awake--

		}

		// A game is won if all animals are asleep
		if awake == 0 {
			return true
		}

		// Game is lost if no more hay
		if hay < 0 {
			return false
		}
	}
}
