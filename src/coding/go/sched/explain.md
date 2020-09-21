# Explain Scheduler of Goroutine

>The `GMP` story  
* G goroutine  
    who's job is defined by go `func` and holds context like stack frames  
* M os-thread
    who's actually executing the job (goroutines)  
* P processor  
    **`NOTE`**: it is not the CPU processor. It is a struct holding a local list of `G`. It must be tied to a `M` in order to execute `G`s.


---

## Tools
* `trace`
* `debug`

>`trace`  
`trace` comes with runtim info dump togather with a WebUI to inspect into it.

`go run trace.go && go tool trace trace.out`
The `View trace` page should be viewable as described at `http://127.0.0.1:33479`

>`debug`  
`debug` goes in a non-intrusive way unlike `trace`.  

`$ go build debug.go`

    $ GODEBUG=schedtrace=1000 ./trace2 
    SCHED 0ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=1 idlethreads=1 runqueue=0 [0 0]
    Hello World
    SCHED 1003ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
    Hello World
    SCHED 2014ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
    Hello World
    SCHED 3015ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
    Hello World
    SCHED 4023ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
    Hello World

>Results  
gomaxprocs := `P`  
idleprocs := ?
threads := `M` of os-threads (including some like 'sysmon')   
idlethreds := threads that without active `P`, so hanged  
spinningthreads := threads with `P` but no `G` available  
runqueue := global `G` queue  

>So  
workingprocs := gomaxprocs - idleprocs  
workingthreads := workingprocs - spinningthreads

## Webpage Archive

![part-1](img/golang-gmp-1.png)
![part-2](img/golang-gmp-2.png)


[example]: https://learnku.com/articles/41728