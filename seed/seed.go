package seed

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dfinity/keysmith/util"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/term"
)

func Load(seedFile string, protected bool) ([]byte, error) {
	mnemonic, err := readFileOrStdin(seedFile)
	if err != nil {
		return nil, err
	}
	var password []byte
	if protected {
		fmt.Printf("Password: ")
		password, err = term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println("")
		if err != nil {
			return nil, err
		}
	}
	mnemonic = bytes.TrimSuffix(mnemonic, []byte(util.NewLine))
	return bip39.NewSeedWithErrorChecking(
		string(mnemonic),
		string(password),
	)
}

func readFileOrStdin(file string) ([]byte, error) {
	reader := os.Stdin
	var err error
	if file != "-" {
		reader, err = os.Open(file)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
	}
	return ioutil.ReadAll(reader)
}
