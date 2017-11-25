package nachtsImStall

// Play plays a full game and returns true if game is won.
// func() is the randomizer [0..6)
func Play(rnd func() int) bool {
	// All animals are awake
	awake := 12
	hay := 5
	for {
		d := rnd()
		// log.Printf("awake: %d, hay:%d. Rolled a %d\n", awake, hay, d)
		switch d {

		// Hay: decrease number of hays
		case 0:
			hay--

		// Cock awakes all animals
		case 1:
			awake = 12

		// everything else (moon) makes one animal go asleep
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
