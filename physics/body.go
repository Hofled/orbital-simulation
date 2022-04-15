package physics

import "gonum.org/v1/gonum/mat"

type Body struct {
	Mass   float64
	Radius float64
	// represents 2d velocity
	Velocity *mat.VecDense
	// position of the center of the mass
	Position *mat.VecDense
}

func NewBody(mass, radius float64, initialPos *mat.VecDense) *Body {
	return &Body{
		Mass:     mass,
		Radius:   radius,
		Position: initialPos,
		Velocity: mat.NewVecDense(2, []float64{0, 0}),
	}
}
