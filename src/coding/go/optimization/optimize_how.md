# Optimization `golang` App

>Goal  

To find out where bottlenecks lie and specificly optimize the bit.

    * Benchmarking
    * Profiling

>Benchmarking  
```go
import "testing"

func BenchmarkTargetFunc(b *testing.B) {
    for i:=0; i< b.N; i++ {
        TargetFunc()
    }
}
```
`$ go test -run=Bench -bench=.` will kick off benchmarking flow with arguments passed on.
* -run=regexp: filter which function the batch includes

>Output
```shell
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkCalculate-8 2000000000 0.30 ns/op
PASS
ok _/Users/elliot/Documents/Projects/tutorials/golang/go-testing-tutorial  0.643s
```

>Explain

`200000000` times of benchmark   
`0.30 ns/op` avg time each benchmark
`0.643s` entire benchmark  

>Profiling  

