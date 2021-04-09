package codec

import (
	"encoding/asn1"
)

func SECP256K1() asn1.ObjectIdentifier {
	return asn1.ObjectIdentifier{1, 3, 132, 0, 10}
}
