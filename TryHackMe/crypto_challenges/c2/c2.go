package main

import (
	"flag"
	"fmt"
	"encoding/hex"
)

func main() {
	w1Ptr := flag.String("w1", "0", "word 1")
	w2Ptr := flag.String("w2", "1", "word 2")

	flag.Parse()

	if len(*w1Ptr) != len(*w2Ptr) {
		fmt.Println("Length are different")
		return 
	}

	output := make([]byte, 0)

	hex1, _ := hex.DecodeString(*w1Ptr)
	fmt.Printf("%+v \n\n", hex1)
	hex2, _ := hex.DecodeString(*w2Ptr)
	fmt.Printf("%+v \n\n", hex2)

	for i:=0; i < len(hex1); i++ {
		output = append(output, hex1[i] ^ hex2[i])
	}

	fmt.Println(string(output))
	return
}

