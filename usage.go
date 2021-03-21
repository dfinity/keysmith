package main

import (
	"fmt"
)

func Usage() string {
	return fmt.Sprintf(`usage: keysmith <command> [<args>]

Available Commands:
    %s            Generate your mnemonic seed.
    %s         Write your private key to a file.
    %s        Print your extended public key.
    %s          Print your public key.
    %s             Print your legacy address.
    %s           Print your principal identifier.
`,
		GENERATE_CMD,
		PRIVATE_KEY_CMD,
		X_PUBLIC_KEY_CMD,
		PUBLIC_KEY_CMD,
		ADDRESS_CMD,
		PRINCIPAL_CMD,
	)
}
