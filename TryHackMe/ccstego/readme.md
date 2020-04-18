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

#### Server on 80

#### Exam1

Download png file

```sh
strings on exam1.png
#password=admin
steghide extract -p admin -sf exam1.jpeg -v -xf exam1.txt
#the key is: superkeykey
```

#### Exam2

Download the .wav file
Run on Online Spectrum analyzer -> found a [URL](https://imgur.com/KTrtNI5)
Download the image

```sh
zsteg image.png
#rKey: fatality"
```

#### Exam3

QRCode
Scanned and nothing found

```sh
stegveritas qrcode.png
```

On the results folder, it generated images with correct coloring of QRCode
scan it -> key=killshot
