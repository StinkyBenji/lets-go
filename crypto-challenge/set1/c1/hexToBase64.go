package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func hexToBase64( s string ) string {
	hexDecoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The hexdecimal decoded string is %s\n", hexDecoded)

	base64Encoded := base64.StdEncoding.EncodeToString([]byte(hexDecoded))
	fmt.Printf("the base64 encoded string is %s\n", base64Encoded)
	
	return base64Encoded
}

func main() {
	var str string
	hexToBase64(str)
}
