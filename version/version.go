package version

import (
	"fmt"
)

type Version struct {
	major int
	minor int
	patch int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func NewVersion(major int, minor int, patch int) Version {
	return Version{major, minor, patch}
}

func VersionParse(version string) (Version, error) {
	tv := ""
	for _, c := range version {
		if (c >= '0' && c < '9') || c == '.' {
			tv = tv + string(c)
		}
	}
	version = tv

	var v Version
	_, err := fmt.Sscanf(version, "%d.%d.%d", &v.major, &v.minor, &v.patch)
	return v, err
}

func (v Version) Compare(other Version) int {
	if v.major > other.major {
		return 1
	} else if v.major < other.major {
		return -1
	} else {
		if v.minor > other.minor {
			return 1
		} else if v.minor < other.minor {
			return -1
		} else {
			if v.patch > other.patch {
				return 1
			} else if v.patch < other.patch {
				return -1
			} else {
				return 0
			}
		}
	}
}

func (v *Version) IsNewerThan(other *Version) bool {
	return v.Compare(*other) > 0
}
