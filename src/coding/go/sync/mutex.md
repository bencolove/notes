# `sync`
* primtives
* patterns
    * fan-in
    * fan-out

Primtives | Usage
---|---
`sync.Mutex` | Mutual exclusion
`sync.RWMutex` | Read/Write mutex
`sync.WaitGroup` | 
`sync.Once` |
`sync.Cond` |

## How `sync.Mutex` works
```go
type Mutext struct {
    state int32
    sema  uint32
}
```

### `Mutext.Lock`
Locking up simply set _mutexLocked_ flag bit in _Mutex.state_ 
`atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked)`

If already locked, current goro should wait:
1. self spin (do not yield control)
2. 30 _PAUSE_ 4 times to take on CPU
3. check state
4. try to acquire the lock by `sync.runtime_SemacquireMutex` which will yield gosched.

>Self-spin  
* check
* spin

`sync.runtime_canSpin()` decides by looking at:
1. multi-cored
1. Spined less than four times
1. at least one P (goro processor) with empty goro list (possible quick kick in)

Spin means 30 times _PAUSE_ machine command keeping current CPU busy and avoid yield.

### `Mutex.Unlock`
Keep trying to clear flag bit _mutexLocked_ by atomic CAS.

If success, `sync.runtime_Semrelease` to notify(awake) one blocked goro.

### Sum up about `sync.Mutex`
* atomic.CAS to set/clear flag on mutext.state
* sync.runtime_SemacqureMutex to wait and yield
* sync.runtime_Semrelease to notify and wake one blocked goro