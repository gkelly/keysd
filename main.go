package main

import (
	"github.com/keys-pub/keysd/service"
)

// build flags passed from goreleaser
var (
	version = "0.0.0-dev"
	commit  = "snapshot"
	date    = ""
)

func main() {
	build := service.Build{
		Version: version,
		Commit:  commit,
		Date:    date,
	}
	service.Run(build)
}