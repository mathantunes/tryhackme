# Matheus Jesus 17-04-2020

## Try Hack Me - [CC Stego](https://tryhackme.com/room/ccstego)

### Steghide

Hiding messages inside images.

#### JPEG1

steghide extract -p password123 -sf jpeg1.jpeg -v -xf msg_jpeg1.txt

cat msg_jpeg1.txt
found pinguftw

### Zsteg

zsteg png1.png

```sh
b1,bgr,lsb,xy       .. text: "nootnoot$"
```

### Exiftool

View and edit image metadata

exiftool jpeg3.jpeg

### Stegoveritas

Support every image file, extract all data.

stegoveritas -steghide -o out_jpeg2 jpeg2.jpeg

cat out_jpeg2/steghide_61270fd72f9a1e5cd47bc4e226a6c524.bin -> kekekekek

### Spectograms

Hide hidden an image inside an audio file`s spectogram.

For audio, they utilize **Sonic Visualize**

[Online Spectrum analyzer](https://academo.org/demos/spectrum-analyzer/)

broke the message ~"Google"

### Final Exam
