//go:build mage
// +build mage

package main

import "github.com/magefile/mage/sh"

func Gen() error {
	return sh.Run(
		"protoc",
		"--go_out=paths=source_relative:.",
		"--go-grpc_out=paths=source_relative:.",
		"stringsvc/transport/pb/stringsvc.proto",
	)
}

func Tidy() error {
	return sh.Run(
		"go",
		"mod",
		"tidy",
	)
}

func Run() error {
	return sh.Run(
		"go",
		"run",
		"stringsvc/cmd/main.go",
	)
}
