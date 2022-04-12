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

type Game struct {
	camera       ui.Camera
	world        *ebiten.Image
	screenHeight float64
	screenWidth  float64
	bodies       []*physics.Body
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

func (g *Game) Update() error {
	for _, b1 := range g.bodies {
		for _, b2 := range g.bodies {
			if b1 == b2 {
				continue
			}
			// gravityForce := physics.Gravitation(b1, b2)

		}
	}

	// handle input
	err := handleInput(g)
	return err
}

// notice that we always draw to the world, as that will be projected onto the screen by the camera
func (g *Game) Draw(screen *ebiten.Image) {
	worldX, worldY := g.camera.ScreenToWorld(ebiten.CursorPosition())
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nMove (WASD/Arrows)\nZoom (QE)\nRotate (R)\nReset (Space)\nCamera Speed (PageUp\\Down)\nEscape to quit", ebiten.CurrentTPS()),
	)

	_r, _g, _b, _a := color.White.RGBA()

	worldW, worldH := g.world.Size()
	worldCenterX := float32(worldW) / 2
	worldCenterY := float32(worldH) / 2

	ui.Circle(g.world, worldCenterX, worldCenterY, 50,
		color.RGBA{
			R: uint8(_r),
			G: uint8(_g),
			B: uint8(_b),
			A: uint8(_a),
		})

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%v\nCursor World Pos: %.2f,%.2f",
			&g.camera,
			worldX, worldY),
		0, int(g.screenHeight)-32,
	)

	// project to screen
	g.camera.Render(g.world, screen)
}

func New(screenWidth, screenHeight float64) *Game {
	g := Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		camera:       ui.Camera{ViewPort: f64.Vec2{screenWidth, screenHeight}, Speed: 1},
		world:        ebiten.NewImage(int(screenWidth), int(screenHeight)),
		bodies: []*physics.Body{
			physics.Earth,
			physics.Moon,
		},
	}
	return &g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	scale := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * scale), int(float64(outsideHeight) * scale)
}
