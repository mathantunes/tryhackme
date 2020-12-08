# Matheus Jesus 24-05-2020

## Try Hack Me - [Peak Hill](https://tryhackme.com/room/peakhill)

### IP 10.10.70.51

### Enum

```sh
nmap -sC -sV -oN nmap 10.10.70.51
```

```text
PORT   STATE  SERVICE  VERSION
20/tcp closed ftp-data
21/tcp open   ftp      vsftpd 3.0.3
| ftp-anon: Anonymous FTP login allowed (FTP code 230)
|_-rw-r--r--    1 ftp      ftp            17 May 15 18:37 test.txt
| ftp-syst:
|   STAT:
| FTP server status:
|      Connected to ::ffff:10.9.5.217
|      Logged in as ftp
|      TYPE: ASCII
|      No session bandwidth limit
|      Session timeout in seconds is 300
|      Control connection is plain text
|      Data connections will be plain text
|      At session startup, client count was 4
|      vsFTPd 3.0.3 - secure, fast, stable
|_End of status
22/tcp open   ssh      OpenSSH 7.2p2 Ubuntu 4ubuntu2.8 (Ubuntu Linux; protocol 2.0)
| ssh-hostkey: 
|   2048 04:d5:75:9d:c1:40:51:37:73:4c:42:30:38:b8:d6:df (RSA)
|   256 7f:95:1a:d7:59:2f:19:06:ea:c1:55:ec:58:35:0c:05 (ECDSA)
|_  256 a5:15:36:92:1c:aa:59:9b:8a:d8:ea:13:c9:c0:ff:b6 (ED25519)
Service Info: OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel
```

### FTP

```sh
ls -la
# test.txt
# .creds
```

Creds is a binary, go to [this](http://icyberchef.com)

Download the file `download.dat`

```sh
python3 read.py > ssh_pass
# User and Password
# user: gherkin
# password: p1ckl3s_@11_@r0und_th3_w0rld
```

### SSH

There is a file cmd_service.pyc

use uncompyle6 to decompile cmd_service
There are credentials hardcoded

```python
from Crypto.Util.number import long_to_bytes
username = 1684630636
password = 2457564920124666544827225107428488864802762356L
user= long_to_bytes(username)
passwd= long_to_bytes(password)

print (user+”:”+passwd)
# dill
# n3v3r_@_d1ll_m0m3nt
```

```sh
# on remote
nc 127.0.0.1 7321
# use credentials for dill
cat /home/dill/user.txt
# f1e13335c47306e193212c98fc07b6a0
```
