package main

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/Hofled/orbital-simulation/consts"
	"github.com/Hofled/orbital-simulation/physics"
	"github.com/Hofled/orbital-simulation/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/math/f64"
)

const (
	maxTPS int = 60
	// TODO debug value scaled down for slower simulation
	dt float64 = 0.05 / float64(maxTPS)
)

type Game struct {
	camera       ui.Camera
	world        *ebiten.Image
	screenHeight float64
	screenWidth  float64
	planets      []*physics.Planet
	// planet pairs for gravitation comparison
	planetPairs [][]*physics.Planet
}

func calcPlanetPairs(g *Game) {
	for i := 0; i < len(g.planets)-1; i++ {
		planet := g.planets[i]
		for j := i + 1; j < len(g.planets); j++ {
			g.planetPairs = append(g.planetPairs, []*physics.Planet{planet, g.planets[j]})
		}
	}
}

func handleInput(g *Game) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.camera.Translate(ui.MoveDirectionLeft)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.camera.Translate(ui.MoveDirectionRight)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.camera.Translate(ui.MoveDirectionUp)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.camera.Translate(ui.MoveDirectionDown)
	}

	if ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		g.camera.UpdateSpeed(0.1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		g.camera.UpdateSpeed(-0.1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		if g.camera.ZoomFactor > -2400 {
			g.camera.ZoomFactor -= 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		if g.camera.ZoomFactor < 2400 {
			g.camera.ZoomFactor += 1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.camera.Rotation += 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.camera.Reset()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New(consts.ErrRegularTermination)
	}

	return nil
}

func simulatePhysics(g *Game) error {
	// calculate gravity and apply forces
	for _, pair := range g.planetPairs {
		p1 := pair[0]
		p2 := pair[1]
		gravityForce := physics.Gravitation(p1.Body, p2.Body, physics.EarthMoonScaleRatio)
		// apply gravitational force
		physics.ApplyForce(p2.Body, gravityForce, dt)
		// update position
		physics.ApplyMovement(p2.Body, dt)
	}

	return nil
}

func (g *Game) Update() error {
	// handle input
	if err := handleInput(g); err != nil {
		return err
	}

	// simulate physics
	if err := simulatePhysics(g); err != nil {
		return err
	}

	return nil
}

// notice that we always draw to the world, as that will be projected onto the screen by the camera
func (g *Game) Draw(screen *ebiten.Image) {
	// clear world
	g.world.Clear()

	// draw planets
	for _, planet := range g.planets {
		newRad := physics.LogScalePlanetSize(planet.Body.Radius)
		ui.Circle(g.world, float32(planet.Body.Position.AtVec(0)), float32(planet.Body.Position.AtVec(1)), float32(newRad), planet.Color)
		ui.DrawArrowTo(g.world, planet.Body.Position.AtVec(0), planet.Body.Position.AtVec(1), 3, physics.GetScaledVec(planet.Body.Velocity, 15), color.RGBA{0xff, 0, 0, 0xff})
	}

	// project to screen
	g.camera.Render(g.world, screen)

	// debug statistics
	// =====================
	worldX, worldY := g.camera.ScreenToWorld(ebiten.CursorPosition())
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nMove (WASD/Arrows)\nZoom (QE)\nRotate (R)\nReset (Space)\nCamera Speed (PageUp\\Down)\nEscape to quit", ebiten.CurrentTPS()),
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%v\nCursor World Pos: %.2f,%.2f",
			&g.camera,
			worldX, worldY),
		0, int(g.screenHeight)-32,
	)
	// =====================
}

func New(screenWidth, screenHeight float64) *Game {
	g := Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		camera:       ui.Camera{ViewPort: f64.Vec2{screenWidth, screenHeight}, Speed: 1},
		world:        ebiten.NewImage(int(screenWidth), int(screenHeight)),
		planets: []*physics.Planet{
			physics.Earth,
			physics.Moon,
		},
	}
	calcPlanetPairs(&g)
	ebiten.SetMaxTPS(maxTPS)
	return &g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	scale := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * scale), int(float64(outsideHeight) * scale)
}
