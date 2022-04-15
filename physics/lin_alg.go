package physics

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// returns the original vector, scaled to the specified length
func GetScaledVec(vec *mat.VecDense, length float64) *mat.VecDense {
	tmp := mat.NewVecDense(2, nil)
	tmp.ScaleVec(length/vec.Norm(2), vec)
	return tmp
}

func GetLogScaledVec(vec *mat.VecDense) *mat.VecDense {
	norm := vec.Norm(2)
	logNorm := math.Log(norm)
	return GetScaledVec(vec, logNorm)
}

func GetCurrentVelocity(b1 *Body, deltaTime, scaleFactor float64) *mat.VecDense {
	tmp := mat.NewVecDense(2, nil)
	tmp.ScaleVec(deltaTime*scaleFactor, b1.Velocity)
	return tmp
}
