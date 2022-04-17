package physics

import (
	"image/color"
	"math"

	"gonum.org/v1/gonum/mat"
)

type Planet struct {
	Body *Body
	// scaled radius in pixels used for drawing
	DrawRadius  float64
	Color       color.RGBA
	IsAttractor bool
}

var (
	earthBody         = NewBody(5.97237e24, 6356.752, mat.NewVecDense(2, []float64{250, 250}), mat.NewVecDense(2, []float64{1, 0}))
	Earth     *Planet = &Planet{
		Body:        earthBody,
		Color:       color.RGBA{0, 0, 0xff, 0xff}, // blue colored
		IsAttractor: true,
		DrawRadius:  LogScalePlanetSize(earthBody.Radius),
	}
	moonBody         = NewBody(7.342e22, 1738.1, mat.NewVecDense(2, []float64{289, 289}), mat.NewVecDense(2, []float64{2, -2}))
	Moon     *Planet = &Planet{
		Body:        moonBody,
		Color:       color.RGBA{0x88, 0x88, 0x88, 0xff}, // gray colored
		IsAttractor: false,
		DrawRadius:  LogScalePlanetSize(moonBody.Radius),
	}
)

// log scales the size of the planet
func LogScalePlanetSize(radius float64) float64 {
	return math.Log(radius) / math.Log(1.5)
}
