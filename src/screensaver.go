package main

import (
	"fmt"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// ----------------------------------------------------------------------------
var fontFace = text.NewGoXFace(bitmapfont.Face)

// ----------------------------------------------------------------------------
type ScreenSaver struct {
	xPos, yPos     int
	xDelta, yDelta int
}

// ----------------------------------------------------------------------------
func NewScreenSaver() *ScreenSaver {
	return &ScreenSaver{
		xPos: 23, yPos: 23, xDelta: 1, yDelta: 1,
	}
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
func (g *ScreenSaver) Update() error {
	return nil
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Draw(screen *ebiten.Image) {
	str := fmt.Sprintf("(v%s) %.0f FPS vs %.0f TPS.", APP_VERSION, ebiten.ActualFPS(), ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, str)
}

// ----------------------------------------------------------------------------
func (g *ScreenSaver) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
