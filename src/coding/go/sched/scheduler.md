# scheduler

The:
Acronym | Usage
---|---
M (thread) | OS thread, managed by OS 
G (goroutine) | task
P (processor) | scheduler on current thread for goros

## `G`
_--runtime.g--_
```go
type g struct {
    // stack 
    stack stack // stack range [stack.lo, stac.hi]
    stackguard0 uintptr

    // preempt
    preempt         bool // flag
    preemptStop     bool 
    preemptShrink   bool

    _panic  *_panic // panic call list
    _defer  *_defer // defer call list

    m       *m      // occupied thread struct
    sched   gobuf   // status
    atmoicstatus uint32
    goid         int64
    ...
}

type gobuf struct {
    sp  uintptr     // stack pointer
    pc  uintptr     // program counter
    g   guintptr    // owner
    ret sys.Uintreg // return value?
    ...
}
```

>gorotine status

Status | Meaning | Code | Stack | list | M | P
---|---|---|---|---|---|---
`_Gwaiting` | suspended(blocked, may be on _channel_) and wait to be into `runq` | x | x | x | x |x
`_Grunnable` | in `runq` and wait for `schedule` |x|x|o|x|x
`_Grunning` | |o|o|x|o|o
`_Gsyscall` | blocked from `syscall` | x|o|x|o|x
`_Gpreempted` | blocked by preempt | x|x|x|x|x

---
## `M`
OS thread limited by:
* maximum 10000
* most blocked/waiting
* GOMAXPROCS active ones
* GOMAXPROCS default to num of cores

_--runtime.m--_
```GO
type m struct {
    g0   *g // stack owner
    curg *g // current goro
    
    p       puintptr // current P
    nextp   puintptr // next available P
    oldp    puintptr // last P

    ...
}
```
The `g0` goro is a special one the first gorotine when creating the `M` and duty for memory allocation, construct the `M` and CGO.

## `P`
The coordinator between `M` and `G`. Scheduler/manager of goros on M.

_--runtime.p--_
```go
type p struct {
    m muintptr // associated M

    // glist
    runqhead uint32         // glist ring buffer
    runqtail uint32         //
    runq     [256]guintptr  //
    runnext  guintptr       // next goro

    status  uint32? // status
}
```

Status | Meaning | M
---|---|---
_Pidle | empty glist | x
_Prunning | executing | o
_Psyscall | m is blocked | o?
_Pgcstop | blocked by gc | o
_Pdead | not available | x

## Flow of scheduler
1. init
1. create goro
1.

### initialize _scheduler_
`runtime.schedinit()` does:  
1. max thread count 10000
1. read env var GOMAXPROCS and make sure of it
1. global processors `allp` match _GOMAXpROCS_
1. associate `m0` to `p0`
1. idlize other `P` other than `p0` by _Pidle

### create gorotine
`go func()` will be transpiled to
1. `runtime.newproc` get
  * args (sp and size)
  * caller goro (get cached or alloc new `g` struct)
  * caller pc
1. `runtime.newproc1`
 * get cached or alloc new `g` struct
 * copy args to `g`'s stack
 * setup `g`'s _Grunnable, sp, pc
 * put `g` into current `p`'s 
    * runnext (schedule right away)
    * runq or schd's runq(gloabl one) (line up)

### loop of scheduler
`runtime.schedule` will find and execute **ONE** `g` once and call it again to loop.

1. `runtime.mstart` initialize `stackguard0` and `stackguard1`
1. `runtime.mstart1` kicks start `runtime.schedule`
    1. find nextg by
        1. global `g` with possibility
        2. current/local `m`'s runq (next `g`)
        3. search for runnable `g` 
            * search global runq
            * search select runq
            * steal from other `p` 
            * blocked and wait
    1. `runtime.execute` found `g`
        1. associate `g` to current `m`
            * `g.m` and `m.curg`
        1. sp (stack lo)
        1. `gogo` schedule on current `m`(OSthread)
1. `runtime.goexit` to clean executed `g`
    * transmitted to _Gdead
    * clean associated `m`, sp, pc
    * put back `g`(struct) to gFree list
    * `runtime.schedule` to loop over agin

>Scheduler kicked off timing(who called `schedule`):
1. `runtime.mstart` and `goexit0` as above
1. suspend `runtime.gopark` -> `park_m`
1. syscall `runtime.exitsyscall` -> `exitsyscall0`
1. cooperative `runtime.Gosched`(by user)
1. preemptive `runtime.sysmon`(by sysmon) -> `retake` ->`premptone`

### Suspend (by runtime)
Suspended by runtime using `gopark`, and runtime will:
1. prepare switch `m`'s curg `g`
1. make it _Grunning and sleep(not yet in `runq`)
1. `schedule` for another `g`
1. `goready` the suspended `g`(_Grunnable and into `runq`)

### Syscall (`m` OSThread will be blocked)
Those **BLOCKING** `syscall` will need to block the cuurent `m`(OSThread) so, it needs to separte `m` and `p` and let `m` to be blocked and `p` associated with another `m` to continue `g`s.

Flow:
1. `entersyscall`  
    prepare pc, sp
1. `reentersyscall`
    * current `m`(OSThread) will be blocked and sleep
1. `runtime.exitsyscall` after syscall finishes
    1. `exitsyscalfast`
    1. scheduler's goro `exitsyscall0`

>Whichever path is taken, `schedule` will be called. So after syscall, scheduling will happen after that.

### Cooperative
`runtime.Gosched` will ativel yield controll to other goros.

`Gosched` itself will not suspend current goro, instead, it relay to `schedule` to do the job after setting current `g` to _Grunnable and put it into runq.

### Preempt (by signal SIGUSR)

