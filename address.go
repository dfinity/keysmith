package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

const ADDRESS_CMD = "address"

type AddressCmd struct {
	FlagSet *flag.FlagSet
	Args    *AddressCmdArgs
}

type AddressCmdArgs struct {
	SeedFile  *string
	Index     *uint
	Protected *bool
}

func NewAddressCmd() *AddressCmd {
	fset := flag.NewFlagSet(ADDRESS_CMD, flag.ExitOnError)
	args := &AddressCmdArgs{
		SeedFile:  fset.String("f", "seed.txt", "Seed file."),
		Index:     fset.Uint("i", 0, "Child index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &AddressCmd{fset, args}
}

func (cmd *AddressCmd) Run() error {
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
	childXPubKey, err := childXPrivKey.Neuter()
	if err != nil {
		return err
	}
	childECPubKey, err := childXPubKey.ECPubKey()
	if err != nil {
		return err
	}
	address := crypto.PubkeyToAddress(*childECPubKey.ToECDSA())
	output := strings.ToLower(strings.TrimPrefix(address.String(), "0x"))
	fmt.Println(output)
	return nil
}
