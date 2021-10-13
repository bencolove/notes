# Connect from within a Container to Docker Host
1. use `host` network (same network stack with host machine)
1. modify container's networking

---

## the `host` network
`docker run --network=host ...` the container will share the network stack with host

```sh
[host:~] $ ip a show eth0
...

[host:~] $ docker run -it --rm --network=host alpine ip a show eth0
...

# they should look the same
```

---

## Modify Container's Networking
`/etc/hosts`, `/etc/resolv.conf` and `/etc/hostname` are mapped(mounted) from `host:/var/lib/docker/containers/CONTAINER_ID` 

We may manually insert a record into _`/etc/hosts`_:
```txt
192.168.65.2    host.docker.internal
```


>Windows or MAC docker-desktop  
`host.docker.internal` can be used to specify the host's ip address

>Linux  
1. `sudo ip a show docker0` IP addr of the default bridge network  
`ip r | grep default | awk '{ print $3 }'`  
`netstat -nr | grep '^0\.0\.0\.0' | awk '{print $2}'`  
2. `docker run -it --rm --add-host=local:<YOUR_IP> --add-host=local_host:<YOUR_IP> image prog` to set `local` and `local_host` into container

>`host-gateway`  
`--add-host=host.docker.internal:host-gateway` will do the trick to replace `host-gateway` with `host` IP.

### Add `host` from Dockerfile
```txt

USER root

RUN mkdir /scripts
COPY run.sh /scripts/
RUN chmod +x /scripts/run.sh
RUN /bin/sh /scripts/run.sh

```
and 
_`run.sh`_
```sh
echo "172.10.10.1 host.docker.internal" >> /etc/hosts
```