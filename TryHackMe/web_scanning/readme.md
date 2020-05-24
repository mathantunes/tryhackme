# Matheus Jesus 16-05-2020

## Try Hack Me - [RP: Web Scanning](https://tryhackme.com/room/rpwebscanning)

Intro to Web Scanning.

### IP 10.10.189.175

### Nikto

```sh
#Basic usage
nikto -h 10.10.189.175
# Server: Apache/2.4.7 (Ubuntu)
```

### GoBuster

```sh
gobuster dir -u http://10.10.189.175:80 -w /usr/share/dirbuster/wordlists/directory-list-2.3-medium.txt
# /docs (Status: 301)
# /external (Status: 301)
# /config (Status: 301)
```

### Owasp-ZAP

Set URL to Attack: 10.10.189.175 -> **Attack**

It shows that the server is vulnerable to XSS,
Brute force HTTP login
HttpOnly flag not set
