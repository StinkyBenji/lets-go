package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

func decodeHexBytes( input []byte ) []byte {
	// hexDecoded, err := hex.DecodeString(s)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(hexDecoded)
	dst := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(dst, input)
	if err != nil {
		log.Fatal(err)
	}
	return dst
}
 
func encodeHexBytes ( input []byte ) []byte {
	res := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(res, input)
	return res
}


func fixedXOR( input1, input2  []byte ) ([]byte, error) {
	if len(input1) != len(input2) {
		return nil, errors.New("the length of the inputs must be identical")
	}

	result := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		result[i] = input1[i] ^ input2[i]
	}
	return result, nil
}


func main () {
	in1 := decodeHexBytes([]byte("1c0111001f010100061a024b53535009181c"))
	in2 := decodeHexBytes([]byte("686974207468652062756c6c277320657965"))

	xored, _ := fixedXOR(in1, in2)

	fmt.Printf("%s", encodeHexBytes(xored))
}