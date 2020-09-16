# `bash` Basics

## Debugging
`strace -e open,dup2,pipe,write -f bash -c PROGRAM [ARGS]`
It will display lower systemcalls involved like:  
```bash
$ strace -e open,dup2,pipe,write -f bash -c 'cat <<EOF
> test
> EOF'
strace: Process 1562 attached
[pid  1562] write(6, "test\n", 5)       = 5
[pid  1562] dup2(6, 0)                  = 0
[pid  1562] write(1, "test\n", 5test)   = 5
[pid  1562] +++ exited with 0 +++
--- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=1562, si_uid=1000, si_status=0, si_utime=0, si_stime=0} ---
+++ exited with 0 +++
```
`write(1, "test\n", 5test)` means systemcall `write` content `"test\n"` to `1` (/dev/stdin).

>Explains:  
`ssize_t write(int fd, const void *buf, size_t count)` write to a file descriptor  
`int dup2(int oldfd, int newfd)` duplicate file descriptor  


## Input TEXT 
* `<<<` (here-string)
* `<<` (here-document)
* `<` (STDIN redirect from FILE)  

1. `<<` aka the **here-document** structure for capture TEXT from **STDIN** and feed to **STDIN** as a whole.
```bash
$ wc << EOF
> one two three
> four five
> EOF
>>> 2 5 24
# line-count word-count byte-count for STDIN
```
compared to
```bash
$ wc
one two
three
(Ctl-D)
>>> 2 3 13
```
Under the hood, **here-document** is implemented by write into a temporary file usually like `/tmp/sh-thd.<random string>`


1. `<<<` aka the **here-string** structure for consume directly the following pre-made TEXT
`$ bc <<< 5*4` 
Observe:
```bash
$ ls -l /proc/self/fd <<< "TEST"
lr-x------ 1 roger roger 0 Sep  4 11:43 0 -> '/tmp/sh-thd.XiybQ9 (deleted)'
lrwx------ 1 roger roger 0 Sep  4 11:43 1 -> /dev/tty3
lrwx------ 1 roger roger 0 Sep  4 11:43 2 -> /dev/tty3
lr-x------ 1 roger roger 0 Sep  4 11:43 3 -> 'pipe:[156]'
```
`/proc/self/fd` contains file symbolic links with current shell process. And there is a **deleted** one, which is the deleted temporary file used to can `<<< "TEST"` **here-string** and right next used as **STDIN**.


## Command Substitution `$(commands)` and **\`commands\`**  
> the **back-tick** **\`** is not quotation mark    
> and **\`commands\`** <=>  `$(commands)`  
> commands := command[;command]  
> command := <bash-command>  

The commands will be evaluated and will be replaced with the result before executing the rest parts of the surrounding command.

So `$(ls | wc -l)` will be replaced with a number string (substitution).

```bash
$ echo num of files: $(ls | wc -l)
>>> num of files: 57
```

---
## Group Commands: PIPE or Process Substition
Commands can be grouped by `|` PIPE or `()` process substitution.

While they have similary usage like:  
```bash
$ cat <(date)
$ cat < <(date)
$ date | cat
# three commands have identical output
# cat <(date): direct proc-subs with cat's STDIN
# cat < <(date): proc-subs into a FILE first and redirect I/O to cat's STDIN
# date | cat: simple PIPE as we know it
```
process substition can help when two FILEs are involved:  
`diff <(ls /bin) <(ls /usr/bin)`

## **Pipe** `|`
As depicted by [stdio_buffer], a PIPE `|` operation is taken via buffers in kernel land.
For example, when executing `cmd1 | cmd2`, the shell is instructing kernel to bufferedly read STDOUT from `cmd1` and then write to `cmd2`'s STDIN. The buffers involved may be controlled by some means. 

Operator `|` (pipe) only pass through **STDOUT**, not including **STDERR** which directly outputs on terminal or be redirected.

## Process Substitution
While `$(commands)` will be replaced with its results as **TEXT**, `<(commands)` will first output its results to a tmporary file `/dev/fd/nn` and it will be replaced with the file name. It effectively looks like to users that it replaces **STDIN** for commands accepts omitted **File** argument as `-` **STDIN**.

```bash
$ echo <(date)
# echo consumes string contents
# it outputs file name, only available in the current command(shell process)
>>> /dev/fd/63

$ cat <(date)
# cat reads FILE or - (STDIN if omitted)
>>> Fri Sep 4 10:11:22 CST 2020
```
---

## Input of Bash Commands
Two categories mainly for bash scripts or built-in commands:
1. STDIN, output end of PIPE
1. FILEs

>**NOTE:**  
>Most commands do not deal with `TEXT` contents, they expect `FILEs` (including `STDIN`). Only that PIPE `|` is passing TEXT contents from upstream's `STDOUT` through to downstream's `STDIN` **NOT** FILE handle


Most commands will distuiguish them by looking at their manual. Among them, some will regard an omitted FILE argument as `-` (STDIN)


## Ways of Conversion among **TEXT**, **FILE** and **STDIN**
1. **TEXT** --> **FILE**
`<(echo TEXT)` 
1. **FILE** --> **TEXT**
`cat FILE`

1. **STDIN** --> **FILE**


[command-substitution]: http://www.gnu.org/software/bash/manual/html_node/Command-Substitution.html
[stdio_buffer]: http://www.pixelbeat.org/programming/stdio_buffering/