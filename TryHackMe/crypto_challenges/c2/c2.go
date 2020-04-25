package main

import (
	"flag"
	"fmt"
	"encoding/hex"
)

/*
[Challenge 2] Fixed XOR

Write a function that takes 2 equal length buffers and produces their XOR combination. 

If it works properly when you feed the string: 1c0111001f010100061a024b53535009181c
.. after hex decoding and XOR'd against: 686974207468652062756c6c277320657965
.. it should produce: 746865206b696420646f6e277420706c6179

go run c2.go -w1 1c0111001f010100061a024b53535009181c -w2 686974207468652062756c6c277320657965
746865206b696420646f6e277420706c6179
*/

func main() {
	w1Ptr := flag.String("w1", "0", "word 1")
	w2Ptr := flag.String("w2", "1", "word 2")

	flag.Parse()
	if len(*w1Ptr) != len(*w2Ptr) {
		fmt.Println("W1 and W2 cannot differ in length")
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

