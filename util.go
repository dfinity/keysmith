package main

import (
	"bytes"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base32"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/term"
)

type ECPrivKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

type ECPubKeyMetadata struct {
	ECPubKeyOID   asn1.ObjectIdentifier
	NamedCurveOID asn1.ObjectIdentifier
}

type ECPubKey struct {
	Metadata  ECPubKeyMetadata
	PublicKey asn1.BitString
}

func LoadSeed(seedFile string, protected bool) ([]byte, error) {
	mnemonic, err := ioutil.ReadFile(seedFile)
	if err != nil {
		return nil, err
	}
	var password []byte
	if protected {
		fmt.Printf("Password: ")
		password, err = term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println("")
		if err != nil {
			return nil, err
		}
	}
	mnemonic = bytes.TrimSuffix(mnemonic, []byte("\n"))
	return bip39.NewSeedWithErrorChecking(string(mnemonic), string(password))
}

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

func ECPrivKeyToPEM(privKey *btcec.PrivateKey) ([]byte, error) {
	der1, err := EncodeECParams()
	if err != nil {
		return nil, err
	}
	der2, err := EncodeECPrivKey(privKey)
	if err != nil {
		return nil, err
	}
	block1 := &pem.Block{
		Type:  "EC PARAMETERS",
		Bytes: der1,
	}
	block2 := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der2,
	}
	data1 := pem.EncodeToMemory(block1)
	data2 := pem.EncodeToMemory(block2)
	return append(data1, data2...), nil
}

func ECPubKeyToPrincipal(pubKey *btcec.PublicKey) (string, error) {
	der, err := EncodeECPubKey(pubKey)
	if err != nil {
		return "", err
	}
	return SelfAuthenticating(der), nil
}

func EncodeECParams() ([]byte, error) {
	return asn1.Marshal(SECP256K1())
}

func EncodeECPrivKey(privKey *btcec.PrivateKey) ([]byte, error) {
	curve := btcec.S256()
	point := privKey.PubKey().ToECDSA()
	return asn1.Marshal(ECPrivKey{
		Version:       1,
		PrivateKey:    privKey.D.Bytes(),
		NamedCurveOID: SECP256K1(),
		PublicKey: asn1.BitString{
			Bytes: elliptic.Marshal(curve, point.X, point.Y),
		},
	})
}

func EncodeECPubKey(pubKey *btcec.PublicKey) ([]byte, error) {
	curve := btcec.S256()
	point := pubKey.ToECDSA()
	return asn1.Marshal(ECPubKey{
		Metadata: ECPubKeyMetadata{
			ECPubKeyOID:   asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1},
			NamedCurveOID: SECP256K1(),
		},
		PublicKey: asn1.BitString{
			Bytes: elliptic.Marshal(curve, point.X, point.Y),
		},
	})
}

func SECP256K1() asn1.ObjectIdentifier {
	return asn1.ObjectIdentifier{1, 3, 132, 0, 10}
}

func SelfAuthenticating(der []byte) string {
	digest := sha256.Sum224(der)
	tag := []byte{2}
	data := append(digest[:], tag...)
	crc := make([]byte, 4)
	binary.BigEndian.PutUint32(crc, crc32.ChecksumIEEE(data))
	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	str := encoder.EncodeToString(append(crc, data...))
	return strings.Join(SplitN(strings.ToLower(str), 5), "-")
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
