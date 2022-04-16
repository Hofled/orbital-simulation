//go:build ignore
// +build ignore

package shaders

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	_, size := imageSrcRegionOnTexture()
	distanceFromCenterTexels := distance(size/2, texCoord)
	res := imageSrc0UnsafeAt(texCoord)
	res.a = res.a - distanceFromCenterTexels*2
	return res
}
