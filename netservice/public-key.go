package netservice

import (
	"encoding/hex"
	"github.com/dfinity/keysmith/crypto"
	"github.com/tyler-smith/go-bip39"
)

func handlePublicKey(mnemonic string) (string, error) {
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
	output := hex.EncodeToString(grandchildECPubKey.SerializeUncompressed())
	return output, nil
}
