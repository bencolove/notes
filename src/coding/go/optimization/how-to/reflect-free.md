# How to Avoid Reflection

ref: [go-json]
---


```go
// Destructure an interface{} into two pointers
type emptyinterface struct {
    typ unsafe.Pointer
    ptr unsafe.Pointer
}

var typeToEncoder = map[uintptr]func(unsafe.Pointer)([]byte, error){}

// Do not rely on switch on v.(type)
// Instead compare the pointer to a fixed type
func Marshal(v interface{}) ([]byte, error) {
    iface := (*emptyinterface)(unsafe.Pointer)(&v)

    typeptr := uintptr(iface.typ)

    if enc, exists := typeToEncoder[typeptr]; exists {
        return enc(iface.ptr)
    }
}
```

[go-json]: https://github.com/goccy/go-json