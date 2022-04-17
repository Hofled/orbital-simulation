//go:build ignore
// +build ignore

package shaders

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	origin, size := imageSrcRegionOnTexture()
	normalizedTexCoord := texCoord - origin
	normalizedTexCoord /= size
	// distance from center of texture in texels
	dist := normalizedTexCoord - 0.5
	// draws a circle that spans the entire rectangle's width and height of the texture it is drawn to
	val := 1.0 - smoothstep(
		0.99,
		1.01,
		dot(dist, dist)*4.0)
	return vec4(imageSrc0UnsafeAt(texCoord).rgb, val)
}
