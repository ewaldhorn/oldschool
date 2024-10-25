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
func (s *ScreenSaver) performMove() bool {
	status := true

	s.xPos += s.xDelta
	s.yPos += s.yDelta

	if s.xPos <= 0 {
		s.xPos = 1
		s.xDelta *= -1
		status = false
	}

	if s.yPos <= 0 {
		s.yPos = 1
		s.yDelta *= -1
		status = false
	}

	return status
}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) updatePosition() {
	if s.performMove() {
		if s.xPos <= 0 || s.xPos >= float32(SCREEN_WIDTH)-s.size {
			s.xDelta *= -1
		}

		if s.yPos <= 0 || s.yPos >= float32(SCREEN_HEIGHT)-s.size {
			s.yDelta *= -1
		}
	}
}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) Update() error {

	s.updatePosition()
	s.colour.updateColour()

	return nil
}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(s.ebitenImage, s.xPos, s.yPos, s.size, s.colour.toColour(), true)

	screen.DrawImage(s.ebitenImage, ops)
}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
