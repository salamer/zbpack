package python

import (
	"github.com/salamer/zbpack/internal/utils"
)

const defaultPython3Version = "3.10"

func getPython3Version(versionRange string) string {
	return utils.ConstraintToVersion(versionRange, defaultPython3Version)
}
