package main

import (
	"flag"
	"fmt"
	"os"
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

func NewPrincipalCmd() *PrincipalCmd {
	fset := flag.NewFlagSet(PRINCIPAL_CMD, flag.ExitOnError)
	args := &PrincipalCmdArgs{
		SeedFile:  fset.String("f", "seed.txt", "Seed file."),
		Index:     fset.Uint("i", 0, "Child index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &PrincipalCmd{fset, args}
}

func (cmd *PrincipalCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	seed, err := LoadSeed(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
	masterXPrivKey0, err := masterXPrivKey.Child(0)
	if err != nil {
		return err
	}
	childXPrivKey, err := masterXPrivKey0.Child(uint32(*cmd.Args.Index))
	if err != nil {
		return err
	}
	childECPrivKey, err := childXPrivKey.ECPrivKey()
	if err != nil {
		return err
	}
	childECPubKey := childECPrivKey.PubKey()
	output, err := ECPubKeyToPrincipal(childECPubKey)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
