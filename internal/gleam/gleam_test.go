package gleam_test

import (
	"testing"

	"github.com/salamer/zbpack/internal/gleam"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDockerfile(t *testing.T) {
	t.Parallel()

	dockerfile, err := gleam.GenerateDockerfile(map[string]string{})

	assert.NoError(t, err)
	assert.Contains(t, dockerfile, "\nWORKDIR /app")
}
