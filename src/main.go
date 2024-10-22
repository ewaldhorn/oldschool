package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
var TPS int = 50

const (
	APP_VERSION       = "0.0.3"
	IS_DEBUGGING      = false
	SCREEN_WIDTH  int = 1024
	SCREEN_HEIGHT int = 768
)

// ----------------------------------------------------------------------------
func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Old School " + APP_VERSION)
	ebiten.SetTPS(TPS)

	err := ebiten.RunGame(NewScreenSaver())

	if err != nil {
		log.Fatal(err)
	}
}
