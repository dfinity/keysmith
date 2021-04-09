package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/enzoh/keysmith/account"
	"github.com/enzoh/keysmith/crypto"
	"github.com/enzoh/keysmith/seed"
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
	_, childECPubKey, err := crypto.DeriveChildECKeyPair(
		masterXPrivKey,
		[]uint32{0, uint32(*cmd.Args.Index)},
	)
	if err != nil {
		return err
	}
	accountId, err := account.FromECPubKey(childECPubKey)
	if err != nil {
		return err
	}
	fmt.Println(accountId.String())
	return nil
}
