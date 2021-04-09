package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/enzoh/keysmith/codec"
	"github.com/enzoh/keysmith/crypto"
	"github.com/enzoh/keysmith/seed"
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
	_, err := os.Stat(*cmd.Args.OutputFile)
	if !os.IsNotExist(err) {
		return fmt.Errorf(
			"Output file already exists: %s",
			*cmd.Args.OutputFile,
		)
	}
	seed, err := seed.Load(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
	childECPrivKey, _, err := crypto.DeriveChildECKeyPair(
		masterXPrivKey,
		[]uint32{0, uint32(*cmd.Args.Index)},
	)
	if err != nil {
		return err
	}
	output, err := codec.ECPrivKeyToPEM(childECPrivKey)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(*cmd.Args.OutputFile, output, 0600)
}
