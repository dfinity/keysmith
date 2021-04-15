package principal

import (
	"bytes"
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/dfinity/keysmith/codec"
)

type PrincipalId interface {
	Bytes() []byte
	String() string
}

type ManagementId struct{}

func (principalId *ManagementId) Bytes() []byte {
	return nil
}

func (principalId *ManagementId) String() string {
	return show(nil)
}

type OpaqueId struct {
	data []byte
}

func (principalId *OpaqueId) Bytes() []byte {
	return principalId.data
}

func (principalId *OpaqueId) String() string {
	return show(principalId.data)
}

type SelfAuthenticatingId struct {
	data []byte
}

func NewSelfAuthenticatingId(der []byte) PrincipalId {
	hash := sha256.Sum224(der)
	data := append(hash[:], []byte{2}...)
	return &SelfAuthenticatingId{data: data}
}

func FromECPubKey(pubKey *btcec.PublicKey) (PrincipalId, error) {
	der, err := codec.EncodeECPubKey(pubKey)
	if err != nil {
		return nil, err
	}
	return NewSelfAuthenticatingId(der), nil
}

func (principalId *SelfAuthenticatingId) Bytes() []byte {
	return principalId.data
}

func (principalId *SelfAuthenticatingId) String() string {
	return show(principalId.data)
}

type DerivedId struct {
	data []byte
}

func (principalId *DerivedId) Bytes() []byte {
	return principalId.data
}

func (principalId *DerivedId) String() string {
	return show(principalId.data)
}

type AnonymousId struct{}

func (principalId *AnonymousId) Bytes() []byte {
	return []byte{4}
}

func (principalId *AnonymousId) String() string {
	return show([]byte{4})
}

type UnassignedId struct {
	data []byte
}

func (principalId *UnassignedId) Bytes() []byte {
	return principalId.data
}

func (principalId *UnassignedId) String() string {
	return show(principalId.data)
}

func FromString(str string) (PrincipalId, error) {

	// Decode.
	decoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	str32 := strings.ToUpper(strings.Replace(str, "-", "", -1))
	envelope, err := decoder.DecodeString(str32)
	if err != nil {
		return nil, err
	}

	// Check length.
	if len(envelope) < 4 {
		msg := "Invalid principal identifier: Invalid length: %s"
		return nil, fmt.Errorf(msg, str)
	}

	// Check checksum.
	reader := bytes.NewReader(envelope[0:4])
	var expect uint32
	binary.Read(reader, binary.BigEndian, &expect)
	data := envelope[4:]
	actual := crc32.ChecksumIEEE(data)
	if expect != actual {
		msg := "Invalid principal identifier: Invalid checksum: %s"
		return nil, fmt.Errorf(msg, str)
	}

	// Match type.
	var principalId PrincipalId
	if len(data) == 0 {
		principalId = &ManagementId{}
	} else {
		switch data[len(data)-1] {
		case 1:
			principalId = &OpaqueId{data: data}
		case 2:
			principalId = &SelfAuthenticatingId{data: data}
		case 3:
			principalId = &DerivedId{data: data}
		case 4:
			principalId = &AnonymousId{}
		default:
			principalId = &UnassignedId{data: data}
		}
	}

	// Check textual format.
	if str != principalId.String() {
		msg := "Invalid principal identifier: Abnormal textual format: %s"
		return nil, fmt.Errorf(msg, str)
	}

	// Return.
	return principalId, nil
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
