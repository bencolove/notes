# Byte-Order
* Big-endian
* Little-endian
* [example](https://blog.gtwang.org/programming/difference-between-big-endian-and-little-endian-implementation-in-c/)

## Difference
> Big-endian  
![be-img](https://blog.gtwang.org/wp-content/uploads/2018/05/big-endian-20180517-1-816x382.png)  

> Little-endian  
![le-img](https://blog.gtwang.org/wp-content/uploads/2018/05/little-endian-20180517-1.png)


## Determin Machine Byte-Order
Simply achive this by putting an 32bit(4byte) word like `0X12345678` into an variable(type `long` in `C` or `int32` in `GO`). And then read the last byte(positioned at 3) against `0x12`(Little-endian) or `0x78`(Big-endian)
```c
#include <stdio.h>
typedef union {
  unsigned long l;
  unsigned char c[4];
} EndianTest;
int main() {
  EndianTest et;
  et.l = 0x12345678;
  printf("本系統位元組順序為：");
  if (et.c[0] == 0x78 && et.c[1] == 0x56 && et.c[2] == 0x34 && et.c[3] == 0x12) {
    printf("Little Endiann");
  } else if (et.c[0] == 0x12 && et.c[1] == 0x34 && et.c[2] == 0x56 && et.c[3] == 0x78) {
    printf("Big Endiann");
  } else {
    printf("Unknown Endiann");
  }
  printf("0x%lX 在記憶體中的儲存順序：n", et.l);
  for (int i = 0; i < 4; i++) {
    printf("%p : 0x%02Xn", &et.c[i], et.c[i]);
  }
  return 0;
}
```

```go
// 0x12345678 左至右是從高位到地位
// 0x12是高位byte
// 0x78是地位byte
i32 := 0x12345678
// i32[0]是低位址
// check the first byte
// 0x78 little-endian 低位byte存在低位位址
// 0x12 big-endian 高位byte存在低位位址
ptri32 := (*byte)(unsafe.Pointer(&i32))
fmt.Printf("i32&=%p *byte=%p \n", &i32, ptri32)

// the first(0th) byte address
ptrFirstByte := (*byte)((unsafe.Pointer)(0 + uintptr(unsafe.Pointer(&i32))))
firstByte := *ptrFirstByte
fmt.Printf("last=%x @%p\n", firstByte, ptrFirstByte)

if firstByte == 0x12 {
    fmt.Println("big-endian")
} else if firstByte == 0x78 {
    fmt.Println("little-endian")
} else {
    fmt.Println("unknown")
}
```


<style>
img[alt$=-img] {
   width:400px;
   height:200px;
   background-color: white
}
</style>