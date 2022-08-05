package examples

import (
	"fmt"

	"github.com/guionardo/go-gstools/version"
)

func VersionExample() {
	v1, _ := version.VersionParse("1.0.0")
	v2, _ := version.VersionParse("v1.2.0")

	fmt.Println(v1.Compare(v2))	// -1
}
