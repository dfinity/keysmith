package main

import (
	"fmt"
	"os"
)

func main() {

	// Check if a subcommand was provided.
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, Usage())
		os.Exit(1)
	}

	// Run the subcommand.
	var err error
	switch os.Args[1] {
	case ACCOUNT_CMD:
		err = NewAccountCmd().Run()
	case GENERATE_CMD:
		err = NewGenerateCmd().Run()
	case LEGACY_ADDRESS_CMD:
		err = NewLegacyAddressCmd().Run()
	case PRINCIPAL_CMD:
		err = NewPrincipalCmd().Run()
	case PRIVATE_KEY_CMD:
		err = NewPrivateKeyCmd().Run()
	case PUBLIC_KEY_CMD:
		err = NewPublicKeyCmd().Run()
	case VERSION_CMD:
		err = NewVersionCmd().Run()
	case X_PUBLIC_KEY_CMD:
		err = NewXPublicKeyCmd().Run()
	default:
		fmt.Fprintf(os.Stderr, Usage())
		os.Exit(1)
	}

	// Check if an error occurred.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
