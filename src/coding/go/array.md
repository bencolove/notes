# Array and Slice

`array` is underneath data structure dont directly dealt with.

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