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

var (
	Earth *Planet = &Planet{
		Body:  NewBody(5.97237e24, 6356.752, mat.NewVecDense(2, []float64{250, 250}), mat.NewVecDense(2, []float64{0, 0})),
		Color: color.RGBA{0, 0, 0xff, 0xff}, // blue colored
	}
	Moon *Planet = &Planet{
		Body:  NewBody(7.342e22, 1738.1, mat.NewVecDense(2, []float64{289, 289}), mat.NewVecDense(2, []float64{2, -2})),
		Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, // white colored
	}
)

// log scales the size of the planet
func LogScalePlanetSize(radius float64) float64 {
	return math.Log(radius) / math.Log(1.6)
}
