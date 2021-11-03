# `lsof`
The `lsof`(_list open files_) for:
1. regular files
1. directories
1. devices
1. libraries
1. pipes
1. symbollinks
1. network
    * socket
    * unix domain name
    * NFS

## filter
Preceding `^` to _`nagate(exclude)`_

FLAG | FILTER
---|---
-c COMM  | process name(default to COMMAND), or `/^..o.$/i`
-p PID | process PID($$)
-u UID | User name($USER) or UID($UID)
-g GID | process group PGID
-a | above all
-d fds | csv 'cwd, 1, ^2'
+d folder | folder
-i | network filter: <br> `-i TCP:3306` port 3306 <br> `-i @HOST:PROTO:PORT1,PROTO:PORT2`  
-n | no DNS lookup, service name lookup
-s PROTO:STATE | `-s TCP:LISTEN` listening TCP


## select columns
1. COMMAND
1. PID
1. PPID
1. USER
1. PGID
1. FD
    * 0
    * 1
    * 2
    * cwd
    * txt (COMMAND file)
1. TYPE
1. DEVICE
1. SIZE 
1. NODE inode
1. NAME (COMMAND file)

[lsof]: https://linuxtools-rst.readthedocs.io/zh_CN/latest/tool/lsof.html