# Looping
* for loop
* range loop

>Compile Time  
All `for-loop` and `range-loop` will eventually converted to optimized `for-loop`. 
Further more, before entering the counterpart loop, `range-loop` will **COPY** its contents to a temporary object and loop over that.
Last, when `for i, v := range array/slice/map`, a second copy is introduced when again value is copied to `v`.

Because assigning an `array` costs deep copy while assigning a `slice` will only copy its `SliceHead`.

>Better Practice  
Better off to loop over slice other than array: `for i := range arr[:]` and access its value by index (no need to copy to second variable within range loop)
Same for `map`, since assignment will no deep copy map's content.

```go
// test array assignment
arr := [...]int{1,2,3}
// deep copy
arr1 := arr
arr1[1] = 0
// [1 2 3], [1 0 3]
fmt.Printf("src:%+v, assigned:%+v \n", arr, arr1)

// test slice assignment
arr = [...]int{2,2,2}
s1 := arr[:]
// no deep copy
s2 := s1
s2[1] = 0
// [2 0 2], [2 0 2], [2 0 2] 
fmt.Printf("raw:%+v, slice:%+v, assigned:%+v \n", arr, s1, s2)

// test map assignment
m := map[string]int {"a":1}
// no deep copy
mcp := m
mcp["a"] = 2
// map[a:2], map[a:2]
fmt.Printf("src:%+v, assigned:%+v \n",m, mcp)
```

## Closure within Loop
```go
for i := range []int{1,2,3} {
    go func() {
        fmt.Print(i)
    } ()
}
```
It will print `3 3 3`, since goroutine captures(by reference) `i` from range's closure which effectively point to same variable(address).

To avoid this, either use local copied variable or pass to goroutine by value:
```go
for i := range []int{1,2,3} {
    v := i
    go func() {
        fmt.Print(v)
    } ()
}
// Or
for i := range []int{1,2,3} {
    go func(v int) {
        fmt.Print(v)
    } (i)
}
``` 