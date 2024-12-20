package main

import (
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
var ops = &ebiten.DrawImageOptions{}
var fontFace = text.NewGoXFace(bitmapfont.Face)

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
func (s *ScreenSaver) canPerformMove() bool {
	mustUpdate := true

	s.xPos += s.xDelta
	s.yPos += s.yDelta

	if s.xPos <= 0 || s.xPos >= float32(SCREEN_WIDTH)-s.size {
		s.xPos = clampFloat32(s.xPos, 0, float32(SCREEN_WIDTH)-s.size)
		s.xDelta *= -1
		mustUpdate = false
	}

	if s.yPos <= 0 || s.yPos >= float32(SCREEN_HEIGHT)-s.size {
		s.yPos = clampFloat32(s.yPos, 0, float32(SCREEN_HEIGHT)-s.size)
		s.yDelta *= -1
		mustUpdate = false
	}

	return mustUpdate
}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) updatePosition() {
	if s.canPerformMove() {
		if s.xPos <= 0 || s.xPos >= float32(SCREEN_WIDTH)-s.size {
			s.xDelta *= -1
		}

		if s.yPos <= 0 || s.yPos >= float32(SCREEN_HEIGHT)-s.size {
			s.yDelta *= -1
		}
	}
}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) checkForKeys() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if s.xDelta > 0 {
			s.xDelta *= -1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if s.xDelta < 0 {
			s.xDelta *= -1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if s.yDelta > 0 {
			s.yDelta *= -1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if s.yDelta < 0 {
			s.yDelta *= -1
		}
	}
}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) Update() error {
	s.checkForKeys()
	s.updatePosition()
	s.colour.updateColour()

	return nil
}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(s.ebitenImage, s.xPos, s.yPos, s.size, s.colour.toColour(), true)

	screen.DrawImage(s.ebitenImage, ops)

	if !ebiten.IsFocused() {
		textOp := &text.DrawOptions{}

		str := "Click me to interact"
		tw, th := text.Measure(str, fontFace, textOp.LineSpacing)
		textOp.GeoM.Translate(float64(SCREEN_WIDTH)/2-(tw/2), float64(SCREEN_HEIGHT)/2-(th/2))
		text.Draw(screen, str, fontFace, textOp)
	}

}

// ----------------------------------------------------------------------------
func (s *ScreenSaver) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
