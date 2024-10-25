package main

import (
	"image/color"
	"math/rand/v2"
)

// ----------------------------------------------------------------------------
const (
	MIN_COLOUR = 5
	MAX_COLOUR = 250
	MIN_ALPHA  = 5
	MAX_ALPHA  = 95
)

// ----------------------------------------------------------------------------
var COLOUR_DARK_GRAY = color.RGBA{R: 128, G: 128, B: 128, A: 255}

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
		red:        uint8(rand.IntN(255)),
		green:      uint8(rand.IntN(255)),
		blue:       uint8(rand.IntN(255)),
		alpha:      uint8(rand.IntN(100)),
		redDelta:   randomDelta(),
		greenDelta: randomDelta(),
		blueDelta:  randomDelta(),
		alphaDelta: randomDelta(),
	}
}

// ----------------------------------------------------------------------------
func (colour *Colour) updateColour() {
	updateColourValueWithinLimits(&colour.red, &colour.redDelta, MIN_COLOUR, MAX_COLOUR)
	updateColourValueWithinLimits(&colour.green, &colour.greenDelta, MIN_COLOUR, MAX_COLOUR)
	updateColourValueWithinLimits(&colour.blue, &colour.blueDelta, MIN_COLOUR, MAX_COLOUR)

	// handle alpha a bit differently, to add some randomness
	if colour.alphaDelta > 0 {
		colour.alpha += 1
	} else {
		colour.alpha -= 1
	}

	if colour.alpha <= MIN_ALPHA || colour.alpha >= MAX_ALPHA {
		colour.alphaDelta *= -1
	}
}
