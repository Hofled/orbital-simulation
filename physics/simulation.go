package physics

import (
	"math"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/unit/constant"
)

// calculates the gravitational force applied on b2 exerted by b1
// based on Newton's law of universal gravitation
// https://en.wikipedia.org/wiki/Newton%27s_law_of_universal_gravitation#Vector_form
func Gravitation(b1, b2 *Body) *mat.VecDense {
	distanceVecDense := mat.NewVecDense(2, nil)
	// calculate vector from b1 to b2
	distanceVecDense.SubVec(b2.Position.TVec(), b1.Position.TVec())
	distanceNorm := distanceVecDense.Norm(2)
	// normalize the distance vector
	distanceVecDense.ScaleVec(1/distanceNorm, distanceVecDense.TVec())
	coefficient := -float64(constant.Gravitational) * (b1.Mass * b2.Mass) / (math.Pow(distanceNorm, 2))
	distanceVecDense.ScaleVec(coefficient, distanceVecDense.TVec())
	return distanceVecDense
}