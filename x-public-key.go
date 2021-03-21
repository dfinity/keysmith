package main

import (
	"flag"
	"fmt"
	"os"
)

const X_PUBLIC_KEY_CMD = "x-public-key"

type XPublicKeyCmd struct {
	FlagSet *flag.FlagSet
	Args    *XPublicKeyCmdArgs
}

type XPublicKeyCmdArgs struct {
	SeedFile  *string
	Protected *bool
}

func NewXPublicKeyCmd() *XPublicKeyCmd {
	fset := flag.NewFlagSet(PUBLIC_KEY_CMD, flag.ExitOnError)
	args := &XPublicKeyCmdArgs{
		SeedFile:  fset.String("f", "seed.txt", "Seed file."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &XPublicKeyCmd{fset, args}
}

func (cmd *XPublicKeyCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	seed, err := LoadSeed(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
	masterXPubKey, err := masterXPrivKey.Neuter()
	if err != nil {
		return err
	}
	output := masterXPubKey.String()
	fmt.Println(output)
	return nil
}
