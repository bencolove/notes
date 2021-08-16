# Build Directives
With go tags, you can build go binaries with different configuration/setup:
1. for debug/release
1. for cross platform implementations

## Differenct configuration/setup for debug/release versions
- project \
 |- debug_config.go
 |- release_config.go

 -- _`debug_config.go`_ --
```go
//+build debug

package main

func GetConfigurationString() {
    return "it is debug"
}
```

 -- _`release_config.go`_ --
```go
//+build !debug

package main

func GetConfigurationString() {
    return "it is release"
}
```

> build with `debug` config  
`$ go build -tags debug`

> build with `release` config (non debug)  
`$ go build`