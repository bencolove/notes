# `reflect`
[draveness][draveness-golang-reflect]
* static `Type` `reflect.TypeOf()`
* runtime `Value` `reflect.ValueOf()`
* three laws
* `interface{}` runtime representation
* test implementing interfaces
* runtime method invocation
* runtime parameter check

```go
// `Type` is an interface and may be implemented
// by _type used by iface/eface
type Type interface {
    Align() int
    FieldAlign() int
    Method(int) Method
    ...
}
// `Value` is a struct
type Value struct {
    // no Export fields
}
// methods come with struct reciever 
// guess struct is a reference type
func (v Value) Addr() Value
...
```

## Three Laws
Most functions from `reflect` are applied soly on `reflect.Type` or `reflect.Value` reflection objects. And they look like:

concrete objects --> interface value --> reflection objects

>1. refection objects come from interface values  

When getting `reflect.Type` or `reflect.Type` by `reflect.TypeOf(interface{})` or `reflect.ValueOf(interface{})`, an implicit conversion from concrete type to `interface{}` happend behind the scene.

>2. get interface values from reflection objects

`reflect.Value.Interface() interface{}` and from there, type assertion can be used like `reflect.Value.Interface().(int)`

Concrete | Method | Interface Value | Method | Reflection
---|---|---|---|---
concrete | assignment/copy <br> ----> | interface value  | `Type/ValueOf` <br> ----> | reflection objects  
concrete | type assertion <br> <---- | interface value | `Interface()` <br> <---- | reflection objects

>3. value updatable only when assignable  
Only assignable `reflect.value` (pass by reference) can be updated via reflection objects

```go
i := 1
// assignable Value
v := reflect.ValueOf(&i)
// Value.Elem() to get pointer/reference
v.Elem().SetInt(10)
```

## Runtime `interface{}`
`interface{}` in runtime look pretty much the same as in compile time.
```go
type emptyInterface struct {
    typ  *rtype
    word unsafe.Pointer
}
```

When `var i interface{} = v` or `reflect.ValueOf(v)`, implicit conversion to `interface{}` is done during compile time, and what `reflect.TypeOf()` and `reflect.ValueOf()` are receiving are `emptyInterface` data struct:

```go
func TypeOf(i interface{}) Type {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	return toType(eface.typ)
}

func toType(t *rtype) Type {
	if t == nil {
		return nil
	}
	return t
}

func ValueOf(i interface{}) Value {
	if i == nil {
		return Value{}
	}

	escapes(i)

	return unpackEface(i)
}

func unpackEface(i interface{}) Value {
	e := (*emptyInterface)(unsafe.Pointer(&i))
	t := e.typ
	if t == nil {
		return Value{}
	}
	f := flag(t.Kind())
	if ifaceIndir(t) {
		f |= flagIndir
	}
	return Value{t, e.word, f}
}

```

>NOTE  
1. `rtype` from `reflect.rtype` implements `Type` interface 
1. cast from `interface{}` value to `emptyInterface` is done by:
`eface := (*emptyInterface)(unsafe.Pointer(&i))`
1. When construct `Value` by `ValueOf()`, data are copied to heap from between and therefore can be updated if allowed(exported and assignable)


## Test Implementing _interfaces_
Getting `reflect.Type` of _struct_ or _interface_:  
`reflect.TypeOf( (*intf)(nil) ).Elem()`

As to test whether a _struct_ implements an _interface_, one can:  
```go
targetInterface := reflect.TypeOf((*error)nil).Elem()
valuePtr := reflect.TypeOf(&CustomError{}) 
value := reflect.TypeOf(CustomError{})
// Or even more simplier
// valuePtr := reflect.TypeOf((*CustomError)(nil))
// value := reflect.TypeOf((*CustomError)(nil)).Elem()
// true
valuePtr.Implements(targetInterface)
// false
value.Implements(targetInterface)
```
>Note:  
`Type.Implements` will compare both methods to decide whether it implements all methods target defined.

## Runtim Method Invocation
Dynamically call a method in runtime is not easy:
```go
func Add(a, b int) int { return a + b }

func main() {
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int()) // #=> 1
}
```
>What it logically does is:
1. get target method by `reflect.ValueOf` and pass in method interface value
1. get parameters by `reflect.Type.NumIn`
1. set parameter by `reflect.ValueOf` in order
1. call method's `reflect.Value.Call` and passin parameters
1. verify results

[draveness-golang-reflect]: https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/