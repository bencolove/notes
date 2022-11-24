# Unittesting in `golang`

[[quick-intro]]

---


>Test file  
```text
myproject/
 - calc.go
 - calc_test.go
 - main.go
 - main_test.go
```
Those `FILE_test.go` files are testing files which `go test` will include.

>Test function

- testing.T
    - `t.Run()` substest
    - table driven
    - `t.helper()`
- testing.B 
    - benchmarking
- testing.M 
    - testsuit
    - 
- http/net/httptest

```go
// test cases by table driven
func TestFunc(t *testing.T) {
    cases := []struct{
        Name string
        A, B int // params
        Expected int
    }

    for _, c := range cases {
        t.Run(c.Name, func(t *testing.T) {
            if got := Func(c.A, c.B); got != c.Expected {
                t.Fatalf("expected: %v, got %v", c.Expected, got)
            }
        })
    }
}
```

```go
// testsuit
func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    teardown()
    os.Exit(code)
}
```

```go
// Benchmarking
func BenchmarkHello(b *testing.B) {
    for i := 0 ; i < b.N; i ++ {
        tests
    }
}
// Result
type BenchmarkResult struct {
    N         int           // 迭代次数
    T         time.Duration // 基准测试花费的时间
    Bytes     int64         // 一次迭代处理的字节数
    MemAllocs uint64        // 总的分配内存的次数
    MemBytes  uint64        // 总的分配内存的字节数
}
// Reset timer
b.ResetTimer()
// Parallel
b.RunParallel(func(pb *testing.PB) {
    for pb.Next() {
        // gorotine of b.N
    }
})
// Run benchmark
go test -benchmem -bench .
```


```go
func TestTableCalculate(t *testing.T) {
    var tests = []struct {
        input    int
        expected int
    }{
        {2, 4},
        {-1, 1},
        {0, 2},
        {-5, -3},
        {99999, 100001},
    }

    for _, test := range tests {
        if output := Calculate(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
        }
    }
}
```

> httptest

```go
func TestConn(t *testing.T) {
    // test request
    req := httptest.NewRequest("GET", URL, nil)
    // test responseWriter
    w := httptest.NewRecorder()

    // test HttpHandler
    helloHandler(w, req)

    // verify by w
    bytes, _ := ioutils.ReadAll(w.Result().Body)

    if str := string(bytes); str != "hello world" {
        t.Fatal("expected %s, but got %s", "hello world", str)
    }
}
```

>Run test  

`$ go test -v`

>Text coverage  

`$ go test -cover`

>Text coverage visualize

`$ go test -coverprofile=coverage.out`  
`$ go tool cover -html=coverage.out`

>Tag test files for integration only

FILE := `*_integration_test.go`
```go
//go:build integration

package main_test
// ...testing go on
```

Run them exclusively
`$ go test -tags=integration`

[quick-intro]: https://geektutu.com/post/quick-go-test.html