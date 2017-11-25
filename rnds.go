package nachtsImStall

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// stdlib's randomizer
func stdlibs() int {
	return rand.Intn(6)
}

func devUrandom(n int) func() int {
	buf := readDevUrandom(n)
	idx := 0
	return func() int {
		if idx >= len(buf) {
			log.Printf("Accessing illegal index %d, len %d\n",
				idx, len(buf))
		}
		m := int(buf[idx])
		idx++
		// TODO we introduce a little cheating here, as 256 cannot
		// cleanly be divided by 6
		if m > 251 {
			m = 251
		}
		return m / 42
	}
}

func readDevUrandom(n int) []byte {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, n)
	read, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	if n != read {
		log.Fatalf("Expected to read %d bytes but read %d\n", n, read)
	}
	f.Close()
	return buf
}
