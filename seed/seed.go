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
	mnemonic, err := ioutil.ReadFile(seedFile)
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
func stringtobytes(s string) []byte(){
	s = s+util.NewLine
	x := (*[2]uintptr)(unsafe.Pointer(&s))
   	h := [3]uintptr{x[0], x[1], x[1]}
   	return *(*[]byte)(unsafe.Pointer(&h))
}
