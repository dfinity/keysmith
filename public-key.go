package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

const PUBLIC_KEY_CMD = "public-key"

type PublicKeyCmd struct {
	FlagSet *flag.FlagSet
	Args    *PublicKeyCmdArgs
}

type PublicKeyCmdArgs struct {
	SeedFile  *string
	Index     *uint
	Protected *bool
}

func NewPublicKeyCmd() *PublicKeyCmd {
	fset := flag.NewFlagSet(PUBLIC_KEY_CMD, flag.ExitOnError)
	args := &PublicKeyCmdArgs{
		SeedFile:  fset.String("f", "seed.txt", "Seed file."),
		Index:     fset.Uint("i", 0, "Child index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &PublicKeyCmd{fset, args}
}

func (cmd *PublicKeyCmd) Run() error {
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
	output := hex.EncodeToString(childECPubKey.SerializeUncompressed())
	fmt.Println(output)
	return nil
}
