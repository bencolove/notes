# Connect from within a Container to Docker Host

>Windows or MAC docker-desktop  
`host.docker.internal` can be used to specify the host's ip address

>Linux  
1. `sudo ip a show docker0` IP addr of the default bridge network
2. `docker run -it --rm --add-host=local:<YOUR_IP> --add-host=local_host:<YOUR_IP> image prog` to set `local` and `local_host` into container
