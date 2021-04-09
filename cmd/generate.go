package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tyler-smith/go-bip39"
)

const GENERATE_CMD = "generate"

type GenerateCmd struct {
	FlagSet *flag.FlagSet
	Args    *GenerateCmdArgs
}

type GenerateCmdArgs struct {
	OutputFile *string
}

func NewGenerateCmd() *GenerateCmd {
	fset := flag.NewFlagSet(GENERATE_CMD, flag.ExitOnError)
	args := &GenerateCmdArgs{
		OutputFile: fset.String("o", "seed.txt", "Seed file."),
	}
	return &GenerateCmd{fset, args}
}

func (cmd *GenerateCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	_, err := os.Stat(*cmd.Args.OutputFile)
	if !os.IsNotExist(err) {
		return fmt.Errorf(
			"Output file already exists: %s",
			*cmd.Args.OutputFile,
		)
	}
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return err
	}
	contents := []byte(mnemonic + "\n")
	return ioutil.WriteFile(*cmd.Args.OutputFile, contents, 0600)
}
