package shaders

import (
	"embed"
)

//go:embed *
var ShadersFS embed.FS

func ReadShadersSource() (map[string][]byte, error) {
	testShader, err := ShadersFS.ReadFile("test.go")
	if err != nil {
		return nil, err
	}

	return map[string][]byte{
		"Test": testShader,
	}, nil
}
