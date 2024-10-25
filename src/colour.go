package main

import (
	"fmt"
	"image/color"
	"math/rand/v2"
)

// ----------------------------------------------------------------------------
const (
	MIN_COLOUR = 5
	MAX_COLOUR = 99
	MIN_ALPHA  = 5
	MAX_ALPHA  = 95
)

// ----------------------------------------------------------------------------
type Colour struct {
	red, green, blue, alpha                     uint8
	redDelta, greenDelta, blueDelta, alphaDelta int8
}

// ----------------------------------------------------------------------------
func (colour Colour) toColour() color.RGBA {
	return color.RGBA{
		R: colour.red,
		G: colour.green,
		B: colour.blue,
		A: colour.alpha,
	}
}

// ----------------------------------------------------------------------------
func CreateNewRandomColourStruct() Colour {
	return Colour{
		red:        uint8(rand.IntN(99)),
		green:      uint8(rand.IntN(99)),
		blue:       uint8(rand.IntN(99)),
		alpha:      uint8(rand.IntN(50)),
		redDelta:   randomDelta(),
		greenDelta: randomDelta(),
		blueDelta:  randomDelta(),
		alphaDelta: randomDelta(),
	}
}

// ----------------------------------------------------------------------------
func (colour *Colour) updateColour() {

	if randomDelta() == 1 {
		updateColourValueWithinLimits(&colour.red, &colour.redDelta, MIN_COLOUR, MAX_COLOUR)
	}

	if randomDelta() == 1 {
		updateColourValueWithinLimits(&colour.green, &colour.greenDelta, MIN_COLOUR, MAX_COLOUR)
	}

	if randomDelta() == 1 {
		updateColourValueWithinLimits(&colour.blue, &colour.blueDelta, MIN_COLOUR, MAX_COLOUR)
	}

	// handle alpha a bit differently, to add some randomness
	if colour.alphaDelta > 0 {
		colour.alpha += 1
	} else {
		colour.alpha -= 1
	}

	if colour.alpha <= MIN_ALPHA || colour.alpha >= MAX_ALPHA {
		colour.alphaDelta *= -1

		if colour.redDelta%2 == 0 {
			colour.redDelta = randomDelta()
		} else {
			colour.greenDelta = randomDelta()
		}

		if colour.blue%3 == 0 {
			colour.blueDelta = randomDelta()
		} else {
			colour.redDelta = randomDelta()
		}

	}

	if IS_DEBUGGING {
		colour.reportOnColours()
	}
}

// ----------------------------------------------------------------------------
func (colour *Colour) reportOnColours() {
	fmt.Printf("R: %3d (%2d)  G: %3d (%2d)  B: %3d (%2d)  A: %3d (%2d)\n",
		colour.red, colour.redDelta, colour.green, colour.greenDelta,
		colour.blue, colour.blueDelta, colour.alpha, colour.alphaDelta)
}
