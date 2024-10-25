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

// ----------------------------------------------------------------------------
func NewScreenSaver() *ScreenSaver {
	return &ScreenSaver{
		xPos:        -5,
		yPos:        -5,
		xDelta:      4,
		yDelta:      4,
		size:        12,
		colour:      CreateNewRandomColourStruct(),
		ebitenImage: ebiten.NewImage(SCREEN_WIDTH, SCREEN_HEIGHT),
	}
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) performMove() bool {
	status := true

	g.xPos += g.xDelta
	g.yPos += g.yDelta

	if g.xPos <= 0 {
		g.xPos = 1
		g.xDelta *= -1
		status = false
	}

	if g.yPos <= 0 {
		g.yPos = 1
		g.yDelta *= -1
		status = false
	}

	return status
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) updatePosition() {
	if g.performMove() {
		if g.xPos <= 0 || g.xPos >= float32(SCREEN_WIDTH)-g.size {
			g.xDelta *= -1
		}

		if g.yPos <= 0 || g.yPos >= float32(SCREEN_HEIGHT)-g.size {
			g.yDelta *= -1
		}
	}
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Update() error {

	g.updatePosition()
	g.colour.updateColour()

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
