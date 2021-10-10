# Channel Patterns
* data channel gather with error channel (combine them into one struct)
* priority over multiple channel-select (extra code)

## Data and Error Channels
Combine them into one struct and pipeline it

>DON'T  
```go
dCh := make(chan int)
eCh := make(chan error)

if err != nil {
    select {
        case <- ctx.Done():
        case ech <- err:
    }
} else {
    select {
        case <- ctx.Done():
        case dCh <- data:
    }
}
```

>DO 
```go
type Result struct {
    Data int
    Err error
}

out := make(chan *Result)

select {
    case <- ctx.Done():
    case out <- &Result{Data: data, Err: err}:
}
```

## Priority on Multiple Ready Channels
What we want is: when both ch1 and ch2 are ready, we prefer ch1 over ch2 (process data from ch1 first)

```go

for {
    select {
    case <-stopCh:
        return
    case d1 := <- ch1:
        process(d1)
    case d2 := <- ch2:
    priority:
        for {
            select {
            case d1 := <- ch1:
                process(d1)
            default:
                // break in for-select only breaks out select clause
                // not outer for, same like continue
                break priority
            }
        }
        process(d2)
    }
}

```