package codec

import (
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/btcsuite/btcd/btcec"
)

func PEMToECPrivKey(data []byte) (*btcec.PrivateKey, error) {
	var block *pem.Block
	remainder := data
	for {
		block, remainder = pem.Decode(remainder)
		if block == nil {
			message := "Cannot find label \"EC PRIVATE KEY\" in PEM data"
			return nil, errors.New(message)
		}
		if block.Type == "EC PRIVATE KEY" {
			break
		}
	}
	var object ECPrivKey
	_, err := asn1.Unmarshal(block.Bytes, &object)
	if err != nil {
		return nil, err
	}
	if !IsSecp256k1(object.NamedCurveOID) {
		message := fmt.Sprintf(
			"Invalid curve type: %v",
			object.NamedCurveOID,
		)
		return nil, errors.New(message)
	}
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), object.PrivateKey)
	return privKey, nil
}

func LoadECPrivKey(identityFile string) (*btcec.PrivateKey, error) {
	contents, err := ioutil.ReadFile(identityFile)
	if err != nil {
		return nil, err
	}
	privKey, err := PEMToECPrivKey(contents)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}
