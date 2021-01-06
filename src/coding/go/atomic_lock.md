# Lock with `atomic`

```go
func netpoolGenericInit() {
    // atomic check
    // first load flag
    if atomic.Load(&netpollInited) == 0 {
        // lock
        // mutex another goro acquiring netPollInitLock
        lock(&netpollInitLock)

        if netpollInited == 0 {
            netpollInit()
            // set flag
            atomic.Store(&netpollInited, 1)
        }

        unlock(&netpollInitLock)
    }
}
```