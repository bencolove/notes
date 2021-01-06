# Basic Data Types
* array
* slice
* map

## Declaration

DataType | Declare | Value | Size(amd64) | Comment
---|---|---|---|---
array | `var [3]int` | `[0 0 0]` | 3 words | allocated on stack
slice | `var []int` | **`nil`** | 24 bytes <br> 3 words | reference type
pointer | `var *int` | **`nil`** | 8 bytes <br> 1 word | reference type
map | `var map[string]int` | **`nil`** | 8 bytes <br> 1 word | reference type


```go
var varSlice []int
fmt.Printf("var varSlice []int\nvarSlice==nil: %t\nlen(varSlice)=%d\nsizeOf(varSlice)=%d\n\n", varSlice==nil, len(varSlice), unsafe.Sizeof(varSlice))

makeSlice := make([]int,10)
fmt.Printf("makeSlice := make([]int,10)\nmakeSlice==nil: %t\nlen(makeSlice)=%d\nsizeOf(makeSlice)=%d\n\n", makeSlice ==nil, len(makeSlice ), unsafe.Sizeof(makeSlice))

newSlice := new([]int)
fmt.Printf("newSlice := new([]int)\nnewSlice==nil: %t\n*newSlice==nil: %t\nlen(*newSlice)=%d\nsizeOf(newSlice)=%d\n\n", newSlice==nil, *newSlice==nil, len(*newSlice), unsafe.Sizeof(newSlice))


var varMap map[string]int
fmt.Printf("var varMap map[string]int\nvarMap==nil: %t\nlen(varMap)=%d\nsizeOf(varMap)=%d\n\n", varMap ==nil, len(varMap), unsafe.Sizeof(varMap))

makeMap := make(map[string]int)
fmt.Printf("makeMap := make(map[string]int)\nmakeMap==nil: %t\nlen(makeMap)=%d\nsizeOf(makeMap)=%d\n\n", makeMap==nil, len(makeMap), unsafe.Sizeof(makeMap))

newMap := new(map[string]int)
fmt.Printf("newMap := new(map[string]int)\nnewMap==nil: %t\n*newMap==nil: %t\nlen(newMap)=%d\nsizeOf(*newMap)=%d\n\n", newMap==nil, *newMap==nil, len(*newMap), unsafe.Sizeof(newMap))

/* Outputs
var varSlice []int
varSlice==nil: true
len(varSlice)=0
sizeOf(varSlice)=24

makeSlice := make([]int,10)
makeSlice==nil: false
len(makeSlice)=10
sizeOf(makeSlice)=24

newSlice := new([]int)
newSlice==nil: false
*newSlice==nil: true
len(*newSlice)=0
sizeOf(newSlice)=8

var varMap map[string]int
varMap==nil: true
len(varMap)=0
sizeOf(varMap)=8

makeMap := make(map[string]int)
makeMap==nil: false
len(makeMap)=0
sizeOf(makeMap)=8

newMap := new(map[string]int)
newMap==nil: false
*newMap==nil: true
len(newMap)=0
sizeOf(*newMap)=8
*/
``` 

## Assignment and Copy

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