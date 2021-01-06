# Array and Slice

`array` is underneath data structure dont directly dealt with.

## Declare
DataType | Declare | Value | Comment
---|---|---|---
array | `var [3]int` | `[0 0 0]` | allocated on stack
slice | `var []int` | **`nil`** | reference type
pointer | `var *int` | **`nil`** | reference type
map | `var map[string]int` | **`nil`** | [reference type][golang-map]

## Copy
```go
arr := [...]int{1,2,3}
// deep copy
arr1 := arr
arr1[1] = 0
fmt.Printf("%+v, %+v \n", arr, arr1)

arr = [...]int{2,2,2}
s1 := arr[:]
s2 := s1
s2[1] = 0
fmt.Printf("%+v, %+v, %+v \n", arr, s1, s2)
```
>Output  
[1 2 3], [1 0 3]   
[2 0 2], [2 0 2], [2 0 2]

It turns out, when assign an `array` like `arr_cp := arr` will deep copy array's underlying data structure, on the other hand assigning a `slice` like `sli_cp := sli` will only copy a `SliceHead` of three words !!

## add element
`newSlice := append(slice, elements...)`

Due to the `cap` left for the old array underneath, the reference(address) to array **MAY** change(expend).

The `SliceHead` is changed in these case, so always to reassign back to the variable.


## remove from slice
To remove an element at `idx`  
`newSlice := append(oldSlice[:idx], oldSlice[idx+1:]...)`

## deep copy
1. `make` + `copy`
`newSlice := make([]T, len(oldSlice))`  
`copy(newSlice, oldSlice)`  
2. `append`
`newSlice := append([]T(nil), oldSlice...)`  

>Diff:
The resultent slices may have different `cap`, which is dependent on whether `append` expands for low-capacity

```go
// test deep copy
	cpv4 := append([]int(nil), v4...)
	info(v4)
	info(cpv4)
	fmt.Printf("%v \n", cpv4)

	cpv42 := make([]int, len(v4))
	copy(cpv42, v4)
	info(cpv42)
    fmt.Printf("%v \n", cpv42)
```
OUT:  
[1 2 4 5 6] 
0xc0000120a0: len=5, cap=10
# append 
0xc000078060: len=5, cap=6 
[1 2 4 5 6]
# make + copy 
0xc000078090: len=5, cap=5 
[1 2 4 5 6] 

[golang-map]: https://blog.golang.org/maps
