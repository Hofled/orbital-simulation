package physics

import (
	"image/color"
	"math"

	"gonum.org/v1/gonum/mat"
)

type Planet struct {
	Body  *Body
	Color color.RGBA
}

const (
	// this value is the inverse of the ratio between the real earth-to-moon distance and pixel distance
	EarthMoonScaleRatio = float64(1) / 3844
)

var (
	Earth *Planet = &Planet{
		Body:  NewBody(5.97237e24, 6356.752, mat.NewVecDense(2, []float64{100, 100})),
		Color: color.RGBA{0, 0, 0xff, 0xff}, // blue colored
	}
	Moon *Planet = &Planet{
		Body:  NewBody(7.342e22, 1738.1, mat.NewVecDense(2, []float64{200, 200})),
		Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, // white colored
	}
)

// log scales the size of the planet
func LogScalePlanetSize(radius float64) float64 {
	return math.Log(radius)
}
