package shaders

import (
	"embed"
)

//go:embed *
var ShadersFS embed.FS

func ReadShadersSource() (map[string][]byte, error) {
	circleShader, err := ShadersFS.ReadFile("circle.go")
	if err != nil {
		return nil, err
	}

	return map[string][]byte{
		"Circle": circleShader,
	}, nil
}
