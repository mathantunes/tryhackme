# Matheus Jesus 22-04-2020

## Try Hack Me - [Crypto Challenges](https://tryhackme.com/room/cryptochallenges)

"This is a different way to learn about crypto than taking a class or reading a book. We give you problems to solve. They're derived from weaknesses in real-world systems and modern cryptographic constructions. We give you enough info to learn about the underlying crypto concepts yourself. When you're finished, you'll not only have learned a good deal about how cryptosystems are built, but you'll also understand how they're attacked."

### Challenge 1

Hex to Base64
Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

```sh
echo "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d" -n  | xxd -r -p | base64
#SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
```

### Challenge 2

Fixed XOR
Write a function that takes 2 equal length buffers and produces their XOR combination.

```sh
go run c2.go -w1=1c0111001f010100061a024b53535009181c -w2=686974207468652062756c6c277320657965
```

### Challenge 3

Single byte XOR cipher

```sh
go run c3.go -w1 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
#X Cooking MC's like a pound of bacon
```

### Challenge 4

```sh
go run c4.go -w set1_challenge4 > out.txt
# {Now that the party is jumping 2.03479 53}
```

### Challenge 5

```sh
go run c5.go -w="Burning 'em, if you ain't quick and nimble" -k=ICE
# 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20
```

### Challenge 6

Break repeating-key XOR

```sh
go run c6.go -w set1_challenge6 
# Key:  tER(InAtOR x: brinG tHe NOiSe
```

### Challenge 7

```sh
python c7.py > output.txt
```
