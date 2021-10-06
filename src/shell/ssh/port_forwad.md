# SSH Port Forward
Parties:
1. LOCAL machine that executes SSH commands
1. SSHD machine that runs SSH server
1. REMOTE target machine to access to

Three modes:
1. Port forward from LOCAL to REMOTE via SSHD (REMOTE may be the SSHD)
1. Reverse forward from SSHD to REMOTE via LOCAL (REMOTE may be the LOCAL)
1. Dynamic SOCKS5 server on SSHD

## `ssh` Options
* -L : forward to remote 
* -R : remote to local
* -D : dynamic 
* -n : no stdout
* -N : no scripting
* -f : background


## LOCAL to REMOTE
`ssh -L local_bind_address:local_port:remote:remote_port <SSH_SERVER>` 

Relay network traffic from **`LOCAL_BIND:PORT`** to **`SSHD`** and forward from **`SSHD`** to **`REMOTE:PORT`**.
![](https://johnliu55.tw/ssh-tunnel/images/local_scenario1_solved.png)

---

## REMOTE to LOCAL
`ssh -R remote_bind_address:remote_port:local_host:local_port <SSH_SERVER>`  

![](https://johnliu55.tw/ssh-tunnel/images/remote_scenario2_solved.png)

>NOTE  
For security purpose, the remote bind address is default to `localhost` and in order to change it configurations are to be made:  
>> --*/etc/ssh/sshd_config*--  
>> GatewayPorts yes  
* no : default to `localhost`
* yes : wildcard `0.0.0.0`
* clientspecified: specified by `ssh` `remote_bind_address`

Open a REMOTE port on SSH server, bound to remote_bind_address and relay the coming network traffic to `local_host:local_port`.

---

## DYNAMIC
`ssh -D local_bind_address:local_port <SSH_SERVER>`  

Open and listen to a local port `local_bind_address:local_port` and relay traffic over to SSH server (temporarily started `SOCKS5` server) acting as a `SOCKS5` server.

![](https://johnliu55.tw/ssh-tunnel/images/dynamic.png)

```sh
# open dynamic socks5 server on SSH server and relay from local 1080 port
ssh -D 1080 user@sshserver

# use SSH tunneled SOCKS5 server
# -k insecure
curl -x socks5://localhost:1080 \
  -k \
  https://intranet_server:3133

# -n no DNS lookup
nc -proxy-type socks5 \
  -proxy 127.0.0.1:1080 \
  -n \
  -v
  intranet_server 3133

google-chrome --user-data-dir=~/proxied-chrome \
  --proxy-server=socks5://localhost:1080

```

[ssh-tunnel]: https://johnliu55.tw/ssh-tunnel.html
[ssh-tunnel-dynamic]: https://securityintelligence.com/posts/socks-proxy-primer-what-is-socks5-and-why-should-you-use-it/
[tinyproxy]: http://tinyproxy.github.io/