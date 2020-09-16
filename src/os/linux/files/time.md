# `atime`, `ctime` and `mtime`

Time | Name | Description | usecases
---|---|---|---
atime | access time | last time when file is accessed | `cat`, `less`, `more` will not change `atime`
ctime | change time | last time when file inode changed | `chmod` will change `ctime`
mtime | modify time | last time when file content changed | 

---

## Display
`stat FILE` can be used to show them,or  
`ls -l --time=ctime|atime` by default it uses `mtime`  
```bash
$ stat test_sed
  File: test_sed
  Size: 25              Blocks: 0          IO Block: 512    regular file
Device: 2h/2d   Inode: 3659174698235511  Links: 1
Access: (0662/-rw-rw--w-)  Uid: ( 1000/   roger)   Gid: ( 1000/   roger)
Access: 2020-09-11 10:14:05.349721500 +0800
Modify: 2020-09-11 14:08:40.582316500 +0800
Change: 2020-09-11 14:09:23.852180500 +0800
 Birth: -
$
```

## Change 
`touch -a FILE`  
`touch -a -t CCYYMMDDhhmm.ss FILE`  
* -a : set `atime` only
* -m : set `mtime` only
>NOTE:  
when `atime` or `mtime` is changed and it is newer than original `ctime`, `ctime` will be updated with the same value.