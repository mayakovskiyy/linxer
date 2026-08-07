package linxerl

import (
	"math/rand/v2"
	"fmt"
)

func LinkGen() string {
	
	var charLib = []string{"A", "a", "B", "b", "C", "c", "D", "d", "E", "e", "F", "f", "G", "g", "H", "h", "I", "i", "J", "j", "K", "k", "L", "l", "M", "m", "N", "n", "O", "o", "P", "p", "Q", "q", "R", "r", "S", "s", "T", "t", "U", "u", "V", "v", "W", "w", "X", "x", "Y", "y", "Z", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	var shortLink string

	for i := 0; i < 5; i++ {
		r := rand.IntN(62)
		shortLink += charLib[r]
		fmt.Println(shortLink)
	}

	return shortLink
}