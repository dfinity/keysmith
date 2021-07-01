package cmd

import (
	"fmt"
	"strings"
)

const SHORTLIST_CMD = "shortlist"

type ShortlistCmd struct {
}

func NewShortlistCmd() *ShortlistCmd {
	return &ShortlistCmd{}
}

func (cmd *ShortlistCmd) Run() error {
	output := strings.Join([]string{
		ACCOUNT_CMD,
		GENERATE_CMD,
		LEGACY_ADDRESS_CMD,
		PRINCIPAL_CMD,
		PRIVATE_KEY_CMD,
		PUBLIC_KEY_CMD,
		SHORTLIST_CMD,
		VERSION_CMD,
		X_PRIVATE_KEY_CMD,
		X_PUBLIC_KEY_CMD,
	}, " ")
	fmt.Println(output)
	return nil
}
