package main

import "math/rand/v2"

// ----------------------------------------------------------------------------
func randomDelta() int8 {
	if rand.IntN(500) > 250 {
		return -1
	} else {
		return 1
	}
}

// ----------------------------------------------------------------------------
func updateColourValueWithinLimits(colour *uint8, change *int8, minColour uint8, maxColour uint8) {
	if *change > 0 {
		*colour += 1
	} else {
		*colour -= 1
	}

	if *colour <= minColour || *colour >= maxColour {
		*change *= -1
	}
}
