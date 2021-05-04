package netservice

import (
	"github.com/tyler-smith/go-bip39"
)

func handleGenerate() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)

	if err != nil {
		return "", err
	}
	return mnemonic, err
}
