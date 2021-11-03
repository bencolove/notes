# `find`
`find -options [PATH...] [EXPR] `  
> defaults:  
PATH $PWD  
EXPR -print  

EXPR := OPERATORS, OPTIONS, TESTS, ACTIONS
OPERATORS := (EXPR), ! EXPR, -not EXPR, -and EXPR, -or EXPR

OPTIONS := 
* -daystart
* -follow
* -regextype 
* -depth
* -maxdepth LEVELS
* -mindepth LEVELS
* -mount
* -noleaf
* -xdev

TESTS | Description | Cases
---|---|---  
-name PATTERN | match file name |
-path PATTERN | match file path |
-type [bcdpflsD] | match file type | `d`(ir) <br> `f`(ile) <br> `p`(ipe)
-size N[bcwkMG] | size bigger than ? |
-uid N | match UID |
-user NAME | match USER |
-group GROUP | match GROUP |
-perm [-/]MODE | match permission | 
-regex PATTERN | match name? using regex |
time-related | |
-amin N | `atime` in last N `minutes` |
-atime [+]N | `atime` in last N `days` or beyond N `days`|
-anewer FILE | `atime` newer than FILE's | 
-empty | test empty |


ACTIONS :=  
* -print | -print0
* -printf FORMAT
* -fprintf FILE FORMAT
* -exec COMMAND
* -exec COMMAND {}
* -ok COMMAND

## Search
1. by name -- find in current location recursively for files/dirs ending with ".log"  
`find . -name "*.log"`
2. by permission  
`find . -perm 777`
3. by type  
`find . -type f -name "*.log"`
4. by size -- files bigger than 1K   
`find . -size +1000c -print`
5. by `atime`  
```bash
# create files and set atime by touch
$ touch -a -t 202009101133.33 one_day_ago
$ touch -a -t $(date -d "-2 days -8 hour" +'%Y%m%d%H%M.%S') two_days_ago
$ touch -a -t $(date -d "-3 days -8 hour" +'%Y%m%d%H%M.%S') three_days_ago
$ touch -a -t $(date -d "-2 hour" +'%Y%m%d%H%M.%S') within_one_day

# view atime
$ ls -l --time atime

# result
$ find -type f -atime -4
./one_day_ago
./three_days_ago
./two_days_ago
./within_one_day
$ find -type f -atime -3
./one_day_ago
./two_days_ago
./within_one_day
$ find -type f -atime -2
./one_day_ago
./within_one_day
$ find -type f -atime -1
./within_one_day
$ find -type f -atime 0
./within_one_day
$ find -type f -atime 1
./one_day_ago
$ find -type f -atime 2
./two_days_ago
$ find -type f -atime 3
./three_days_ago
$ find -type f -atime 4

$ find -type f -atime +0
./one_day_ago
./three_days_ago
./two_days_ago
$ find -type f -atime +1
./three_days_ago
./two_days_ago
$ find -type f -atime +2
./three_days_ago
$ find -type f -atime +3

```
>What's all about `-atime [-]n`:  

N values | Meaning
---|---
-4 | [4 days, now) | between -4 days and now
-1 | [1 day, now) same as 0 | between -1 day and now
0  | [1 day, now) | 
1  | [2 day ago, 1 day ago] | between -2 days and -1 day
2  | [3 day ago, 2 day ago] | between -3 days and -2 days
+0 | [, 1 day ago] | > 1 day
+1 | [, 2 day ago] | > 2 days


---
## Actions
### `-exec COMMAND {} \;`  
`find . -type f -exec ls -l {} \;`
```bash
# find and list
$ find -type f -exec ls -l {} \;
-rw-rw-rw- 1 roger roger 0 Sep 11 15:39 ./one_day_ago
-rw-rw-rw- 1 roger roger 0 Sep 11 15:43 ./three_days_ago
-rw-rw-rw- 1 roger roger 0 Sep 11 15:38 ./two_days_ago
-rw-rw-rw- 1 roger roger 0 Sep 11 15:40 ./within_one_day

# {} will be replaced by result from find one at a time
$ find -type f -exec ls -l tag{} \;
ls: cannot access 'tag./one_day_ago': No such file or directory
ls: cannot access 'tag./three_days_ago': No such file or directory
ls: cannot access 'tag./two_days_ago': No such file or directory
ls: cannot access 'tag./within_one_day': No such file or directory

# safe mode, prompt for y/n before exec COMMAND
$ find -type f -ok ls -l {} \;
< ls ... ./one_day_ago > ? y
-rw-rw-rw- 1 roger roger 0 Sep 11 15:39 ./one_day_ago
< ls ... ./three_days_ago > ? n
< ls ... ./two_days_ago > ? n
< ls ... ./within_one_day > ? n   

# with grep
$ find /etc -name "passwd*" -exec ls {} \;
/etc/cron.daily/passwd
/etc/pam.d/passwd
/etc/passwd
/etc/passwd-
find: ‘/etc/polkit-1/localauthority’: Permission denied
find: ‘/etc/ssl/private’: Permission denied

$ find /etc -name "passwd*" -exec grep "root" {} \;
root:x:0:0:root:/root:/bin/bash
root:x:0:0:root:/root:/bin/bash
find: ‘/etc/polkit-1/localauthority’: Permission denied
find: ‘/etc/ssl/private’: Permission denied
```
>NOTE:  
The `{}` in `-exec COMMAND {} \;` or `-ok COMMAND {} \;` will be replaced as it is from `find`'s result one at a time. So you will can do `mv`, `cp`, even `gred` with `{}` as long as it is a FILE path.

>NOTE:  
It is FOREVER recommended to use `ls` in `-exec` before doing `-ok` in safe mode when modifying a thing. 

---

## Advanced
Sometimes