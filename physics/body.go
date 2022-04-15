package physics

import "gonum.org/v1/gonum/mat"

type Body struct {
	// mass represented in kilograms
	Mass float64
	// radius represented in kilometers
	Radius float64
	// represents 2d velocity in Newtons
	Velocity *mat.VecDense
	// position of the center of the mass
	Position *mat.VecDense
}

func NewBody(mass, radius float64, initialPos, initialVelocity *mat.VecDense) *Body {
	return &Body{
		Mass:     mass,
		Radius:   radius,
		Position: initialPos,
		Velocity: initialVelocity,
	}
}
