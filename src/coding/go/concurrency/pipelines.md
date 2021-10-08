# Pipelne and Fan-in/Fan-out Mode

## Pipeline

A Pipeline is generally
> _`upstream`_ ==inbound==> _`procss`_ ==outbound==> _`downstream`_  

With 
* `upstream` alias `source` `producer`  
* `downstream` alias `sink` `consumer`  

## Upstrea, Producer and Source
```go
// the source is definite
// what about infinite case
func producer(nums ...int) <-chan int {
    // create an upstream channel and later return
    out := make(chan int)
    go func() {
        // write values into the returned channel
        for _, n := range nums {
            out <- n
        }
        // close when done
        close(out)
    }()
    return out
}
```

## Processor, Stage in-between
1. read values from upstream channel
1. process in Goroutine and write to downstream channel
1. return downstream channel
```go
func process(in <-chan int) <-chan int {
    // still create a downstream channel
    out := make(chan int)

    go func() {
        // process and write values into downstream in Goro
        for n := range in {
            out <- n * n
        }
        // close when doen
        close(out)
    }()

    return out
}
```


## Fan-in Merge
* `Fan-out`: distribute `1` goro for each each upstream channel
* `Fan-in`: merge into `1` wait goro

```go
func merge(cs ...<--chan int) <-chan int {
    // wait signal
    var wg sync.WaitGroup

    // fan-ined downstream channel
    out := make(chan int)

    collect := func(c <-chan int) {
        // collect from upstream channel
        for n := range c {
            out <- n
        }
        // notify when done via WaitGroup
        wg.Done()
    }

    // add all wait signals
    wg.Add(len(cs))

    // collect from Goro
    for _, c := range cs {
        go collect(c)
    }

    // wait in another Goro
    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

## Stop Short and Explicit Cancellation
By `context.Context` or `chan struct{}`  

```go

func producer(ctx context.Context, nums ...int) <-chan int {
    // downstream channel
    out := make(chan int)

    go func() {
        // must close after use or cancelled
        defer close(out)
        for _, n := range nums {
            select {
                case out <- n:
                // check whether cancelled
                case <- ctx.Done(): return
            }
        }
    } ()

    return out
}

func process(ctx context.Context, in <-chan int) <-chan int {
    // downstream channel
    out := make(chan int)

    go func() {
        // must close after use or cancelled
        defer close(out)
        for n := range in {
            select {
                case out <- n*n:
                // check whether cancelled
                case <- ctx.Done(): return
            }
        }
    } ()

    return out
}

func merge(ctx context.Context, cs ...<-chan int) <-chan int {
    // cond lock
    var wg sync.WaitGroup

    // downstream channel
    out := make(chan int)

    collect := func(c <-chan int) {
        // signal done when finish
        defer wg.Done()

        for n := range c {
            select {
                case out <- n:
                // check whether cancelled
                case <- ctx.Done(): return
            }
        }
    }

    // init cond lock
    wg.Add(len(cs))

    for _, c := range cs {
        // distribute over downstream
        go collect(c)
    }

    // goro to wait on cond lock and close downstream channel
    go func() {
        wg.Wait()
        close(out)
    } ()

    return out
}
```

_`main.go`_
```go
func main() {

    // setup context
    ctx, cancelFn := context.WithCancel(context.Background())
    time.AfterFunc(2 * time.Second, func() {
        fmt.Println("2 seconds up, cancel it")
        // cancel after 2 seconds
        cancelFn()
    })
    // above equals to
    // ctx, cancelFn := context.WithTimeout(context.Background(), 2 * time.Second)

    in := producer(ctx, 1,2,3,4,5,7)
    c1 := process(ctx, in)
    c2 := process(ctx, in)

    out := merge(ctx, c1, c2)

    for n := range out {
        fmt.Printf("collect %d\n", n)
    }

    fmt.Println("all done")
}

```