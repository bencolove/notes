# SliceHeader & StringHeader
1. Defs of them
1. Normal conversion between `[]byte` and `string`
1. Shared-backing-data conversion
1. save `struct` without parsing

>Conclusion  

When GO convers using `string([]byte)` and `([]byte)string` resulting into a copy of backing data for reasons:
1. `[]byte` mutable and `string` immutable
1. change immutable variables needs a copy of it

If space (memory) efficiency is needed, a direct map from `SliceHeader` (from `[]byte`) to `StringHeader` (from `string`) would work as long as they are **READ-ONLY**.

>Definitions  
```go
// reflect
type StringHeader struct {
    Data uintptr
    Len  int
}
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```
>NOTE  

By looking at the defs, `StringHeader` has one less Field than `SliceHeader` does and first two Fields `Data` and `Len` are aligned (same ocurrence order and type). It makes possible that direct mapping a `SliceHeader` to `StringHeader` works. That said, it is possible to have a `string` view of underlying `[]byte` without any copy.


## Shallow Copy

Make a `string` out of `[]byte` without coping data:
```go
var byteSlice = []byte{'a', 'b', 'c'}
// byte slice has layout of SliceHeader
var sliceData = (*reflect.SliceHeader)(unsafe.Pointer(&byteSlice)).Data

// direct map `SliceHeader` to `StringHeader` (layout of `string`)
var str = *(*string)(unsafe.Pointer(&byteSlice))
// check backing data
// string has layout of StringHeader
var strData = (*reflect.StringHeader)(unsafe.Pointer(&str)).Data

// they should be identical
if sliceData != strData {
    t.Fail()
}
```

## Normal Conversion
```go
var byteSlice = []byte{'a', 'b', 'c'}
var str = string(byteSlice)

// check whether they share underlying backing data
var byteSliceHeader = (*reflect.SliceHeader)(unsafe.Pointer(&byteSlice))
var strHeader = (*reflect.StringHeader)(unsafe.Pointer(&str))

// they have different backing data
if byteSliceHeader.Data == strHeader.Data {
    t.Fail()
}

// content of backing data are same
// Unicode string, backing data length are byte array length
// check length
if byteSliceHeader.Len != strHeader.Len {
    t.Fail()
}
// check content byte by byte
for i:=0; i<byteSliceHeader.Len; i++ {
    var b = *(*byte)((unsafe.Pointer)(byteSliceHeader.Data + uintptr(i)))
    var s = *(*byte)((unsafe.Pointer)(strHeader.Data + uintptr(i)))
    if b != s {
        t.Fail()
    }
}

```

## `struct` and `[]byte`
How to do the conversion from `struct` to `[]byte`:
1. `encoding/gob`
1. `encoding/json`
1. `encoding/binary`
1. *black magic* `unsafe.Pointer`

```go
type MyStruct struct {
    A int
    B int
}

var data = MyStruct{
    A: 1, B: 2,
}

var sizeMyStruct = int(unsafe.Sizeof(data))

var x reflect.SliceHeader
x.Len = sizeMyStruct
x.Cap = x.Len
x.Data = uintptr(unsafe.Pointer(&data))
// struct to []byte
var backup = *(*[]byte)(unsafe.Pointer(&x))

// []byte to struct
unsafe.Pointer(&backup)
```
