package netservice

import (
	"github.com/dfinity/keysmith/crypto"
	eth "github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"strings"
)

func handleLegacyAddress(mnemonic string) (string, error) {
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
	address := eth.PubkeyToAddress(*grandchildECPubKey.ToECDSA())
	output := strings.ToLower(strings.TrimPrefix(address.String(), "0x"))
	return output, nil
}
