package firebase

import (
	"math/rand"
	"sync"
	"time"
)

// Modeled after base64 web-safe chars, but ordered by ASCII.
const pushChars = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

var (
	lastPushTimeMs int64
	lastRandChars  [12]int
	mu             sync.Mutex
	rnd            *rand.Rand
)

// KeyGenerate for firebase ref key generate
func KeyGenerate() string {

	var id [8 + 12]byte

	mu.Lock()
	timeMs := time.Now().UTC().UnixNano() / 1e6

	if timeMs == lastPushTimeMs {

		// increment lastRandChars
		for i := 0; i < 12; i++ {

			lastRandChars[i]++
			if lastRandChars[i] < 64 {
				break
			}
			// increment the next byte
			lastRandChars[i] = 0
		}
	} else {
		genRandPart()
	}

	lastPushTimeMs = timeMs
	// put random as the second part
	for i := 0; i < 12; i++ {
		id[19-i] = pushChars[lastRandChars[i]]
	}

	mu.Unlock()

	// put current time at the beginning
	for i := 7; i >= 0; i-- {
		n := int(timeMs % 64)
		id[i] = pushChars[n]
		timeMs = timeMs / 64
	}

	return string(id[:])
}

func init() {
	// seed to get randomness
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func genRandPart() {
	for i := 0; i < len(lastRandChars); i++ {
		lastRandChars[i] = rnd.Intn(64)
	}
}
