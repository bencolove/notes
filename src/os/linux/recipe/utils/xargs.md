# `xargs`
`xargs`   
1. reads items from **`STDIN`** separated by _SPACE_ and _LINE-FEED_  
2. convert them into Space-Separated-Values  
3. feed result items to following command
4. `xargs` by default is followed by `echo`

```bash
$ cat testxargs
a b c d
e f g h i
j k l m n o
p q r s t u v
w
x y z
# extract items from STDIN 
# separated by SPACE ' ' and NEWLINE '\n'
# output Space-Separated-Values results
$ cat testxargs | xargs
a b c d e f g h i j k l m n o p q r s t u v w x y z

# output one result every TWO lines from STDIN
$ cat testxargs | xargs -l2
a b c d e f g h i
j k l m n o p q r s t u v
w x y z
# output one result every 3 items from SOURCE(STDIN)
$ cat testxargs | xargs -n3
a b c
d e f
g h i
j k l
m n o
p q r
s t u
v w x
y z
# it seems -l -n will override each other
# and the last one takes controll
$ cat testxargs | xargs -l2 -n3
a b c
d e f
g h i
j k l
m n o
p q r
s t u
v w x
y z
$ cat testxargs | xargs -n3 -l2
a b c d e f g h i
j k l m n o p q r s t u v
w x y z
```

## Extract by delimitors
`xargs -d'\n' STDIN`  
`xargs -d'\0' STDIN`  
`xargs -0 STDIN`  

## Feed Specific Number of Arguments at a Time (batch)
```bash
$ echo {1..9}
1 2 3 4 5 6 7 8 9
$ echo {1..9} | xargs -n4 echo 0
0 1 2 3 4
0 5 6 7 8
0 9
```

## Feed Same Argument Multiple Times
Sometimes, a feeded item from `xargs` to the following command can make use the same item more than once by:  
`echo blarblarfile | xargs -I FILE echo FILE; ls FILE; rm FILE`  
`echo blarblarfile | xargs sh -c "echo $1; rm $1" sh`  

## Parallel
`xargs -p 0` : run at most n batches at once. _0_ means as many as possible. 

## Safe Mode
`xargs -P` : prompt for confirm before each run  
`xargs -t` : print and run


