# Matheus Jesus 19-04-2020

## Try Hack Me - [Biohazard](https://tryhackme.com/room/biohazard)

### IP 10.10.65.222

### Recon

* NMAP identifies an HTTP server on 80
* SSH on 22
* FTP on 21

### Roleplaying lol

#### Alpha team

* Chris **MISSING**
* Jill **Investigating /diningRoom/**
* Barry **Investigating /diningRoom/**
* Weasker
* Joseph **DEAD**

#### Bravo team

* Kenneth **DEAD** on /teaRoom/
* Richard **DEAD** on /attic

found this SG93IGFib3V0IHRoZSAvdGVhUm9vbS8= on diningroom **NOT YET USEFUL**
It is base64 encoded

```sh
echo 'SG93IGFib3V0IHRoZSAvdGVhUm9vbS8=' | base64 -d
#How about the /teaRoom/
```

Get Emblem **emblem{fec832623ea498e20bf4fe1821d58727}**

##### On TeaRoom

lock_pick{037b35e2ff90916a9abf99129c8e1837}

##### On ArtRoom

Location:
/diningRoom/
/teaRoom/
/artRoom/
/barRoom/
/diningRoom2F/
/tigerStatusRoom/
/galleryRoom/
/studyRoom/
/armorRoom/
/attic/

##### On BarRoom

Got in with lockpick flag
Play the Piano?

Moonlight Somata -> NV2XG2LDL5ZWQZLFOR5TGNRSMQ3TEZDFMFTDMNLGGVRGIYZWGNSGCZLDMU3GCMLGGY3TMZL5

it is base32 

```sh
echo 'NV2XG2LDL5ZWQZLFOR5TGNRSMQ3TEZDFMFTDMNLGGVRGIYZWGNSGCZLDMU3GCMLGGY3TMZL5' | base32 -d
#music_sheet{362d72deaf65f5bdc63daece6a1f676e}
```

Play the piano with music_sheet

##### On barRoom357162e3db904857963e6e0b64b96ba7/barRoomHidden.php (Secret Bar Room)

Got gold emblem gold_emblem{58a8c41a9d08b8a4e38d02a4d7ff4843}

refresh and write emblem flag -> rebecca

##### Back on dininRoom

Put gold emblem flag

klfvg ks r wimgnd biz mpuiui ulg fiemok tqod. Xii jvmc tbkg ks tempgf tyi_hvgct_jljinf_kvc

Using [this website](https://www.guballa.de/vigenere-solver)

there is a shield key inside the dining room. The html page is called the_great_shield_key

shield_key{48a7a9227cd7eb89f0a062590798cbac}

##### On Gallery Room

crest 2:
GVFWK5KHK5WTGTCILE4DKY3DNN4GQQRTM5AVCTKE
Hint 1: Crest 2 has been encoded twice
Hint 2: Crest 2 contanis 18 letters
Note: You need to collect all 4 crests, combine and decode to reavel another path
The combination should be crest 1 + crest 2 + crest 3 + crest 4. Also, the combination is a type of encoded base and you need to decode it

##### On Armor Room

crest 3:
MDAxMTAxMTAgMDAxMTAwMTEgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAwMTEgMDAxMDAwMDAgMDAxMTAxMDAgMDExMDAxMDAgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAxMTAgMDAxMDAwMDAgMDAxMTAxMDAgMDAxMTEwMDEgMDAxMDAwMDAgMDAxMTAxMDAgMDAxMTEwMDAgMDAxMDAwMDAgMDAxMTAxMTAgMDExMDAwMTEgMDAxMDAwMDAgMDAxMTAxMTEgMDAxMTAxMTAgMDAxMDAwMDAgMDAxMTAxMTAgMDAxMTAxMDAgMDAxMDAwMDAgMDAxMTAxMDEgMDAxMTAxMTAgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTEwMDEgMDAxMDAwMDAgMDAxMTAxMTAgMDExMDAwMDEgMDAxMDAwMDAgMDAxMTAxMDEgMDAxMTEwMDEgMDAxMDAwMDAgMDAxMTAxMDEgMDAxMTAxMTEgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAxMDEgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAwMDAgMDAxMDAwMDAgMDAxMTAxMDEgMDAxMTEwMDAgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAwMTAgMDAxMDAwMDAgMDAxMTAxMTAgMDAxMTEwMDA=
Hint 1: Crest 3 has been encoded three times
Hint 2: Crest 3 contanis 19 letters
Note: You need to collect all 4 crests, combine and decode to reavel another path
The combination should be crest 1 + crest 2 + crest 3 + crest 4. Also, the combination is a type of encoded base and you need to decode it

##### On Attic

crest 4:
gSUERauVpvKzRpyPpuYz66JDmRTbJubaoArM6CAQsnVwte6zF9J4GGYyun3k5qM9ma4s
Hint 1: Crest 2 has been encoded twice
Hint 2: Crest 2 contanis 17 characters
Note: You need to collect all 4 crests, combine and decode to reavel another path
The combination should be crest 1 + crest 2 + crest 3 + crest 4. Also, the combination is a type of encoded base and you need to decode it

##### On dinning room 2f

Lbh trg gur oyhr trz ol chfuvat gur fgnghf gb gur ybjre sybbe. Gur trz vf ba gur qvavatEbbz svefg sybbe. Ivfvg fnccuver.ugzy

You get the blue gem by pushing the status to the lower floor. The gem is on the diningRoom first floor. Visit sapphire.html

/diningRoom/sapphire.html -> blue_jewel{e1d457e96cac640f863ec7bc475d48aa}

crest 1:
S0pXRkVVS0pKQkxIVVdTWUpFM0VTUlk9
Hint 1: Crest 1 has been encoded twice
Hint 2: Crest 1 contanis 14 letters
Note: You need to collect all 4 crests, combine and decode to reavel another path
The combination should be crest 1 + crest 2 + crest 3 + crest 4. Also, the combination is a type of encoded base and you need to decode it

#### CRESTS

1 - S0pXRkVVS0pKQkxIVVdTWUpFM0VTUlk9
2 - GVFWK5KHK5WTGTCILE4DKY3DNN4GQQRTM5AVCTKE
3 - MDAxMTAxMTAgMDAxMTAwMTEgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAwMTEgMDAxMDAwMDAgMDAxMTAxMDAgMDExMDAxMDAgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAxMTAgMDAxMDAwMDAgMDAxMTAxMDAgMDAxMTEwMDEgMDAxMDAwMDAgMDAxMTAxMDAgMDAxMTEwMDAgMDAxMDAwMDAgMDAxMTAxMTAgMDExMDAwMTEgMDAxMDAwMDAgMDAxMTAxMTEgMDAxMTAxMTAgMDAxMDAwMDAgMDAxMTAxMTAgMDAxMTAxMDAgMDAxMDAwMDAgMDAxMTAxMDEgMDAxMTAxMTAgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTEwMDEgMDAxMDAwMDAgMDAxMTAxMTAgMDExMDAwMDEgMDAxMDAwMDAgMDAxMTAxMDEgMDAxMTEwMDEgMDAxMDAwMDAgMDAxMTAxMDEgMDAxMTAxMTEgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAxMDEgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAwMDAgMDAxMDAwMDAgMDAxMTAxMDEgMDAxMTEwMDAgMDAxMDAwMDAgMDAxMTAwMTEgMDAxMTAwMTAgMDAxMDAwMDAgMDAxMTAxMTAgMDAxMTEwMDA=
4 - gSUERauVpvKzRpyPpuYz66JDmRTbJubaoArM6CAQsnVwte6zF9J4GGYyun3k5qM9ma4s

```sh
echo S0pXRkVVS0pKQkxIVVdTWUpFM0VTUlk9 | base64 -d | base32 -d
# 1 - RlRQIHVzZXI6IG
echo GVFWK5KHK5WTGTCILE4DKY3DNN4GQQRTM5AVCTKE | base32 -d #AND base 58 (online)
# 2 - h1bnRlciwgRlRQIHBh
# 3 base64 | binary | hex - c3M6IHlvdV9jYW50X2h
# 4 base58 | hex  pZGVfZm9yZXZlcg==
# RlRQIHVzZXI6IGh1bnRlciwgRlRQIHBhc3M6IHlvdV9jYW50X2hpZGVfZm9yZXZlcg==
echo RlRQIHVzZXI6IGh1bnRlciwgRlRQIHBhc3M6IHlvdV9jYW50X2hpZGVfZm9yZXZlcg== | base64 -d
# FTP user: hunter, FTP pass: you_cant_hide_forever
```

#### FTP In

download files

exiftool 002-key.jpg -> 5fYmVfZGVzdHJveV9

steghide extract -sf 001-key.jpg -> cGxhbnQ0Ml9jYW
binwalk -e 003-key.jpg
cat key-003.txt -> 3aXRoX3Zqb2x0

cGxhbnQ0Ml9jYW5fYmVfZGVzdHJveV93aXRoX3Zqb2x0

```sh
echo cGxhbnQ0Ml9jYW5fYmVfZGVzdHJveV93aXRoX3Zqb2x0 | base64 -d
# plant42_can_be_destroy_with_vjolt
gpg -d helmet_key.txt.gpg
# gpg: AES256 encrypted data
# gpg: encrypted with 1 passphrase
# helmet_key{458493193501d2b94bbab2e727f8db4b}
```

##### On Hidden Closet

wpbwbxr wpkzg pltwnhro, txrks_xfqsxrd_bvv_fy_rvmexa_ajk
SSH password: T_virus_rules

##### On Study Room

SSH user: umbrella_guest

### SSH

```sh
cd .jailcell
cat chris.txt
Jill: Chris, is that you?
# Chris: Jill, you finally come. I was locked in the Jail cell for a while. It seem that weasker is behind all this.
# Jil, What? Weasker? He is the traitor?
# Chris: Yes, Jill. Unfortunately, he play us like a damn fiddle.
# Jill: Let's get out of here first, I have contact brad for helicopter support.
# Chris: Thanks Jill, here, take this MO Disk 2 with you. It look like the key to decipher something.
# Jill: Alright, I will deal with him later.
# Chris: see ya.

# MO disk 2: albert
```

Use albert as key for wpbwbxr wpkzg pltwnhro, txrks_xfqsxrd_bvv_fy_rvmexa_ajk

decrypted - weasker login password, stars_members_are_my_guinea_pig

Weasker is ROOT

```sh
sudo /bin/bash
cd /root
cat root.txt
```

monster name is Tyrant
root flag 3c5794a00dc56c35f2bf096571edf3bf
