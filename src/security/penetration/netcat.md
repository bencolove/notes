# The `netcat` or `nc`
The basic form (-t | __tcp__ by default)  
`nc [options] host port|port-range`
or using __udp__  
`nc -u host por3`

---

## Sending or Discovering

1. Port Scanning  
>`$ nc -z -v domain.com 1-1000`
* -z zero data, scan mode
* -v verbose, for filtering
* -n _DO NOT_ __DNS__ to resolve IP address, use __IP__ instead  

>`$ nc -z -v -n 192.168.99.99 1-10000 2>&1 | grep succeeded`  
* redirect **STDERR** to **STDOUT** to filter **succeeded** from **failed**  

[tutor-1]: https://www.digitalocean.com/community/tutorials/how-to-use-netcat-to-establish-and-test-tcp-and-udp-connections-on-a-vps