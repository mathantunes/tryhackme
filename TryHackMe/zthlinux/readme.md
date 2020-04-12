# Matheus Jesus 01-04-2020

## Try Hack Me - ZTHLinux (https://tryhackme.com/room/zthlinux)

### IP 10.10.127.152

#### First access

Access user shiba1:shiba1 @ SSH
Run shiba1 executable
Get password for shiba2 -> shiba2:pinguftw

```text
shiba2:pinguftw
```

#### Shiba2 User

Change user to shiba2

su shiba2
password: pinguftw

##### Basic Linux Commands

* \> used for writing output to file
  * echo hello > test.txt

* \>> used for appending output to file
  * echo world >> test.txt

* && Combine commands
  * touch file.txt && cat file.txt

* & Run on background
  * sleep 5 &

* $ Environment variables
  * $HOME

* | pipe output from one command into another
  * cat test | grep root

* ; Combine commands but failed commands won't break the chain
  * dsadsadsad ; ls

To get shiba3 password

export test1234=$USER
./shiba2

shiba3:happynootnoises

#### Shiba3 User

```text
shiba3:happynootnoises
```

CHMOD
1 Execute
2 Write
4 Read
ls -la
USER -- GROUP -- EVERYONE

CHOWN
Change Owner of file
Only when on higher user

RM
Remove files

MV
Move Files
mv file /tmp

CP
Copy Files
cp file /tmp

LN
Hard linking
Copy file and link it to change according to original
ln file /tmp

FIND
find files
find directory
find file -user
find file -group

find / -userd paradox

GREP
Filter

USER MANAGEMENT
adduser
addgroup
usermod -a -G group user

Nano
nano s.sh
\#!/bin/bash
ctrlX + Y

Important files
/etc/passwd - user info
/etc/shadow - passwords
/tmp - temporary files deleted on shutdown
/etc/sudoers - control sudo permissions
/home downloads and documents live
/root
/usr all software installations
/bin and /sbin - critical system files
/var - Linux miscellaneous directory
$PATH - Environment variable for executables

APT

GET
INSTALL


#### Shiba4

```text
shiba4:test1234
```

#### Gain Root Acess

Run
find / -user shiba2 -type f 2>>/dev/null

* Found suspicious file /var/log/test1234
* cat /var/log/test1234
* nootnoot:notsofast
* su nootnoot
* pw notsofast

* sudo -l shows that nootnoot is root
* sudo cat /root/root.txt (ad91979868d06e19d8e8a9c28be56e0c)

DONE!
