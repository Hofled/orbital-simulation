package ui

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type Camera struct {
	ViewPort   f64.Vec2
	Position   f64.Vec2
	ZoomFactor int
	Speed      float64
	Rotation   int
}

const (
	minCameraSpeed float64 = 0.1
	maxCameraSpeed float64 = 20
)

const (
	MoveDirectionUp = iota
	MoveDirectionDown
	MoveDirectionRight
	MoveDirectionLeft
)

func (c *Camera) String() string {
	return fmt.Sprintf(
		"Top Left: %.1f, Rotation: %d, Zoom: %d, Speed: %f",
		c.Position, c.Rotation, c.ZoomFactor, c.Speed,
	)
}

func (c *Camera) viewportCenter() f64.Vec2 {
	return f64.Vec2{
		c.ViewPort[0] * 0.5,
		c.ViewPort[1] * 0.5,
	}
}

func (c *Camera) UpdateSpeed(delta float64) {
	newSpeed := c.Speed + delta
	if newSpeed < minCameraSpeed {
		c.Speed = minCameraSpeed
		return
	}
	if newSpeed > maxCameraSpeed {
		c.Speed = maxCameraSpeed
		return
	}
	c.Speed = newSpeed
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position[0], -c.Position[1])
	// We want to scale and rotate around center of image / screen
	m.Translate(-c.viewportCenter()[0], -c.viewportCenter()[1])
	m.Scale(
		math.Pow(1.01, float64(c.ZoomFactor)),
		math.Pow(1.01, float64(c.ZoomFactor)),
	)
	m.Rotate(float64(c.Rotation) * 2 * math.Pi / 360)
	m.Translate(c.viewportCenter()[0], c.viewportCenter()[1])
	return m
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

func (c *Camera) ScreenToWorld(posX, posY int) (float64, float64) {
	inverseMatrix := c.worldMatrix()
	if inverseMatrix.IsInvertible() {
		inverseMatrix.Invert()
		return inverseMatrix.Apply(float64(posX), float64(posY))
	} else {
		// When scaling it can happend that matrix is not invertable
		return math.NaN(), math.NaN()
	}
}

func (c *Camera) Translate(moveDirection int) {
	switch moveDirection {
	case MoveDirectionUp:
		{
			c.Position[1] -= c.Speed
			break
		}
	case MoveDirectionDown:
		{
			c.Position[1] += c.Speed
			break
		}
	case MoveDirectionRight:
		{
			c.Position[0] += c.Speed
			break
		}
	case MoveDirectionLeft:
		{
			c.Position[0] -= c.Speed
			break
		}
	}
}

func (c *Camera) Reset() {
	c.Position[0] = 0
	c.Position[1] = 0
	c.Rotation = 0
	c.ZoomFactor = 0
	c.Speed = 1
}
