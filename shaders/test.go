//go:build ignore
// +build ignore

package shaders

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	baseColor := vec4(1, 0, 0, 1)

	origin, _ := imageSrcRegionOnTexture()
	distanceFromCenter := distance(origin, texCoord)
	baseColor.a = baseColor.a - distanceFromCenter
	return baseColor
}
