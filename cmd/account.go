package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dfinity/keysmith/account"
	"github.com/dfinity/keysmith/crypto"
	"github.com/dfinity/keysmith/seed"
)

const ACCOUNT_CMD = "account"

type AccountCmd struct {
	FlagSet *flag.FlagSet
	Args    *AccountCmdArgs
}

type AccountCmdArgs struct {
	SeedFile  *string
	Index     *uint
	Protected *bool
}

func NewAccountCmd() *AccountCmd {
	fset := flag.NewFlagSet(ACCOUNT_CMD, flag.ExitOnError)
	args := &AccountCmdArgs{
		SeedFile:  fset.String("f", "seed.txt", "Seed file."),
		Index:     fset.Uint("i", 0, "Child index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &AccountCmd{fset, args}
}

func (cmd *AccountCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	seed, err := seed.Load(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
	childXPrivKey, err := masterXPrivKey.Derive(0)
	if err != nil {
		return err
	}
	_, grandchildECPubKey, err := crypto.DeriveChildECKeyPair(
		childXPrivKey,
		uint32(*cmd.Args.Index),
	)
	if err != nil {
		return err
	}
	accountId, err := account.FromECPubKey(grandchildECPubKey)
	if err != nil {
		return err
	}
	fmt.Println(accountId.String())
	return nil
}
