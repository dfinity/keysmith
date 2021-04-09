package codec

import (
	"encoding/asn1"
)

type ECPrivKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

type ECPubKey struct {
	Metadata  []asn1.ObjectIdentifier
	PublicKey asn1.BitString
}
