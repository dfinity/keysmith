package cmd

import (
	"flag"
	"os"

	"github.com/dfinity/keysmith/util"
	"github.com/tyler-smith/go-bip39"
)

const GENERATE_CMD = "generate"

type GenerateCmd struct {
	FlagSet *flag.FlagSet
	Args    *GenerateCmdArgs
}

type GenerateCmdArgs struct {
	Bits       *int
	OutputFile *string
}

func NewGenerateCmd() *GenerateCmd {
	fset := flag.NewFlagSet(GENERATE_CMD, flag.ExitOnError)
	args := &GenerateCmdArgs{
		Bits:       fset.Int("b", 128, "Bits of entropy."),
		OutputFile: fset.String("o", "seed.txt", "Seed file."),
	}
	return &GenerateCmd{fset, args}
}

func (cmd *GenerateCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	entropy, err := bip39.NewEntropy(*cmd.Args.Bits)
	if err != nil {
		return err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return err
	}
	output := []byte(mnemonic + util.NewLine)
	return writeFileOrStdout(*cmd.Args.OutputFile, output, 0600)
}
