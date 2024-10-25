package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
var ops = &ebiten.DrawImageOptions{}

// ----------------------------------------------------------------------------
type ScreenSaver struct {
	xPos, yPos     float32
	size           float32
	xDelta, yDelta float32
	colour         Colour
	ebitenImage    *ebiten.Image
}

func NewScreenSaver() *ScreenSaver {
	return &ScreenSaver{
		xPos:        1,
		yPos:        1,
		xDelta:      3,
		yDelta:      3,
		size:        12,
		colour:      CreateNewRandomColourStruct(),
		ebitenImage: ebiten.NewImage(SCREEN_WIDTH, SCREEN_HEIGHT),
	}
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) updateColour() {
	updateColourValueWithinLimits(&g.colour.red, &g.colour.redDelta)
	updateColourValueWithinLimits(&g.colour.green, &g.colour.greenDelta)
	updateColourValueWithinLimits(&g.colour.blue, &g.colour.blueDelta)

	// handle alpha a bit differently, to add some randomness
	if g.colour.alphaDelta > 0 {
		g.colour.alpha += 1
	} else {
		g.colour.alpha -= 1
	}

	if g.colour.alpha <= 5 || g.colour.alpha >= 95 {
		g.colour.alphaDelta *= -1
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
	vector.DrawFilledCircle(g.ebitenImage, g.xPos, g.yPos, g.size, g.colour.toColour(), true)

	screen.DrawImage(g.ebitenImage, ops)
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
