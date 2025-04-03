package dockerfile

import (
	"testing"

	"github.com/salamer/zbpack/pkg/plan"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDockerFile(t *testing.T) {
	const expectedContent = "FROM alpine"

	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "Dockerfile", []byte(expectedContent), 0o644)

	config := plan.NewProjectConfigurationFromFs(fs, "")

	ctx := GetMeta(plan.NewPlannerOptions{Source: fs, Config: config})
	packer := NewPacker()
	actualContent, err := packer.GenerateDockerfile(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedContent, actualContent)
}

func TestNoMatchedDockerfile(t *testing.T) {
	fs := afero.NewMemMapFs()
	config := plan.NewProjectConfigurationFromFs(fs, "")

	ctx := GetMeta(plan.NewPlannerOptions{Source: fs, Config: config})

	assert.Equal(t, plan.Continue(), ctx)
}
