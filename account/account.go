package account

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"

	"github.com/btcsuite/btcd/btcec"
	"github.com/dfinity/keysmith/codec"
	"github.com/dfinity/keysmith/principal"
)

type AccountId interface {
	String() string
}

type accountId struct {
	data []byte
}

func FromECPubKey(pubKey *btcec.PublicKey) (AccountId, error) {
	der, err := codec.EncodeECPubKey(pubKey)
	if err != nil {
		return nil, err
	}
	hash := sha256.New224()
	hash.Write([]byte("\x0Aaccount-id"))
	hash.Write(principal.NewSelfAuthenticating(der).Bytes())
	hash.Write(make([]byte, 32))
	data := hash.Sum(nil)
	return &accountId{data: data}, nil
}

func (accountId *accountId) String() string {
	crc := make([]byte, 4)
	binary.BigEndian.PutUint32(crc, crc32.ChecksumIEEE(accountId.data))
	return hex.EncodeToString(append(crc, accountId.data...))
}
