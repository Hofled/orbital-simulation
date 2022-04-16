package main

import (
	"log"

	"github.com/Hofled/orbital-simulation/consts"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetFullscreen(true)
	screenWidth, screenHeight := ebiten.ScreenSizeInFullscreen()
	ebiten.SetWindowTitle("Orbital Simulation")
	simulation, err := New(float64(screenWidth), float64(screenHeight))
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(simulation); err != nil && err.Error() != consts.ErrRegularTermination {
		log.Fatal(err)
	}
}
