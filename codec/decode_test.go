package codec

import (
	"testing"
)

func TestLoadAlternate(test *testing.T) {
	_, err := LoadECPrivKey("test/alternate.pem")
	if err == nil {
		test.Fatal("Failed to reject alternate identity file.")
	}
}

func TestLoadCorrupt(test *testing.T) {
	_, err := LoadECPrivKey("test/corrupt.pem")
	if err == nil {
		test.Fatal("Failed to reject corrupt identity file.")
	}
}

func TestLoadEmpty(test *testing.T) {
	_, err := LoadECPrivKey("test/empty.pem")
	if err == nil {
		test.Fatal("Failed to reject empty identity file.")
	}
}

func TestLoadValid(test *testing.T) {
	_, err := LoadECPrivKey("test/valid.pem")
	if err != nil {
		test.Fatal("Failed to accept valid identity file.")
	}
}
