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
	path := []uint32{0, uint32(*cmd.Args.Index)}
	_, childECPubKey, err := DeriveChildECKeyPair(seed, path)
	if err != nil {
		return err
	}
	output, err := ECPubKeyToPrincipal(childECPubKey)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
