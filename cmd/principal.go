package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dfinity/keysmith/crypto"
	"github.com/dfinity/keysmith/principal"
	"github.com/dfinity/keysmith/seed"
)

const PRINCIPAL_CMD = "principal"

type PrincipalCmd struct {
	FlagSet *flag.FlagSet
	Args    *PrincipalCmdArgs
}

type PrincipalCmdArgs struct {
	SeedFile  *string
	Index     *uint
	Protected *bool
}

func NewPrincipalCmd(filename string) *PrincipalCmd {
	fset := flag.NewFlagSet(PRINCIPAL_CMD, flag.ExitOnError)
	args := &PrincipalCmdArgs{
		SeedFile:  fset.String("f", filename, "Seed file."),
		Index:     fset.Uint("i", 0, "Child index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &PrincipalCmd{fset, args}
}

func (cmd *PrincipalCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	seed, err := seed.Load(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
	_, grandchildECPubKey, err := crypto.DeriveGrandchildECKeyPair(
		masterXPrivKey,
		uint32(*cmd.Args.Index),
	)
	if err != nil {
		return err
	}
	principalId, err := principal.FromECPubKey(grandchildECPubKey)
	if err != nil {
		return err
	}
	fmt.Println(principalId.String())
	return nil
}
