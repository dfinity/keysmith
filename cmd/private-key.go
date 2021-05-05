package cmd

import (
	"flag"
	"os"

	"github.com/dfinity/keysmith/codec"
	"github.com/dfinity/keysmith/crypto"
	"github.com/dfinity/keysmith/seed"
)

const PRIVATE_KEY_CMD = "private-key"

type PrivateKeyCmd struct {
	FlagSet *flag.FlagSet
	Args    *PrivateKeyCmdArgs
}

type PrivateKeyCmdArgs struct {
	Index      *uint
	OutputFile *string
	Protected  *bool
	SeedFile   *string
}

func NewPrivateKeyCmd() *PrivateKeyCmd {
	fset := flag.NewFlagSet(PRIVATE_KEY_CMD, flag.ExitOnError)
	args := &PrivateKeyCmdArgs{
		Index:      fset.Uint("i", 0, "Child index."),
		OutputFile: fset.String("o", "identity.pem", "Output file."),
		Protected:  fset.Bool("p", false, "Password protection."),
		SeedFile:   fset.String("f", "seed.txt", "Seed file."),
	}
	return &PrivateKeyCmd{fset, args}
}

func (cmd *PrivateKeyCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	seed, err := seed.Load(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
	grandchildECPrivKey, _, err := crypto.DeriveGrandchildECKeyPair(
		masterXPrivKey,
		uint32(*cmd.Args.Index),
	)
	if err != nil {
		return err
	}
	output, err := codec.ECPrivKeyToPEM(grandchildECPrivKey)
	if err != nil {
		return err
	}
	return writeFileOrStdout(*cmd.Args.OutputFile, output, 0600)
}
