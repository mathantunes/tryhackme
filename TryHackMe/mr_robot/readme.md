# Matheus Jesus 10-04-2020

## Try Hack Me - MR Robot (https://tryhackme.com/room/mrrobot)

### VPN IP 10.9.5.234
### IP 10.10.77.104

#### nmap

found http server on 80 and 443

#### GOBUSTER

gobuster dir -u http://10.10.77.104:80 -w /usr/share/dirbuster/wordlists/directory-list-2.3-medium.txt

found a couple of useful directories

/readme
/robots

```text
User-agent: *
fsocity.dic
key-1-of-3.txt
```

by running /fsocity.dic -> found a dictionary (maybe used for cracking)
by runnning /key-1-of-3.txt -> found the first key 073403c8a58a1f80d943455fb30724b9

#### NIKTO

run nikto -h 10.10.77.104

#### wpscan

wpscan --url 10.10.77.104 --enumerate vp

#### wp-admin

try login on wordpress

admin:admin -> invalid username
elliot:admin -> invalid password

try bruteforce with dictionary

wpscan --url 10.10.77.104 -P fsocity.dic --usernames elliot

##### Stupid dictionary has ER28-0652 on last item

elliot:ER28-0652

##### CRACK MD5

on /home/robot md5 file
robot:c3fcd3d76192e4007dfb496cca67e13b

abcdefghijklmnopqrstuvwxyz

cat key-2-
822c73956184f694993bede3eb39f959

#### PRIVILEGE ESCALATION

NOT DONE
cd /root
