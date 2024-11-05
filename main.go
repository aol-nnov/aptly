//go:generate go run github.com/swaggo/swag/cmd/swag init --markdownFiles docs
package main

import (
	"os"

	"github.com/aptly-dev/aptly/cmd"
)

func main() {
	os.Exit(cmd.Run(cmd.RootCommand(), os.Args[1:], true))
}
