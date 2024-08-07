package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

func decodeHexStr ( s string ) []byte {
	hexDecoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return hexDecoded
}

func fixedXOR2( s1, s2  string) ([]byte, error) {
	decode1 := decodeHexStr(s1)
	decode2 := decodeHexStr(s2)

	if len(decode1) != len(decode2) {
		return nil, errors.New("the length of the inputs must be identical")
	}

	result := make([]byte, len(decode1))
	for i := 0; i < len(decode1); i++ {
		result[i] = decode1[i] ^ decode2[i]
	}

	return result, nil
}

func main () {
	xored, _ := fixedXOR2("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	
	result := hex.EncodeToString(xored)

	fmt.Printf("%s", result)
}