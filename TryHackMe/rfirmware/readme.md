# Matheus Jesus 18-04-2020

## Try Hack Me - [Dumping Router Firmware](https://tryhackme.com/room/rfirmware)

### Setup

Download linksys.img [here](https://www.linksys.com/us/support-article?articleNum=165487)

```sh
sudo pip install cstruct
git clone https://github.com/sviehb/jefferson
cd jefferson && sudo python setup.py install
```

### Exploring

``` sh
strings linksys.img
binwalk -e linksys.img
strings 6870
```

Got Kernel version and a couple of system files

### Dumping

mount the firmware with the mount_firmware script

cd /mnt/jffs2_file

Explore!

#### Important Folders

* etc
  * SSH Configs

* JNAP
  * Used to be a security flaw. contains lua scripts and detailed lan, wan
