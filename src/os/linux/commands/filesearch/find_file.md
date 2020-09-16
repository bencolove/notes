# Find Files
* `which` executables from **`PATH`**
* `whereis` from .db
* `locate` look into .db
* `find` the ultimate tool looking into disk

>NOTE:  
>Search sources:
>* PATH
>* .db: only updated likely once a week, results may be outdated 
>* disk

---

## `which`
`which` searchs env **`PATH`** for the first occurence.

>NOTE:  
built-in will not be found since they won't on on PATH

## `whereis`  
`whereis [-bmsu] [-BMS <dir>... -f] <name>`
* -b: for binaries
* -m: for manuals and infos
* -s: for sources
* -B: define binary lookup path
* -M: define muanls lookup path
* -S: define sources lookup path


## `locate`  
`locate -l n PATTERN`  
`locate` seems to TEXT match PATTERN against the `.db`, so glob pattern `*` can be used.
* -l n: limit output lines

```bash
$ locate -l 2  "/usr/*pwd*"
/usr/bin/pwdx
/usr/lib/python3/dist-packages/twisted/python/fakepwd.py
$
```