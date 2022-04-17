package physics

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"gonum.org/v1/gonum/mat"
)

type Planet struct {
	Body        *Body
	PlanetImage *ebiten.Image
	DrawRadius  float64
	Color       color.RGBA
	IsAttractor bool
}

var (
	Earth = NewPlanet(5.97237e24, 6356.752, 250.0, 250.0, 1.0, 0.0, color.RGBA{0, 0, 0xff, 0xff})
	Moon  = NewPlanet(7.342e22, 1738.1, 289.0, 289.0, 2.0, -2.0, color.RGBA{0x88, 0x88, 0x88, 0xff})
)

// body := NewBody(5.97237e24, 6356.752, mat.NewVecDense(2, []float64{250, 250}), mat.NewVecDense(2, []float64{1, 0}))
func NewPlanet(mass, radius, initialPosX, initialPosY, initialVelocityX, initialVelocityY float64, c color.RGBA) *Planet {
	body := NewBody(mass, radius, mat.NewVecDense(2, []float64{initialPosX, initialPosY}), mat.NewVecDense(2, []float64{initialVelocityX, initialVelocityY}))
	drawRadius := LogScalePlanetSize(body.Radius * 2)
	img := ebiten.NewImage(int(drawRadius), int(drawRadius))
	img.Fill(c)
	return &Planet{
		Body:        body,
		Color:       c,
		IsAttractor: true,
		DrawRadius:  drawRadius,
		PlanetImage: img,
	}
}

// log scales the size of the planet
func LogScalePlanetSize(radius float64) float64 {
	return math.Log(radius) / math.Log(1.2)
}
