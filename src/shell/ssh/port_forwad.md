# SSH Port Forward
Parties:
1. LOCAL machine that executes SSH commands
1. SSHD machine that runs SSH server
1. REMOTE target machine to access to

Three modes:
1. Port forward from LOCAL to REMOTE via SSHD (REMOTE may be the SSHD)
1. Reverse forward from SSHD to REMOTE via LOCAL (REMOTE may be the LOCAL)
1. Dynamic SOCKS5 server on SSHD

## LOCAL to REMOTE
`ssh -L local_bind_address:local_port:remote_host:remote_port` 

[ssh-tunnel]: https://johnliu55.tw/ssh-tunnel.html