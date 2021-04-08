package main

import (
	"fmt"
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
		ACCOUNT_CMD,
		GENERATE_CMD,
		LEGACY_ADDRESS_CMD,
		PRINCIPAL_CMD,
		PRIVATE_KEY_CMD,
		PUBLIC_KEY_CMD,
		VERSION_CMD,
		X_PUBLIC_KEY_CMD,
	)
}
