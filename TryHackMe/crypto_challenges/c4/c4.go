package main

import (
	"flag"
	"fmt"
	"encoding/hex"
	"strings"
	"sort"
	"os"
	"bufio"
)

type MessageEval struct {
	message string
	score float64
	key rune
}


/* 
[Challenge 4] Detect single-character XOR
One of the 60-character strings in the file attached to this task has been encrypted by single-character XOR.
Your mission is to find it.

go run c4.go -w set1_challenge4
{Now that the party is jumping 2.03479 53} 

*/

func readlines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
	}
    return lines, scanner.Err()
}

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
	w1Ptr := flag.String("w", "0", "list1")
	flag.Parse()

	words, _ := readlines(*w1Ptr)
	
	output := make([][]MessageEval, 0)

	for _, word := range words {
		bytes1 := []byte(word)
		hex1 := make([]byte, hex.DecodedLen(len(bytes1)))
		hex.Decode(hex1, []byte(word))

		msgs := TryAllCharacters(hex1)
		sort.Slice(msgs[:], func(i, j int) bool {
			return msgs[i].score > msgs[j].score
		})
		output = append(output, msgs)
	}
	

	for _, v := range output {
		for _, i := range v {
			if i.score > 2 {
				fmt.Printf("%v AND String character is %s \n\n", i, string([]rune{i.key}) )
			}
		}
	}
	return
}

