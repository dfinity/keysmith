package main

import (
	"fmt"
	"os"

	"github.com/dfinity/keysmith/cmd"
	"github.com/dfinity/keysmith/util"
)

var (
	MAJOR = "0"
	MINOR = "0"
	PATCH = "0"
	BUILD = "unknown"
)

func main() {
	filename := "seed.txt"
	// Check if a subcommand was provided.
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout, Usage())
		os.Exit(0)
	}

	if len(os.Args) > 2 {
		filename = os.Args[2]
	}
	// Run the subcommand.
	var err error
	switch os.Args[1] {
	case cmd.ACCOUNT_CMD:
		err = cmd.NewAccountCmd(filename).Run()
	case cmd.GENERATE_CMD:
		err = cmd.NewGenerateCmd(filename).Run()
	case cmd.LEGACY_ADDRESS_CMD:
		err = cmd.NewLegacyAddressCmd(filename).Run()
	case cmd.PRINCIPAL_CMD:
		err = cmd.NewPrincipalCmd(filename).Run()
	case cmd.PRIVATE_KEY_CMD:
		err = cmd.NewPrivateKeyCmd(filename).Run()
	case cmd.PUBLIC_KEY_CMD:
		err = cmd.NewPublicKeyCmd(filename).Run()
	case cmd.VERSION_CMD:
		err = cmd.NewVersionCmd(version()).Run()
	case cmd.X_PUBLIC_KEY_CMD:
		err = cmd.NewXPublicKeyCmd(filename).Run()
	default:
		fmt.Fprintf(os.Stderr, Usage())
		os.Exit(1)
	}

	// Check if an error occurred.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v"+util.NewLine, err)
		os.Exit(1)
	}
}

func version() string {
	return fmt.Sprintf("%s.%s.%s-%s", MAJOR, MINOR, PATCH, BUILD)
}
