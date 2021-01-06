# Config WIFI
By:
1. default networking manager `networking`
1. `wap_supplicant`
1. dhcp manager `dhcpcd`
---
## Default setup `wpa` + `dhcp`
[source][source1]  
_--`/etc/wpa_supplicant/wpa_supplicant.config`--_
```txt
...
network={
    ssid="SSID"
    psk="YOUR_PASSWORD"
    // for network without password
    key_mgmt=NONE
    // for hidden network
    scan_ssid=1
}
```
>Verify Status  
```sh
$ ifconfig wlan0
$ ip a show wlan0  

inet xx.xx.xx.xx/xx
...
# A valid inet address means connected

$ ip r | grep default | awk '{ print $3 }'
10.1.1.1
# router addr

$ ip -4 a | grep global | awk '{ print $3 }'
10.1.1.107/20
# assigned ip address

```

>Apply Changes  
1. reconfig WIFI  
`wpa_cli -i wlan0 reconfigure`
2. verify  
`ip a show wlan0`
3. unblock interface
```sh
$ rfkill list all

0: phy0: Wireless LAN
    Soft blocked: yes
    Hard blocked: no
1: hci0: Bluetooth
    Soft ...

$ rfkill unblock 0
```
the `rfkill`'s `blocked/unblocked` status corresponds to GUI's turn Wi-Fi on/off

---
## Configiure Networking

### Default debian networking manager:  
`sudo systemctl enable networking`

>Config files  

_`--/etc/network/interfaces--`_
```txt
auto eth0
iface eth0 inet static
        address 10.1.1.30
        netmask 255.255.255.0
        gateway 10.1.1.1

allow-hotplug wlan0
iface wlan0 inet static
        address 10.1.1.31
        netmask 255.255.255.0
        gateway 10.1.1.1
    wpa-conf /etc/wpa_supplicant/wpa_supplicant.conf
```

### DHCP Manager
1. leave `/etc/network/interfaces` as default
2. manager app `sudo systemctl enable dhcpcd`
2. config files  
_`--/etc/dhcpcd.conf--`_
```txt
interface eth0
static ip_address=10.1.1.30/24
static routers=10.1.1.1
static domain_name_servers=10.1.1.1

interface wlan0
static ip_address=10.1.1.31/24
static routers=10.1.1.1
static domain_name_servers=10.1.1.1
```

[source1]: https://raspberrypi.stackexchange.com/questions/37920/how-do-i-set-up-networking-wifi-static-ip-address-on-raspbian-raspberry-pi-os