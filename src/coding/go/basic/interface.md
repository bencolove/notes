# The `interface`
[From draveness][draveness-golang-interface]
* internal representation
* conversion
* recievers when implementing interfaces

Internal Repr | Interface Def | Meaning
---|---|--- 
iface | `type Name interface { DEFS }` | with methods definitions
eface | `interface{}` | empty interface

>Internal:  
```go
// interface{} consists of two elements: <type, elemPtr>
type eface struct {
    _type *_type
    data unsafe.Pointer
}
// *itab are function tables
type iface sturct {
    tab *itab
    data unsafe.Pointer
}
// Above all, they are both two words long in size
```

>`_type`  
`_type` is kinda like `reflect.Type` contains descriptions of the type like `size`, `hash` (for equality) etc.

>`itab`  
`itab` is 32 bytes wide like:
```go
type itab struct {
    inter *interfacetype // containing interface
    _type *_type // contained concrete type
    hash  uint32 // from `_type.hash`
    _     [4]byte
    fun   [1]uintptr // dynamic, refer to contained type's function pointer table
}
```

## 2. Conversion

>From concrete to interface
An struce of either `iface` or `eface` will be constructed.


>From interface to concrete  
`iface{}.(type) {}`

Recall that `iface.itab.hash` or `eface._type.hash` will be compared to statically compiled `_type.hash`, it knowns whether they are of same type.

## 3. Recievers
Programatically speaking, an interface can be implemented with:
1. `(s Struct)` struct reciever (pass by value)
1. `(p *Struct)` pointer reciever (pass by reference)

What makes the biggest difference is not how they are used (called) but the efficency impact. Under the hood when implementing an interface with struct reciever, a copy by value is implied everytime a method is invoked.

>Best practice  
Use pointer reciever to implement an interface

[draveness-golang-interface]: https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/