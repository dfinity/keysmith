package netservice

import (
	"github.com/dfinity/keysmith/crypto"
	"github.com/tyler-smith/go-bip39"
)

func handleXPublicKey(mnemonic string) (string, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return "", err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return "", err
	}
	masterXPubKey, err := masterXPrivKey.Neuter()
	if err != nil {
		return "", err
	}
	output := masterXPubKey.String()
	return output, nil
}
