package main

import (
	"flag"
	"fmt"
	"encoding/hex"
	b64 "encoding/base64"
)

type MessageEval struct {
	message string
	score float64
	key rune
}

/* 

Encrypt the following with the key "ICE", using repeating-key XOR.

Burning 'em, if you ain't quick and nimble

In repeating-key XOR, you'll sequentially apply each byte of the key; the first byte of plaintext will be XOR'd
against I, the next C, the next E, then I again for the 4th byte, and so on.
	
go run c5.go -w="Burning 'em, if you ain't quick and nimble" -k=ICE
0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20
*/
func main() {
	w1Ptr := flag.String("w", "0", "word 1")
	keyPtr := flag.String("k", "ICE", "key")
	flag.Parse()
	
	bytes1 := []byte(*w1Ptr)
	bytes2 := []byte(*keyPtr)

	output := make([]byte, 0)
	current := 0
	for _, v := range bytes1 {
		xor := v ^ bytes2[current]
		current++
		if current > len(bytes2)-1 {
			current = 0
		}
		output = append(output, xor)
	}

	base64_encode := b64.StdEncoding.EncodeToString(output)
	base64_decode, err := b64.StdEncoding.DecodeString(base64_encode)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(base64_decode))

	// 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20

	return
}

