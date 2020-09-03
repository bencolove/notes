# Service Account VS User Account
Essential differences between those:

Diff | User Account | Service Account
---|---|---
UID range | <1000 | no limit
Login shell | no limit | /bin/false or /usr/sbin/nologin
Password | modifiable | not available

They look effectively differently in `/etc/passwd`:

`showroom:x:112:119::/home/showroom:/usr/sbin/nologin`

against normal user account:
`ubuntu:x:1000:1000:Ubuntu:/home/ubuntu:/bin/bash`

refer to 
[passwd format](../users/passwd.mg)

## Create an Service Account
`useradd` is nativate binary while
`adduser` is more user friendly scripts

```bash
$ adduser -h
# add normal user
adduser [--home DIR] [--shell SHELL] [--no-create-home] [--uid ID]
[--firstuid ID] [--lastuid ID] [--gecos GECOS] [--ingroup GROUP | --gid ID]
[--disabled-password] [--disabled-login] [--add_extra_groups]
[--encrypt-home] USER
   Add a normal user
# add service account
adduser --system [--home DIR] [--shell SHELL] [--no-create-home] [--uid ID]
[--gecos GECOS] [--group | --ingroup GROUP | --gid ID] [--disabled-password]
[--disabled-login] [--add_extra_groups] USER
   Add a system user
# add user group
adduser --group [--gid ID] GROUP
addgroup [--gid ID] GROUP
   Add a user group
# add service group
addgroup --system [--gid ID] GROUP
   Add a system group
# add existing user to existing group
adduser USER GROUP
   Add an existing user to an existing group
``` 

## Test `binary` and `script` Executables
Binary file starting with magic number `23 21` (`#!`) is script while `7F 45 4C 46` is binary executable.
```bash
$ od -N2 -tx1 -c `which adduser`
0000000  23  21
          #   !
0000002
$ od -N4 -tx1 -c `which useradd`
0000000  7f  45  4c  46
        177   E   L   F
0000004

```