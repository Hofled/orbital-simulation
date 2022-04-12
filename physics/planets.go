package physics

import "gonum.org/v1/gonum/mat"

var (
	Earth *Body = NewBody(5.97237e24, 6356.752, *mat.NewVecDense(2, []float64{100, 100}))
	Moon  *Body = NewBody(7.342e22, 1738.1, *mat.NewVecDense(2, []float64{200, 200}))
)
