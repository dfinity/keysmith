package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

const LEGACY_ADDRESS_CMD = "legacy-address"

type LegacyAddressCmd struct {
	FlagSet *flag.FlagSet
	Args    *LegacyAddressCmdArgs
}

type LegacyAddressCmdArgs struct {
	SeedFile  *string
	Index     *uint
	Protected *bool
}

func NewLegacyAddressCmd() *LegacyAddressCmd {
	fset := flag.NewFlagSet(LEGACY_ADDRESS_CMD, flag.ExitOnError)
	args := &LegacyAddressCmdArgs{
		SeedFile:  fset.String("f", "seed.txt", "Seed file."),
		Index:     fset.Uint("i", 0, "Child index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &LegacyAddressCmd{fset, args}
}

func (cmd *LegacyAddressCmd) Run() error {
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
	address := crypto.PubkeyToAddress(*childECPubKey.ToECDSA())
	output := strings.ToLower(strings.TrimPrefix(address.String(), "0x"))
	fmt.Println(output)
	return nil
}
