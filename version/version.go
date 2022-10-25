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
	n1 := v.major*100000 + v.minor*1000 + v.patch
	n2 := other.major*100000 + other.minor*1000 + other.patch
	if n1 > n2 {
		return 1
	}
	if n1 < n2 {
		return -1
	}
	return 0
}

func (v *Version) IsNewerThan(other *Version) bool {
	return v.Compare(*other) > 0
}

func (v *Version) IsEqual(other *Version) bool {
	return v.Compare(*other) == 0
}
