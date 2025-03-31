package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

const input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func main() {
	//convert hexadecimal string to a slice of bytes
	bytes, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
		return
	}

	// encode this slice of bytes to base64
	// using EncodeToString inside base64 encoding package

	fmt.Println(base64.StdEncoding.EncodeToString(bytes))

}
