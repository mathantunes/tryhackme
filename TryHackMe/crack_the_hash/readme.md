# Matheus Jesus 04-04-2020

## Try Hack Me - Crack the Hash (https://tryhackme.com/room/crackthehash)

### Level 1

#### MD5

hashid 48bb6e862e54f2a795ffc4e541caed4d

hashcat -m 0 -a 0 md5.txt /usr/share/wordlists/rockyou.txt --force --show

s48bb6e862e54f2a795ffc4e541caed4d -> https://crackstation.net/ -> easy

#### SHA1

CBFDAC6008F9CAB4083784CBD1874F76618D2A97 -> password123

hashcat -m 100 -a 0 sha1.txt /usr/share/wordlists/rockyou.txt --force

#### SHA256

1C8BFE8F801D79745C4631D09FFF36C82AA37FC4CCE4FC946683D7B336B63032 -> letmein

hashcat -m 1400 -a 0 sha256.txt /usr/share/wordlists/rockyou.txt --force

#### BCrypt

$2y$12$Dwt1BZj6pcyc3Dy1FWZ5ieeUznr71EeNkJkUlypTsgbX1H68wsRom
Dictionary cache hit:

* Filename..: /usr/share/wordlists/rockyou.txt
* Passwords.: 14344385
* Bytes.....: 139921507
* Keyspace..: 14344385

Took too long for brute forcing on rockyou.txt
Since the TryHackMe website says it's a four letter word, generate a wordlist with all possibilites of 4 letter combinations and brute force it

DONE!!!

**Filter rockyou.txt, save to 4letters.txt only the lines with length == 4**
awk 'length == 4' /usr/share/wordlists/rockyou.txt > 4letters.txt 

#### NTLM

hashcat -m 1000 -a 0 ntlm.txt /usr/share/wordlists/rockyou.txt --force --show > ntlm_cracked.txt

#### HMAC-SHA1

use -m 160 to compute with salt as **hash:salt** on the hmac_sha1.txt file

hashcat -m 160 -a 0 hmac_sha1.txt /usr/share/wordlists/rockyou.txt --force --show > hmac_sha1_cracked.txt
