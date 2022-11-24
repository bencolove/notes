# Reuse Buffer
- decrese `alloc/op`
- speed up 
- something related to escape from stack?

## `sync.pool`

Stratigies:
1. get cache from a pool
2. operate with it
3. copy result to new one for return
4. put back in the pool

```go
// Wrapper
type buffer struct {
    data []byte
}
// Pool
var bufpool = sync.Pool{
    New: func() any {
        return &buffer{data: make([]bute, 0, 1024)}
    }
}

// 1. get one cache from the pool
buf := bufpool.Get().(*buffer)

// 2. operate on it(write to it)
data := encode(buf.data)

// 3. copy result for return
newBuf := make([]byte, len(data))
copy(newBuf, buf)

// 4. return copy and put back to pool
buf.data = data
bufpool.Put(buf)

```