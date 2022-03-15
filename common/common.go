package common

import (
	"math/rand"
	"time"
)

func GenerateIntVals(valRange int, minVal int) int {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return randomly generated integer
	return newRand.Intn(valRange) + minVal
}

func GenerateFloat64Vals(valRange float64, minVal float64) float64 {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return randomly generated float64
	return (newRand.Float64() * valRange) + minVal
}
