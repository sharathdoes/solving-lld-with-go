package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandom16Digit() (*big.Int) {
	min := new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)
	// The range is [min, max) so we want a number up to but not including 10^16.

	// Generate a random number within the specified range
	// max.Sub(max, min) gives the range size, then add min back
	diff := new(big.Int).Sub(max, min)
	randomNum, err := rand.Int(rand.Reader, diff)
	if err != nil {
		panic(err)
	}
	randomNum.Add(randomNum, min)
	return randomNum
}