package main

import (
	"fmt"
)

const (
	VERSION_CMD = "version"

	MAJOR = 1
	MINOR = 2
	PATCH = 0
)

var RELEASE = "unknown"

type VersionCmd struct {
}

func NewVersionCmd() *VersionCmd {
	return &VersionCmd{}
}

func (cmd *VersionCmd) Run() error {
	fmt.Printf("%d.%d.%d-%s\n", MAJOR, MINOR, PATCH, RELEASE)
	return nil
}
