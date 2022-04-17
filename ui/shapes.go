package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gonum.org/v1/gonum/mat"
)

var circleImg = ebiten.NewImage(1, 1)

// draws a circle on provided image
func Circle(img *ebiten.Image, x float32, y float32, r float32, c color.RGBA) {
	circleImg.Fill(c)

	circlePath := vector.Path{}
	circlePath.MoveTo(x, y)
	circlePath.Arc(x, y, r, 0, math.Pi*2, vector.Clockwise)

	vs, is := circlePath.AppendVerticesAndIndicesForFilling(nil, nil)
	op := &ebiten.DrawTrianglesOptions{FillRule: ebiten.EvenOdd}
	img.DrawTriangles(vs, is, circleImg, op)
}

const (
	minArrowHeadHeight = 5
)

func DrawCircle(src, dst *ebiten.Image, x, y float64, c color.Color, shaderMap map[string]*ebiten.Shader) {
	shader := shaderMap["Circle"]
	src.Fill(c)
	shaderOp := &ebiten.DrawRectShaderOptions{}
	shaderOp.GeoM.Translate(x, y)
	shaderOp.Uniforms = map[string]interface{}{
		"Offset": []float32{float32(x), float32(y)},
	}
	shaderOp.Images[0] = src
	w, h := src.Size()
	dst.DrawRectShader(w, h, shader, shaderOp)
}

// draws an arrow from the x,y origin based on the vector values to a new
func DrawArrowTo(dst *ebiten.Image, x, y, arrowLen float64, arrowBodyWidth int, destVec *mat.VecDense, c color.Color) {
	height := math.Ceil(arrowLen)
	arrowHeadWidth := int(arrowBodyWidth) * 4
	arrowHeadHeight := math.Max(minArrowHeadHeight, height*0.25)

	arrowImg := ebiten.NewImage(arrowHeadWidth, int(height))
	arrowImgOp := ebiten.DrawImageOptions{}

	// draw arrow line from bottom center to top center of image
	ebitenutil.DrawRect(arrowImg,
		float64((arrowHeadWidth/2)-(arrowBodyWidth/2)),
		arrowHeadHeight,
		float64(arrowBodyWidth),
		float64(height),
		c)

	// draw arrow head
	DrawTriangle(arrowImg,
		float32(arrowHeadWidth)*0.5, 0,
		0, float32(arrowHeadHeight),
		float32(arrowHeadWidth), float32(arrowHeadHeight),
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

	dst.DrawImage(arrowImg, &arrowImgOp)
}

var triangleImg = ebiten.NewImage(1, 1)

func DrawTriangle(dst *ebiten.Image, topX, topY, leftX, leftY, rightX, rightY float32, c color.Color) {
	triangleImg.Fill(c)

	triangleVertices := []ebiten.Vertex{
		{
			DstX:   topX,
			DstY:   topY,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1},
		{
			DstX:   leftX,
			DstY:   leftY,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1},
		{
			DstX:   rightX,
			DstY:   rightY,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1},
	}
	triangleIndices := []uint16{0, 1, 2}
	op := &ebiten.DrawTrianglesOptions{FillRule: ebiten.FillAll}

	dst.DrawTriangles(triangleVertices, triangleIndices, triangleImg, op)
}
