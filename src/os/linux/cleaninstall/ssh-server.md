# Install SSH Server _openssh-server_
```bash

ssh -V
# >> OpenSSH_8.2p1 Ubuntu-4ubuntu0.1, OpenSSL 1.1.1f  31 Mar 2020
apt update
#
apt install openssh-server
#
systemctl status sshd
# active
systemctl list-unit-files | grep enabled | grep ssh
# >> sshd.service   enabled    enabled
ufw status
# inactive
```