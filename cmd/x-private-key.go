package cmd

import (
	"flag"
	"os"

	"github.com/dfinity/keysmith/crypto"
	"github.com/dfinity/keysmith/seed"
	"github.com/dfinity/keysmith/util"
)

const X_PRIVATE_KEY_CMD = "x-private-key"

type XPrivateKeyCmd struct {
	FlagSet *flag.FlagSet
	Args    *XPrivateKeyCmdArgs
}

type XPrivateKeyCmdArgs struct {
	OutputFile *string
	SeedFile   *string
	Protected  *bool
}

func NewXPrivateKeyCmd() *XPrivateKeyCmd {
	fset := flag.NewFlagSet(X_PRIVATE_KEY_CMD, flag.ExitOnError)
	args := &XPrivateKeyCmdArgs{
		OutputFile: fset.String("o", "identity.bip32", "Output file."),
		SeedFile:   fset.String("f", "seed.txt", "Seed file."),
		Protected:  fset.Bool("p", false, "Password protection."),
	}
	return &XPrivateKeyCmd{fset, args}
}

func (cmd *XPrivateKeyCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	seed, err := seed.Load(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
	output := []byte(masterXPrivKey.String() + util.NewLine)
	return writeFileOrStdout(*cmd.Args.OutputFile, output, 0600)
}
