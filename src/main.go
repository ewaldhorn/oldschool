package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
const (
	APP_VERSION           = "0.0.8"
	SCREEN_WIDTH     int  = 1024
	SCREEN_HEIGHT    int  = 768
	TICKS_PER_SECOND int  = 50
	IS_DEBUGGING     bool = false
)

// -------------------------------------------------------------------------
func setupWindow() error {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Old School " + APP_VERSION)
	ebiten.SetTPS(TICKS_PER_SECOND)

	return nil
}

// -------------------------------------------------------------------------
func RunOldSchoolScreenSaver() error {
	if err := setupWindow(); err != nil {
		return err
	}

	return ebiten.RunGame(NewScreenSaver())
}

// ----------------------------------------------------------------------------
func main() {
	if err := RunOldSchoolScreenSaver(); err != nil {
		log.Fatalf("Error running screensaver: %v", err)
	}
}
