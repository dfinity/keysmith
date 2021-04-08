package main

import (
	"fmt"
)

const VERSION_CMD = "version"

var (
	MAJOR = 1
	MINOR = 2
	PATCH = 0
	BUILD = "unknown"
)

type VersionCmd struct {
}

func NewVersionCmd() *VersionCmd {
	return &VersionCmd{}
}

func (cmd *VersionCmd) Run() error {
	fmt.Printf("%d.%d.%d-%s\n", MAJOR, MINOR, PATCH, BUILD)
	return nil
}
