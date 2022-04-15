package physics

import (
	"math"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/unit/constant"
)

const (
	// 1 pixel = 10,000 meters = 10 kilometers
	metersToPixelRatio = 0.0001
)

// calculates the gravitational force applied on b2 exerted by b1
// based on Newton's law of universal gravitation
// https://en.wikipedia.org/wiki/Newton%27s_law_of_universal_gravitation#Vector_form
// mass is measured by kilograms, and distance by meters
func Gravitation(b1, b2 *Body) *mat.VecDense {
	gravityVecDense := mat.NewVecDense(2, nil)
	gravityVecDense.SubVec(b2.Position.TVec(), b1.Position.TVec())
	distanceNorm := gravityVecDense.Norm(2)
	// convert pixels to km
	distanceNorm /= metersToPixelRatio
	// normalize the distance vector
	gravityVecDense.ScaleVec(1/distanceNorm, gravityVecDense.TVec())
	coefficient := -float64(constant.Gravitational) * (b1.Mass * b2.Mass) / (math.Pow(distanceNorm, 2))
	gravityVecDense.ScaleVec(coefficient, gravityVecDense.TVec())
	return gravityVecDense
}

func ApplyForce(body *Body, force *mat.VecDense, deltaTime float64) {
	body.Velocity.AddScaledVec(body.Velocity, deltaTime/body.Mass, force)
}

func ApplyMovement(body *Body, deltaTime float64) {
	body.Position.AddScaledVec(body.Position, deltaTime, body.Velocity)
}
