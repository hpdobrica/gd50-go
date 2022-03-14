package main

import (
	"math"
	"math/rand"
)

func randIntBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

func clamp(v, min, max float64) float64 {
	return math.Min(math.Max(v, min), max)
}
