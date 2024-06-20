package util

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandAlphabet() string {
	rand.Seed(time.Now().UnixNano())
	ShuffledAlphabet := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x",
		"y", "z"}
	rand.Shuffle(len(ShuffledAlphabet), func(i, j int) {
		ShuffledAlphabet[i], ShuffledAlphabet[j] = ShuffledAlphabet[j], ShuffledAlphabet[i]
	})

	return strings.Join(ShuffledAlphabet, ",")
}
