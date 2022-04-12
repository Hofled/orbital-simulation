package main

import (
	"log"

	"github.com/Hofled/orbital-simulation/consts"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	screenWidth, screenHeight := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Orbital Simulation")
	simulation := New(float64(screenWidth), float64(screenHeight))
	if err := ebiten.RunGame(simulation); err != nil && err.Error() != consts.ErrRegularTermination {
		log.Fatal(err)
	}
}
