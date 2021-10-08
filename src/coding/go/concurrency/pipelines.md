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
* `Fan-out`: distribute
* `Fan-in`: merge

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
        go output(c)
    }

    // wait in another Goro
    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```