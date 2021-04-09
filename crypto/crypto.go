package crypto

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
)

func DeriveMasterXPrivKey(seed []byte) (*hdkeychain.ExtendedKey, error) {
	masterXPrivKey, err := hdkeychain.NewMaster(
		seed,
		&chaincfg.MainNetParams,
	)
	if err != nil {
		return nil, err
	}
	path, err := accounts.ParseDerivationPath("m/44'/223'/0'")
	if err != nil {
		return nil, err
	}
	for _, i := range path {
		masterXPrivKey, err = masterXPrivKey.Child(i)
		if err != nil {
			return nil, err
		}
	}
	return masterXPrivKey, nil
}

func DeriveChildECKeyPair(
	masterXPrivKey *hdkeychain.ExtendedKey,
	path []uint32,
) (*btcec.PrivateKey, *btcec.PublicKey, error) {
	childXPrivKey := masterXPrivKey
	var err error
	for _, i := range path {
		childXPrivKey, err = childXPrivKey.Child(i)
		if err != nil {
			return nil, nil, err
		}
	}
	childECPrivKey, err := childXPrivKey.ECPrivKey()
	if err != nil {
		return nil, nil, err
	}
	childECPubKey := childECPrivKey.PubKey()
	return childECPrivKey, childECPubKey, nil
}
