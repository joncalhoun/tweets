// A slice of bytes doesn't always convert directly into a valid UTF-8 string, meaning string(byteSlice) won't always render a usable result. One way to avoid this is to encode & decode your byte slices using the hex or base64 packages in #golang.

package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// Let's say we have bytes from hashing a value or something similar...
	origBytes := pretendToHash("some value")

	// One way to encode these into a string is base64
	str := base64.URLEncoding.EncodeToString(origBytes)
	fmt.Println(str)

	// We can decode using the same encoding.
	decodedBytes, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}

	// We can verify this with a test, but for this demo I'll use panic
	if len(origBytes) != len(decodedBytes) {
		panic("our base64 decoded bytes aren the same length as the original bytes!")
	}
	for i, b := range origBytes {
		if b != decodedBytes[i] {
			panic(fmt.Sprintf("base64 bytes differ at index %d", i))
		}
	}

	// We can do something similar with hex
	strHex := hex.EncodeToString(origBytes)
	fmt.Println(strHex)
	decodedBytes, err = hex.DecodeString(strHex)
	if err != nil {
		panic(err)
	}
	// We can verify this with a test, but for this demo I'll use panic
	if len(origBytes) != len(decodedBytes) {
		panic("our hex decoded bytes aren the same length as the original bytes!")
	}
	for i, b := range origBytes {
		if b != decodedBytes[i] {
			panic(fmt.Sprintf("hex bytes differ at index %d", i))
		}
	}
}

func pretendToHash(thing string) []byte {
	return []byte{
		0b10101100, 0b10110100, 0b01010001, 0b01011111, 0b10001000, 0b00111101, 0b01110110, 0b01010001, 0b11110011, 0b00000110, 0b00100111, 0b11101101, 0b00010100, 0b10110011, 0b00010000, 0b11100101, 0b11110101, 0b11011011, 0b11011101, 0b10111101, 0b01100000, 0b11000111, 0b00001100, 0b01000111, 0b00010101, 0b00110011, 0b11000001, 0b11001110, 0b11010100, 0b10010001, 0b01011101, 0b00110001, 0b11110110, 0b10110010, 0b01001110, 0b00101000, 0b00010111, 0b11001001, 0b10011010, 0b01100000, 0b10011101, 0b10001010, 0b00011001, 0b00011001, 0b00100000, 0b01100111, 0b10110011, 0b00010011, 0b10110001, 0b00010101, 0b00011110, 0b01110011, 0b01010000, 0b10000001, 0b00011011, 0b11100000, 0b01111110, 0b01011110, 0b01000000, 0b10111100, 0b01100101, 0b11111111, 0b00001101, 0b00100010, 0b00000010, 0b11011000, 0b00100101, 0b10011001, 0b00000110, 0b11010101, 0b01111000, 0b11011000, 0b11110011, 0b00001100, 0b01100011, 0b01111000, 0b01010001, 0b10001000, 0b01110100, 0b01101011, 0b00011011, 0b11001111, 0b00011101, 0b01001101, 0b00011110, 0b00010110, 0b01100110, 0b11010111, 0b01110001, 0b00011100, 0b00001101, 0b01101110, 0b01000001, 0b11000100, 0b10000010, 0b11001110, 0b00000010, 0b01111011, 0b10011010, 0b10010110, 0b00001001, 0b01110001, 0b00110100, 0b10101010, 0b00110011, 0b00110001, 0b00110000, 0b11111001, 0b10101100, 0b11101111, 0b00010110, 0b00011110, 0b01011001, 0b10100101, 0b01010001, 0b01010010, 0b11100110, 0b01111101, 0b01000011, 0b00111010, 0b11010101, 0b00011010, 0b01000101, 0b10100110, 0b11111010, 0b01010010, 0b00101010, 0b11000010,
	}
}
