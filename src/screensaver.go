package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var COLOUR_DARK_GRAY = color.RGBA{R: 128, G: 128, B: 128, A: 255}

// ----------------------------------------------------------------------------
type ScreenSaver struct {
	xPos, yPos                                  float32
	size                                        float32
	xDelta, yDelta                              float32
	red, green, blue, alpha                     uint8
	redDelta, greenDelta, blueDelta, alphaDelta int8
	ebitenImage                                 *ebiten.Image
}

// ----------------------------------------------------------------------------
func NewScreenSaver() *ScreenSaver {
	return &ScreenSaver{
		xPos:        1,
		yPos:        1,
		xDelta:      3,
		yDelta:      3,
		size:        12,
		red:         uint8(rand.IntN(255)),
		green:       uint8(rand.IntN(255)),
		blue:        uint8(rand.IntN(255)),
		alpha:       uint8(rand.IntN(100)),
		redDelta:    randomDelta(),
		greenDelta:  randomDelta(),
		blueDelta:   randomDelta(),
		alphaDelta:  randomDelta(),
		ebitenImage: ebiten.NewImage(SCREEN_WIDTH, SCREEN_HEIGHT),
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
func (g *ScreenSaver) updateColour() {
	updateColourValueWithinLimits(&g.red, &g.redDelta)
	updateColourValueWithinLimits(&g.green, &g.greenDelta)
	updateColourValueWithinLimits(&g.blue, &g.blueDelta)

	// handle alpha a bit differently, to add some randomness
	if g.alphaDelta > 0 {
		g.alpha += 1
	} else {
		g.alpha -= 1
	}

	if g.alpha <= 5 || g.alpha >= 95 {
		g.alphaDelta *= -1
	}
}

// ----------------------------------------------------------------------------
func updateColourValueWithinLimits(colour *uint8, change *int8) {
	if *change > 0 {
		*colour += 1
	} else {
		*colour -= 1
	}

	if *colour <= 1 || *colour >= 250 {
		*change *= -1
	}
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Update() error {
	g.xPos += g.xDelta
	g.yPos += g.yDelta

	if g.xPos <= 0 || g.xPos >= float32(SCREEN_WIDTH)-g.size {
		g.xDelta *= -1
	}

	if g.yPos <= 0 || g.yPos >= float32(SCREEN_HEIGHT)-g.size {
		g.yDelta *= -1
	}

	g.updateColour()

	return nil
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Draw(screen *ebiten.Image) {

	// vector.DrawFilledRect(g.ebitenImage, g.xPos, g.yPos, g.size, g.size, color.RGBA{R: g.red, G: g.green, B: g.blue, A: g.alpha}, true)
	// vector.StrokeRect(g.ebitenImage, g.xPos, g.yPos, g.size, g.size, 1.0, color.Black, true)

	vector.DrawFilledCircle(g.ebitenImage, g.xPos, g.yPos, g.size, color.RGBA{R: g.red, G: g.green, B: g.blue, A: g.alpha}, true)
	vector.StrokeCircle(g.ebitenImage, g.xPos, g.yPos, g.size, 0.25, COLOUR_DARK_GRAY, true)

	var ops = &ebiten.DrawImageOptions{}
	screen.DrawImage(g.ebitenImage, ops)
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
