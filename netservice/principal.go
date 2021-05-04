package netservice

import (
	"github.com/dfinity/keysmith/crypto"
	"github.com/dfinity/keysmith/principal"
	"github.com/tyler-smith/go-bip39"
)

func handlePrincipal(mnemonic string) (string, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return "", err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return "", err
	}
	_, grandchildECPubKey, err := crypto.DeriveGrandchildECKeyPair(
		masterXPrivKey,
		0,
	)
	if err != nil {
		return "", err
	}
	principalId, err := principal.FromECPubKey(grandchildECPubKey)
	if err != nil {
		return "", err
	}
	return principalId.String(), nil
}
