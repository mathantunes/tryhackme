package main

import (
	"flag"
	"fmt"
	"encoding/hex"
	"strings"
	"sort"
)

type MessageEval struct {
	message string
	score float64
	key rune
}

/*
[Challenge 3] Single-byte XOR cipher

The hex encoded string:

    1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
.. has been XOR'd against a single character. Find the key, decrypt the message.
*/


/* 
	go run c3.go -w1 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
	Cooking MC's like a pound of bacon
*/

func TryAllCharacters(hex1 []byte) []MessageEval {
	characters := map[rune]float64 {
		'a': .08167, 'b': .01492, 'c': .02782, 'd': .04253,
        'e': .12702, 'f': .02228, 'g': .02015, 'h': .06094,
        'i': .06094, 'j': .00153, 'k': .00772, 'l': .04025,
        'm': .02406, 'n': .06749, 'o': .07507, 'p': .01929,
        'q': .00095, 'r': .05987, 's': .06327, 't': .09056,
        'u': .02758, 'v': .00978, 'w': .02360, 'x': .00150,
        'y': .01974, 'z': .00074, ' ': .13000,
	}
	msgs := make([]MessageEval,0)

	for i:=0; i < 256; i++ {

		output := make([]byte, 0)

		hex2 := byte(i)
		for i:=0; i < len(hex1); i++ {
			output = append(output, hex1[i] ^ hex2)
		}

		var score float64
		low := strings.ToLower(string(output))
		for _, r := range []rune(low) {
			score += characters[r]
		}
		msgs = append(msgs, *&MessageEval{
			score : score,
			message : string(output),
			key : rune(i),
		})
	}

	return msgs
}

func main() {
	w1Ptr := flag.String("w1", "0", "word 1")
	flag.Parse()

	bytes1 := []byte(*w1Ptr)
	hex1 := make([]byte, hex.DecodedLen(len(bytes1)))
	hex.Decode(hex1, []byte(*w1Ptr))

	msgs := TryAllCharacters(hex1)
	
	sort.Slice(msgs[:], func(i, j int) bool {
		return msgs[i].score < msgs[j].score
	})

	for _, v := range msgs {
		fmt.Printf("%+v  AND String character is %s \n\n", v, string([]rune{v.key}) )
	}
	return
}

