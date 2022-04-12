package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Circle draws a circle on provided image
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
