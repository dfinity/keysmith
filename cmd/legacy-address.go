package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/enzoh/keysmith/crypto"
	"github.com/enzoh/keysmith/seed"
	eth "github.com/ethereum/go-ethereum/crypto"
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
	seed, err := seed.Load(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
	_, childECPubKey, err := crypto.DeriveChildECKeyPair(
		masterXPrivKey,
		[]uint32{0, uint32(*cmd.Args.Index)},
	)
	if err != nil {
		return err
	}
	address := eth.PubkeyToAddress(*childECPubKey.ToECDSA())
	output := strings.ToLower(strings.TrimPrefix(address.String(), "0x"))
	fmt.Println(output)
	return nil
}
