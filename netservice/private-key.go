package netservice

import (
	"github.com/dfinity/keysmith/codec"
	"github.com/dfinity/keysmith/crypto"
	"github.com/tyler-smith/go-bip39"
)

func handlePrivateKey(mnemonic string) (string, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return "", err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return "", err
	}
	grandchildECPrivKey, _, err := crypto.DeriveGrandchildECKeyPair(
		masterXPrivKey,
		0,
	)
	if err != nil {
		return "", err
	}
	output, err := codec.ECPrivKeyToPEM(grandchildECPrivKey)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
