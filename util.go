package main

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"strings"

	"github.com/btcsuite/btcd/btcec"
)

func ECPubKeyToPrincipal(pubKey *btcec.PublicKey) (string, error) {
	der, err := EncodeECPubKey(pubKey)
	if err != nil {
		return "", err
	}
	return DisplayPrincipal(SelfAuthenticating(der)), nil
}

func ECPubKeyToAccount(pubKey *btcec.PublicKey) (string, error) {
	der, err := EncodeECPubKey(pubKey)
	if err != nil {
		return "", err
	}
	return DisplayAccount(DeriveAccount(der)), nil
}

func SelfAuthenticating(der []byte) []byte {
	hash := sha256.Sum224(der)
	tag := []byte{2}
	return append(hash[:], tag...)
}

func DeriveAccount(der []byte) []byte {
	hash := sha256.New224()
	hash.Write([]byte("\x0Aaccount-id"))
	hash.Write(SelfAuthenticating(der))
	hash.Write(make([]byte, 32))
	return hash.Sum(nil)
}

func DisplayPrincipal(data []byte) string {
	crc := make([]byte, 4)
	binary.BigEndian.PutUint32(crc, crc32.ChecksumIEEE(data))
	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	str := encoder.EncodeToString(append(crc, data...))
	return strings.Join(SplitN(strings.ToLower(str), 5), "-")
}

func DisplayAccount(data []byte) string {
	crc := make([]byte, 4)
	binary.BigEndian.PutUint32(crc, crc32.ChecksumIEEE(data))
	return hex.EncodeToString(append(crc, data...))
}

func SplitN(str string, n int) []string {
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
