//go:build tools

package internal

import (
	_ "git.rootprojects.org/root/go-gitver/v2"
	_ "github.com/air-verse/air"
	_ "github.com/choffmeister/git-describe-semver"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/swaggo/swag/cmd/swag"
)
