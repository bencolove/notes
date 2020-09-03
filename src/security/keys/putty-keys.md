# PuTTY Keys
`.ppk` PuTTY keys' file extension contains both public and private key pair.

## Convert PuTTY Keys to OpenSSH
```bash
$ sudo apt install putty-tools
# convert with puttygen
# private key only
$ puttygen id_dsa.ppk -O private-openssh -o id_dsa
$ chmod 600 id_dsa
# pulibc key only
$ puttygen id_dsa.ppk -O public-openssh -o id_dsa.pub
```