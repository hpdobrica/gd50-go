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

func aabbCollides(aX, aY, aWidth, aHeight, bX, bY, bWidth, bHeight float64) bool {

	if (aX < bX+bWidth) && (aX+aWidth > bX) && (aY < bY+bHeight) && (aY+aHeight > bY) {
		return true
	}
	return false

}
