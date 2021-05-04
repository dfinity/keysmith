package netservice

import (
	"github.com/dfinity/keysmith/account"
	"github.com/dfinity/keysmith/crypto"
	"github.com/tyler-smith/go-bip39"
)

func handleAccount(mnemonic string) (string, error) {
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
		uint32(0),
	)
	if err != nil {
		return "", err
	}
	accountId, err := account.FromECPubKey(grandchildECPubKey)
	if err != nil {
		return "", err
	}
	return accountId.String(), nil
}
