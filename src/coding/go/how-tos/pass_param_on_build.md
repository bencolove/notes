# Passing Params on Build
Injecting values into binaries during building process.
```go

var (
    buildstamp = ""
    githash = ""
    runtimeversion = ""
)
func main() {
    args := os.Args
    if len(args) == 2 && (args[1] == "--version" || args[1]=="-v") {
        fmt.Printf("UTC Build Time: %s\n", buildstamp)
        fmt.Printf("GO Version: %s\n", runtimeversion)
    }
}
```

> pass values by `-X`  
```sh
flags="-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X 'main.runtimeversion=$(go version)'"

$ go build -ldflags "$flags" -x -o executable main.go
```
-x: print commands
-o: output binary name