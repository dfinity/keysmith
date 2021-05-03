package cmd

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/dfinity/keysmith/crypto"
	"github.com/dfinity/keysmith/seed"
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

func NewPublicKeyCmd(filename string) *PublicKeyCmd {
	fset := flag.NewFlagSet(PUBLIC_KEY_CMD, flag.ExitOnError)
	args := &PublicKeyCmdArgs{
		SeedFile:  fset.String("f", filename, "Seed file."),
		Index:     fset.Uint("i", 0, "Child index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &PublicKeyCmd{fset, args}
}

func (cmd *PublicKeyCmd) Run() error {
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
	output := hex.EncodeToString(grandchildECPubKey.SerializeUncompressed())
	fmt.Println(output)
	return nil
}
