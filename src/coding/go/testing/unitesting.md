# Unittesting in `golang`

>Test file layout  
```text
myproject/
 - calc.go
 - calc_test.go
 - main.go
 - main_test.go
```
Those `FILE_test.go` files are testing files which `go test` will include.

>Test file
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
// +build integration

package main_test
// ...testing go on
```

Run them exclusively
`$ go test -tags=integration`