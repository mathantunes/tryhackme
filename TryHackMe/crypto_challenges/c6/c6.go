package main

import (
	"flag"
	"fmt"
	b64 "encoding/base64"
	"strings"
	"sort"
	"os"
	"bufio"
)

import "github.com/steakknife/hamming"

type MessageEval struct {
	message string
	score float64
	key rune
}

func readlines(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var line string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line += scanner.Text()
	}
    return line, scanner.Err()
}

func TryAllCharacters(hex1 []byte) []MessageEval {
	characters := map[rune]float64 {
		'a': .08167, 'b': .01492, 'c': .02782, 'd': .04253,
        'e': .12702, 'f': .02228, 'g': .02015, 'h': .06094,
        'i': .06094, 'j': .00153, 'k': .00772, 'l': .04025,
        'm': .02406, 'n': .06749, 'o': .07507, 'p': .01929,
        'q': .00095, 'r': .05987, 's': .06327, 't': .09056,
        'u': .02758, 'v': .00978, 'w': .02360, 'x': .00150,
        'y': .01974, 'z': .00074,
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

func ComputeHamming(h1, h2 []byte) int {
	var out int
	for i:=0; i< len(h1); i++ {
		out += hamming.Byte(h1[i], h2[i])
	}
	return out
}

type Hammed struct {
	keysize int
	hamming float64
	mean float64
}

func TestComputeHamming() {
	value := ComputeHamming([]byte("this is a test"), []byte("wokka wokka!!!"))
	//fmt.Println(value) //Should be 37 - OK
	if value != 37 {
		// fmt.Errorf("Hammin is not 37 - Fail %v", value)
	}
}

func ProcessHamm(word string, keysize int) (*Hammed, bool) {
	distances := make([]float64, 0)
	for i:=1; i < (len(word)-keysize*2+2); i+= keysize {
		
		first := word[i-1:i-1+keysize] //0:1
		second := word[i-1+keysize:i-1+keysize*2]//2:3
		distance := ComputeHamming([]byte(first), []byte(second))

		floatDistance := float64((float64(distance))/float64(keysize))

		distances = append(distances, floatDistance)
	}

	sort.Slice(distances[:], func(i,j int)bool {
		return distances[i] < distances[j]
	})
	sum := 0.0
	for _, d := range distances {
		sum+=d
	}
	if len(distances) > 0 {
		ham := &Hammed{
			keysize: keysize,
			hamming: distances[0],
			mean: sum/float64(len(distances)),
		}
		return ham, true
	}
	return nil, false
}

func MountBlocks(word string, h Hammed)[]string {
	blocks := make([]string, 0)

	// [HUIfTQsPAh9PE048Gmll H0kcDk4TAQsHThsBFkU2]
	for i:=0; i<len(word)-h.keysize+4; i+= h.keysize {
		blocks = append(blocks, word[i:i+h.keysize])
	}
	return blocks
}

func Transpose(h Hammed, blocks []string)[][]string {
	transp := make([][]string, 0)
	for i:=0; i<h.keysize; i++ {
		intern := make([]string,0)
		for j:=0; j<len(blocks); j++ {
			intern = append(intern, blocks[j][i:(i+1)] )
		}
		transp = append(transp, intern)
	}
	return transp
}

func FindPossibleKeys(transp [][]string)[][]MessageEval {
	possibleKeys := make([][]MessageEval, 0)
	for _, bytesT := range transp {

		var bytesAsString string
		for _, byteT := range bytesT {
			bytesAsString += byteT
		}

		bytes1 := []byte(bytesAsString)
		hex1 := bytes1
		results := TryAllCharacters(hex1)

		possibleKeys = append(possibleKeys, results)
	}
	return possibleKeys
}

func MountKey(possibleKeys [][]MessageEval) string {
	var key string
	for _, res := range possibleKeys {
		sort.Slice(res[:], func(i, j int) bool {
			return res[i].score > res[j].score
		})
		key += string([]rune{res[0].key})
	}

	return key
}
/*
go run c6.go -w set1_challenge6 
Key:  tER(InAtOR x: brinG tHe NOiSe
*/
func main() {
	w1Ptr := flag.String("w", "0", "list1")
	flag.Parse()

	b64Encoded, err := readlines(*w1Ptr)
	if err != nil {
		panic(err)
	}
	dec, _ := b64.StdEncoding.DecodeString(b64Encoded)

	word := string(dec)
	ham := make([]Hammed,0)
	for keysize := 2; keysize < 40; keysize++ {
		h, done := ProcessHamm(word, keysize)
		if done {
			ham = append(ham, *h)
		}
	}

	//Sort by mean
	sort.Slice(ham[:], func(i,j int)bool {
		return ham[i].mean > ham[j].mean
	})

	//Break the ciphertext into blocks of KEYSIZE length
	h0 := ham[len(ham)-1]
	blocks := MountBlocks(word, h0)
	// Transpose the blocks
	// make a block that is the first byte of every block, and a block that is the          
	// second byte of every block, and so on.
	transp := Transpose(h0, blocks)
	// Solve Each Block with Single-Character XOR Technique
	possibleKeys := FindPossibleKeys(transp) 
	key := MountKey(possibleKeys)
	fmt.Println("Key: ", key)
	return
                                                                  
}

