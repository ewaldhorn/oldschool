package main

import (
	"image/color"
	"math/rand/v2"
)

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
func randomDelta() int8 {
	if rand.IntN(500) > 250 {
		return -1
	} else {
		return 1
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
