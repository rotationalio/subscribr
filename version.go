package subscribr

import "fmt"

const (
	VersionMajor = 1
	VersionMinor = 0
	VersionPatch = 0
)

func Version() string {
	if VersionPatch > 0 {
		return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
	}
	return fmt.Sprintf("%d.%d", VersionMajor, VersionMinor)
}
