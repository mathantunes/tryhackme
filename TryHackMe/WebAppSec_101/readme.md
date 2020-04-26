# Matheus Jesus 26-04-2020

## Try Hack Me - [WebAppSec 101](https://tryhackme.com/room/webappsec101)

## IP 10.10.248.127

### Recon

Browse on firefox. A web app will be displayed

### The Application

Go to a random route, get the 404 not found error.
It displays the **Apache/2.4.7 (Ubuntu) Server at 10.10.248.127 Port 80**

Try a nonsense username and password.
Checkout the response headers on firefox DevTools

```json
"X-Powered-By" : "PHP/5.5.9-1ubuntu4.24"
```

### Establishing Methodology

We could try to exploit each pages' functionalities or try to find specific topics that are subject to vulnerabilities

Topics to be tested for vulnerabilities:

#### Authorization

Either by brute-forcing weak passwords or session management.

Register an admin:admin user. It works and it has access to the /admin area

Get Cookies from DevTools Storage.
There are two cookies:

* PHPSESSID on path /
* session on path /admin

By accessing the route *http://10.10.248.127/users/sample.php?userid=1*

We can see the photos of userid=1.

Increment userid to see other users.

userid=11 is bryce, try log in as bryce:bryce

### Cross Site Scripting

[Ref](https://owasp.org/www-community/xss-filter-evasion-cheatsheet) for possible JS tag injection

Test out the possible injection fields with:

```html
<script>alert("hello")</script>
```

* The search engine is vulnerable
* The guestbook is vulnerable

### Injection

Under the Check Password Strength, when creating a new user,

#### Command Injection

Try running commands. The server does not check the input and runs a command server side

```sh
grep ^ls$ /etc/dictionaries-common/words
```

#### SQLi Injection

```sql
' or 1=1--
```

### Miscellaneous & Logic Flaws

#### Parameter Manipulation

On sample, change userid

#### Directory Traversal

Upload an image with directory ../etc/passwd

error message displays:

```text
Warning: mkdir(): Permission denied in /app/pictures/upload.php on line 27
Warning: move_uploaded_file(../upload/../etc/passwd/dsa): failed to open stream: No such file or directory in /app/pictures/upload.php on line 44
Warning: move_uploaded_file(): Unable to move '/tmp/php8nn8k1' to '../upload/../etc/passwd/dsa' in /app/pictures/upload.php on line 44
```

Access the path /upload to see all uploads.

Upload a [reverse shell](./rev_shell.php).

nc -lvnp 1234.

Gain acess to terminal

#### Logic Flaw

Apply discount coupon many times at a purchase.
