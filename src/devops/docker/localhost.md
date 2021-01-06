# Connect from within a Container to Docker Host

>Use `host` network  
`docker run --network=host ...` the container will share the network stack with host

```sh
[host:~] $ ip a show eth0
...

[host:~] $ docker run -it --rm --network=host alpine ip a show eth0
...

# they should look the same
```

>Windows or MAC docker-desktop  
`host.docker.internal` can be used to specify the host's ip address

>Linux  
1. `sudo ip a show docker0` IP addr of the default bridge network
2. `docker run -it --rm --add-host=local:<YOUR_IP> --add-host=local_host:<YOUR_IP> image prog` to set `local` and `local_host` into container
