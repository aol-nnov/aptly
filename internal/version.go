//go:generate go run git.rootprojects.org/root/go-gitver/v2 --package internal
package internal

import (
	"fmt"
)

var (
	commit  = "$Format:%H$"
	version = "$Format:%(describe)$"
	date    = ""
)

type VersionInfo struct {
	Version string
	Commit  string
}

func Version() VersionInfo {
	return VersionInfo{
		Version: version,
		Commit:  commit,
	}
}

func (v VersionInfo) String() string {
	return v.Version
}

func (v VersionInfo) Print() {
	fmt.Println("aptly version:", v.Version)
	// fmt.Println()

	// fmt.Println("Build information:")
	// fmt.Printf("  Go version: %s (%s, %s)\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	// fmt.Println("  Git commit:", v.Commit)
}
