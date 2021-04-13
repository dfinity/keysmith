package codec

import (
	"encoding/asn1"
)

func Secp256k1() asn1.ObjectIdentifier {
	return asn1.ObjectIdentifier{1, 3, 132, 0, 10}
}

func IsSecp256k1(actual asn1.ObjectIdentifier) bool {
	expect := Secp256k1()
	if len(expect) != len(actual) {
		return false
	}
	for i := range actual {
		if expect[i] != actual[i] {
			return false
		}
	}
	return true
}
