# `diff` and Result Insights

```bash
$ diff -u <(printf "%d\n" 1 3 4 5) <(printf "%d\n" 2 3 5 6)
--- /dev/fd/63  2020-09-04 14:10:46.204886800 +0800
+++ /dev/fd/62  2020-09-04 14:10:46.207815300 +0800
@@ -1,4 +1,4 @@
-1
+2
 3
-4
 5
+6
```
```bash
$ diff <(printf "%d\n" 1 3 4 5) <(printf "%d\n" 2 3 5 6)
1c1
< 1
---
> 2
3d2
< 4
4a4
> 6
```

When running `diff olfFile newFile` the command reads **Find diff from right against left**, the ordering is important for understanding the results.

Two output formats:
* -u unified format
* WITHOUT -u old good format

---
## Unified Format
>Input  
`$diff <(printf "%s\n" l1 3 l4 5) <(printf "%s\n" r2 3 5 r6)`
>Output
```
--- /dev/fd/63  2020-09-04 17:21:13.574656500 +0800
+++ /dev/fd/62  2020-09-04 17:21:13.577583200 +0800
@@ -1,4 +1,4 @@
-l1
+r2
 3
-l4
 5
+r6
```
It tells that
>From
``` 
1: l1 
2: 3 
3: l4 
4: 5
```
>To
```
1: r2 
2: 3 
3: 5 
4: r6
```
>You go through
```bash
--- /dev/fd/63  2020-09-04 17:21:13.574656500 +0800 # left file
+++ /dev/fd/62  2020-09-04 17:21:13.577583200 +0800 # right file
@@ -1,4 +1,4 @@ # differs over left(1,4) and right(1,4)
-l1 # left content
+r2 # right content
 3
-l4 # left content
 5
+r6 # right content
``` 

---
## Normal Format  
>Input  

`diff <(printf "%s\n" l1 3 l4 5) <(printf "%s\n" r2 3 5 r6)`  

>Output
```
1c1
< l1
---
> r2
3d2
< l4
4a4
> r6
```
It tells that
>From
``` 
1: l1 
2: 3 
3: l4 
4: 5
```
>To
```
1: r2 
2: 3 
3: 5 
4: r6
```
>You go through
```bash
1c1  # replace left(1:1) of left with right(1:1)
< l1 # left content < l1
---
> r2 # right content > r2
3d2  # delete left(3:3), content left file < l4
< l4
4a4  # add right(4:4) after line left(4), content right file  > r6
> r6
```

Pattern | Denote | Meaning | Left -> right | Right -> left
---|---|---|---|---
**LaR** | a (add, insert) | **`a`** dd <br> **`L`** ine <br> **`R`** ange | `8a12,15` <br> append right(12-15) **`AFTER`** left(8) | `8a12,15` <br> delete right(12-15)
**FcT** | c (change) | **`c`** hange <br> **`F`** rom <br> **`T`** o | `5,7c8,10` <br> swap left(5,7) and right(8,10) | `5,7c8,10` <br> swap right(8,10) and left(5,7)
**RdL** | d (delete, remove) | **`d`** elete <br> **`R`** ange <br> **`L`** ine | `10,12d7` <br> delete left(10, 12) | `10,12d7` <br> add left(10,12) **`AFTER`** right(7) 
