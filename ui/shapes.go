package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gonum.org/v1/gonum/mat"
)

// draws a circle on provided image
func Circle(img *ebiten.Image, x float32, y float32, r float32, c color.RGBA) {
	cimg := ebiten.NewImage(1, 1)
	cimg.Fill(c)

	circlePath := vector.Path{}
	circlePath.MoveTo(x, y)
	circlePath.Arc(x, y, r, 0, math.Pi*2, vector.Clockwise)

	vs, is := circlePath.AppendVerticesAndIndicesForFilling(nil, nil)
	op := &ebiten.DrawTrianglesOptions{FillRule: ebiten.EvenOdd}
	img.DrawTriangles(vs, is, cimg, op)
}

// draws an arrow from the x,y origin based on the vector values to a new
func DrawArrowTo(img *ebiten.Image, x, y float64, arrowBodyWidth int, destVec *mat.VecDense, c color.Color) {
	vecNorm := destVec.Norm(2)
	if math.IsNaN(vecNorm) {
		return
	}
	height := math.Ceil(vecNorm)
	arrowHeadWidth := int(arrowBodyWidth) * 4
	arrowHeadHeight := height * 0.25

	arrowImg := ebiten.NewImage(arrowHeadWidth, int(height))
	arrowImgOp := ebiten.DrawImageOptions{}

	// draw arrow line from bottom center to top center of image
	ebitenutil.DrawRect(arrowImg,
		float64((arrowHeadWidth/2)-(arrowBodyWidth/2)),
		float64(arrowHeadHeight),
		float64(arrowBodyWidth),
		float64(height),
		c)

	// draw arrow head
	DrawTriangle(arrowImg,
		mat.NewVecDense(2, []float64{float64(arrowHeadWidth / 2), 0}),
		mat.NewVecDense(2, []float64{float64(0), arrowHeadHeight}),
		mat.NewVecDense(2, []float64{float64(arrowHeadWidth), arrowHeadHeight}),
		c,
	)

	// flip sign of y value since screen coordinates or flipped
	vecAngle := math.Atan2(destVec.AtVec(0), -destVec.AtVec(1))

	// translate to origin of rotation for image (bottom middle)
	arrowImgOp.GeoM.Translate(float64(-arrowHeadWidth/2), float64(-height))

	// rotate arrow around arrow origin
	arrowImgOp.GeoM.Rotate(vecAngle)

	// translate arrow
	arrowImgOp.GeoM.Translate(x, y)

	img.DrawImage(arrowImg, &arrowImgOp)
}

func DrawTriangle(img *ebiten.Image, top, left, right *mat.VecDense, c color.Color) {
	cimg := ebiten.NewImage(1, 1)
	cimg.Fill(c)

	vectorEndpath := vector.Path{}
	vectorEndpath.MoveTo(float32(top.AtVec(0)), float32(top.AtVec(1)))
	vectorEndpath.LineTo(float32(left.AtVec(0)), float32(left.AtVec(1)))
	vectorEndpath.LineTo(float32(right.AtVec(0)), float32(right.AtVec(1)))
	vs, is := vectorEndpath.AppendVerticesAndIndicesForFilling(nil, nil)
	op := &ebiten.DrawTrianglesOptions{FillRule: ebiten.EvenOdd}

	img.DrawTriangles(vs, is, cimg, op)
}
