package codec

import (
	"crypto/elliptic"
	"encoding/asn1"
	"encoding/pem"

	"github.com/btcsuite/btcd/btcec"
)

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

func EncodeECParams() ([]byte, error) {
	return asn1.Marshal(Secp256k1())
}

func EncodeECPrivKey(privKey *btcec.PrivateKey) ([]byte, error) {
	curve := btcec.S256()
	point := privKey.PubKey().ToECDSA()
	return asn1.Marshal(ECPrivKey{
		Version:       1,
		PrivateKey:    privKey.D.Bytes(),
		NamedCurveOID: Secp256k1(),
		PublicKey: asn1.BitString{
			Bytes: elliptic.Marshal(curve, point.X, point.Y),
		},
	})
}

func EncodeECPubKey(pubKey *btcec.PublicKey) ([]byte, error) {
	curve := btcec.S256()
	point := pubKey.ToECDSA()
	return asn1.Marshal(ECPubKey{
		Metadata: []asn1.ObjectIdentifier{
			asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1},
			Secp256k1(),
		},
		PublicKey: asn1.BitString{
			Bytes: elliptic.Marshal(curve, point.X, point.Y),
		},
	})
}

func EncodeECSig(sig *btcec.Signature) []byte {
	var buf [64]byte
	r := sig.R.Bytes()
	s := sig.S.Bytes()
	copy(buf[(32-len(r)):], r)
	copy(buf[(64-len(s)):], s)
	return buf[:]
}
