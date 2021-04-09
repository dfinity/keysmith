package main

import (
	"fmt"

	"github.com/enzoh/keysmith/cmd"
)

func Usage() string {
	return fmt.Sprintf(`usage: keysmith <command> [<args>]

Available Commands:
    %s             Print your account identifier.
    %s            Generate your mnemonic seed.
    %s      Print your legacy address.
    %s           Print your principal identifier.
    %s         Write your private key to a file.
    %s          Print your public key.
    %s             Print the version number.
    %s        Print your extended public key.
`,
		cmd.ACCOUNT_CMD,
		cmd.GENERATE_CMD,
		cmd.LEGACY_ADDRESS_CMD,
		cmd.PRINCIPAL_CMD,
		cmd.PRIVATE_KEY_CMD,
		cmd.PUBLIC_KEY_CMD,
		cmd.VERSION_CMD,
		cmd.X_PUBLIC_KEY_CMD,
	)
}
