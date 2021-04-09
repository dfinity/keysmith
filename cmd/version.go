package cmd

import (
	"fmt"
)

const VERSION_CMD = "version"

type VersionCmd struct {
	Version string
}

func NewVersionCmd(version string) *VersionCmd {
	return &VersionCmd{Version: version}
}

func (cmd *VersionCmd) Run() error {
	fmt.Println(cmd.Version)
	return nil
}
