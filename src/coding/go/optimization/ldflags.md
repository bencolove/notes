# Golang `ldflags`

`go build -ldflags "-w -s -X main.Version=${INJECT} -X main.Build=${INJECT}"`
- `-s` strip off symbol table, no file name ane lineno from stack trace
- `-w` trim `DWARF` no `gdb` anymore
- `-X` wire data into final binaries