package main

import (
	"testing"
)

func TestDecodeGammaCypher(t *testing.T) {
	encodedData := Text("encoded data")
	salt := []uint{1, 2, 3, 4, 5}

	expectedResult := Text("decoded data")

	result := DecodeGammaCypher(encodedData, salt)

	if result != expectedResult {
		t.Errorf("DecodeGammaCypher failed, expected: %s, got: %s", expectedResult, result)
	}
}
