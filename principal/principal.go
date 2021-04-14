package principal

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"hash/crc32"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/dfinity/keysmith/codec"
)

type PrincipalId interface {
	Bytes() []byte
	String() string
}

type SelfAuthenticating struct {
	data []byte
}

func NewSelfAuthenticating(der []byte) PrincipalId {
	hash := sha256.Sum224(der)
	data := append(hash[:], []byte{2}...)
	return &SelfAuthenticating{data: data}
}

func FromECPubKey(pubKey *btcec.PublicKey) (PrincipalId, error) {
	der, err := codec.EncodeECPubKey(pubKey)
	if err != nil {
		return nil, err
	}
	return NewSelfAuthenticating(der), nil
}

func (principal *SelfAuthenticating) Bytes() []byte {
	return principal.data
}

func (principal *SelfAuthenticating) String() string {
	return show(principal.data)
}

func show(data []byte) string {
	crc := make([]byte, 4)
	binary.BigEndian.PutUint32(crc, crc32.ChecksumIEEE(data))
	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	str := encoder.EncodeToString(append(crc, data...))
	return strings.Join(split(strings.ToLower(str), 5), "-")
}

func split(str string, n int) []string {
	if n >= len(str) {
		return []string{str}
	}
	var chunks []string
	chunk := make([]rune, n)
	i := 0
	for _, r := range str {
		chunk[i] = r
		i++
		if i == n {
			chunks = append(chunks, string(chunk))
			i = 0
		}
	}
	if i > 0 {
		chunks = append(chunks, string(chunk[:i]))
	}
	return chunks
}
