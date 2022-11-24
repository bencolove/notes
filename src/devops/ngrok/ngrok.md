# `ngrok`
1. install from [[zip][install]]
1. install from snap

# Install from ZIP
```shell
$ wget -s -o https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-arm.zip
# or
$ curl -s -O url
$ cd /path/to/install
$ unzip /path/to/ngrok.zip

# open web dashboard for authtoken
# connect to account
$./ngrok authtoken xxxxxxxxxtokenxxxx
# config file installed to ~/.ngrok2/ngrok.yml

# test run
$./ngrok help

# quick run
$./ngrok http 8080

# run with config file
$./ngrok start -config=/home/pi/.ngrok2/ngrok.yml <tunnel1> <tunnel2>

# inspect tunnels
$ curl http://localhost:4040/api/tunnels
```

# Install from `snap`
```shell
$ sudo apt update
$ sudo apt install snapd
$ sudo snap install ngrok
```

# Use Config File

>1. Edit config file  

--_`~/.ngrok2/ngrok.yml`_--
```txt
authtoken xxxxtokenxxx
tunnels:
  ssh:
    proto: tcp
    addr: 22
  redis:
    proto: tcp
    add: 6379

```

>2. Start ngrok service  
`$ngrok start -config=/home/pi/.ngrok2/ngrok.yml ssh redis > /dev/null &`
--log=stdout  

>3. Inspect tunnels  
`$curl http://localhost:4040/api/tunnels`  

>4. Test connection 
For tunnel `tcp://8.tcp.ngrok.io:18658`   
`$ ssh pi@8.tcp.ngrok.io -p 18658`  


[install]: https://dashboard.ngrok.com/get-started/setup
[multiple]: https://ngrok.com/docs#multiple-tunnels