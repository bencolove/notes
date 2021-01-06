# `netpoll`
`netpoller` is not a didicated thread nor a goroutine.

It bridges Golang's scheduler of goros and underlying platform's I/O multiplex system(_epoll_, _kqueue_, _select_, _IOCP_ etc).

>Golang -> OS  

From goro blocked on fileI/O?, netI/O and timeout(sleep), the runtime will check and `gopark` current goro then `schedule` another and submit(register) events to lower platform poll system.

>OS -> Golang  

`sysmon` will periodically `netpoll(0)` (poll the flatform's poller) and when existing ready goros, it will put them to global runq for scheduling.

>Check deadlines(timer or comes from context)

1. `schedule` then `findrunnable` will check kicked timers and execute associated function
1. `sysmon` will periodically sleep before next timer times out and execute it on dedicate `m`??? (no `P` associated).

The `netpoll` methods are implemented on individual platfroms and the flow goes:
1. init
1. add events
1. get triggered events

## Golang Structs
```go
type pollDesc struct {
    link *pollDesc

    lock mutex
    fd   uintptr

    rseq uintptr // fd reset
    rg   uintptr // pdReady goro or nil
    rt   timer
    rd   int64   // read deadline

    wseq uintptr // timer reset
    wg   uintptr // pdWait goro or nil
    wt   timer
    wd   int64   // write dealine
}
```

## Golang Methods
```go
func netpollinit()
func netpollopen(fd uintptr, pd *pollDesc) int32
func netpoll(delta int64) gList
func netpollBreak()
func netpollIsPollDescriptor(fd uintptr) bool
```

## Initialization
`runtime.netpollinit()` will delegate to platform specific implementation like `epoll` on Linux for:
1. `epollcreate1` to create new `epoll` descriptor
1. `runtime.nonblockingpipe` to create a golang pipe
1. `epollctl` to put the descriptor(intesting events) into epoll service.


## Add events
`internal/poll.pollDesc.init` will also create events and put into poller by `netpollopen` and `netpollclose` to remove interesting events from poller.

## Golang `netpoll` Event loop
It bridges underlying OS I/O multiplexer and Golang runtime's goro scheduling.
How:
1. suspend goro and wait for read/write ready events
1. 

The runtime will put the goro into sleep (by `gopark`) when its acquiring read/write not ready which then `schedule` to another goro. When next schedule the runtime checks whether ready event happend(set by netpoller) and decides to make it sleep again or schedule it.



