# File descriptors
Each _`PROCESS`_ is associated with a bunch of _`file descriptors`_ in
`/proc/$PID/fd`:
1. `0` STDIN
1. `1` STDOUT
1. `2` STDERR
1. others: created by `<(COMMAND)`, temporarily captured output into a file
1. `pipe: [number]` PIPES 

## TTY (screens for input and output)
> 1. `tty` (current output screen)  
/dev/pts/5

> 2. `ps -p $$ -o tty=`  
/dev/pts/5

## Prove
```shell
function getfd() {
  echo $(readlink -f "/proc/$$/fd/$1")
}

echo "[$1] PID=$$, stdin=$(getfd 0), stdout=$(getfd 1), stderr=$(getfd 2)" >> pids
```

RUN: `./cmd frist 2> error | ./cmd second`  
`cat pids`  
>OUTPUT
[second] PID=18726, stdin=/proc/18726/fd/pipe:[211606], stdout=/dev/pts/5, stderr=cwd/error
[first] PID=18725, stdin=/dev/pts/5, stdout=/proc/18725/fd/pipe:[211606], stderr=/dev/pts/5

Conclusion:  
1. Each command in a pipe line runs in a dedicated subshell(process) simutaneously with separated _`fd`_s
1. current sreen(tty) is denoted(linked to) as `/dev/pts/NUM` which is the same as `$(tty)`
1. `readlink -f /proc/$$/fd/FD` can tell which file the _`FD`_ is connected
1. STDIN of next command is connected to STDOUT of previous command in a pipeline by a pipe file located at `/proc/$$/fd/pipe:[NUM]` 

[exec]: https://www.itread01.com/p/1390349.html