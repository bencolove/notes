# Redirection Insight

`int pipe(int fds[2])` creates a channel from `fds[0]` to `fds[1]`  
`int dup2(int oldfd, int newfd)` makes newfd be the copy of oldfd  
`ssize_t write(int fd, const void *buf, size_t count)` writes up to count bytes from buf to fd  

```bash
$ strace -e open,dup2,write,pipe -f bash -c 'cat < <(date)'
strace: Process 2889 attached
[pid  2889] pipe([5, 6])                = 0
[pid  2889] dup2(5, 63)                 = 63
strace: Process 2890 attached
[pid  2889] dup2(5, 0)                  = 0
[pid  2890] dup2(6, 1)                  = 1
strace: Process 2891 attached
[pid  2891] write(1, "Fri Sep  4 17:56:07 CST 2020\n", 29) = 29
[pid  2889] write(1, "Fri Sep  4 17:56:07 CST 2020\n", 29Fri Sep  4 17:56:07 CST 2020
 <unfinished ...>
[pid  2891] +++ exited with 0 +++
[pid  2889] <... write resumed> )       = 29
[pid  2890] --- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=2891, si_uid=1000, si_status=0, si_utime=0, si_stime=0} ---
[pid  2890] +++ exited with 0 +++
[pid  2889] --- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=2890, si_uid=1000, si_status=0, si_utime=0, si_stime=0} ---
[pid  2889] +++ exited with 0 +++
--- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=2889, si_uid=1000, si_status=0, si_utime=0, si_stime=0} ---
+++ exited with 0 +++
```

```bash
$ strace -e open,dup2,write,pipe -f bash -c 'cat <(date)'
pipe([5, 6])                            = 0
dup2(5, 63)                             = 63
strace: Process 2895 attached
[pid  2895] dup2(6, 1)                  = 1
strace: Process 2896 attached
[pid  2896] write(1, "Fri Sep  4 17:56:56 CST 2020\n", 29) = 29
[pid  2894] write(1, "Fri Sep  4 17:56:56 CST 2020\n", 29Fri Sep  4 17:56:56 CST 2020
) = 29
[pid  2896] +++ exited with 0 +++
[pid  2895] --- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=2896, si_uid=1000, si_status=0, si_utime=0, si_stime=0} ---
[pid  2895] +++ exited with 0 +++
--- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=2895, si_uid=1000, si_status=0, si_utime=0, si_stime=0} ---
+++ exited with 0 +++
```

```bash
$ strace -e open,dup2,write,pipe -f bash -c 'date | cat'
pipe([5, 6])                            = 0
strace: Process 2900 attached
[pid  2900] dup2(6, 1)                  = 1
strace: Process 2901 attached
[pid  2901] dup2(5, 0)                  = 0
[pid  2900] write(1, "Fri Sep  4 17:57:10 CST 2020\n", 29) = 29
[pid  2900] +++ exited with 0 +++
[pid  2901] write(1, "Fri Sep  4 17:57:10 CST 2020\n", 29Fri Sep  4 17:57:10 CST 2020
) = 29
[pid  2901] +++ exited with 0 +++
--- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=2900, si_uid=1000, si_status=0, si_utime=0, si_stime=0} ---
+++ exited with 0 +++
```