package geometry

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Rotate(point, pivot *mat.VecDense, angle float64) *mat.VecDense {
	c := math.Cos(angle)
	s := math.Sin(angle)

	// translate point to be relative to origin same as it was with pivot
	point.SubVec(point, pivot.TVec())

	newPoint := mat.NewVecDense(2, []float64{
		point.AtVec(0)*c - point.AtVec(1)*s,
		point.AtVec(0)*s + point.AtVec(1)*c})

	// translate new point back
	point.AddVec(newPoint.TVec(), pivot.TVec())

	return point
}
