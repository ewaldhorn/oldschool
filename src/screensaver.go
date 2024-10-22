package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
type ScreenSaver struct {
	xPos, yPos                                  float32
	boxSize                                     float32
	xDelta, yDelta                              float32
	red, green, blue, alpha                     uint8
	redDelta, greenDelta, blueDelta, aplhaDelta int8
	ebitenImage                                 *ebiten.Image
}

// ----------------------------------------------------------------------------
func NewScreenSaver() *ScreenSaver {
	return &ScreenSaver{
		xPos:        1,
		yPos:        1,
		xDelta:      2,
		yDelta:      2,
		boxSize:     20,
		red:         uint8(rand.IntN(255)),
		green:       uint8(rand.IntN(255)),
		blue:        uint8(rand.IntN(255)),
		alpha:       uint8(rand.IntN(255)),
		redDelta:    randomDelta(),
		greenDelta:  randomDelta(),
		blueDelta:   randomDelta(),
		aplhaDelta:  randomDelta(),
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
	updateAndCheckLimits(&g.red, &g.redDelta)

	if g.red > 128 {
		updateAndCheckLimits(&g.green, &g.greenDelta)
	}

	if g.green > 128 {
		updateAndCheckLimits(&g.blue, &g.blueDelta)
	}

	// handle alpha a bit differently, to add some randomness
	if g.aplhaDelta > 0 {
		g.alpha += 1
	} else {
		g.alpha -= 1
	}

	if g.alpha <= 1 || g.alpha >= 75 {
		g.aplhaDelta *= -1
		g.blueDelta = randomDelta()
		g.redDelta = randomDelta()
		g.greenDelta = randomDelta()
	}
}

// ----------------------------------------------------------------------------
func updateAndCheckLimits(colour *uint8, change *int8) {
	if *change > 0 {
		*colour += 1
	} else {
		*colour -= 1
	}

	if *colour <= 1 || *colour >= 254 {
		*change *= -1
	}
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Update() error {
	g.xPos += g.xDelta
	g.yPos += g.yDelta

	if g.xPos <= 0 || g.xPos >= float32(SCREEN_WIDTH)-g.boxSize {
		g.xDelta *= -1
	}

	if g.yPos <= 0 || g.yPos >= float32(SCREEN_HEIGHT)-g.boxSize {
		g.yDelta *= -1
	}

	g.updateColour()

	return nil
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(g.ebitenImage, g.xPos, g.yPos, g.boxSize, g.boxSize, color.RGBA{R: g.red, G: g.green, B: g.blue, A: g.alpha}, true)

	var ops = &ebiten.DrawImageOptions{}
	screen.DrawImage(g.ebitenImage, ops)
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
