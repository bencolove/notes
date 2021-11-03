# Bash test
* `-O FILE`: test if FILE's owner is current effective USER
* `-t FD`: test if FD is open on terminal
    * `-t 0` test is _`stdin`_ is connected to terminal
* `[ -z "${var+x}"]` test variable `set` or not
    * if `var` not set: `"${var+x}"` set empty and `[ -z ]` true
    * otherwise `var` set(empty or not), `"${var+x}"` return `x` and `[ -z ]` false
* `[ -z "$var" ]` test variable `not set` or `empty`


### `test`
`test -e /folder_may_not_existed && echo 'exist' || echo 'not exist'`

> file type

FLAG | MEANING
---|---
-e | exist ?
-f | exist and is regular file ?
-d | exist and is folder ?

> integral

FLAG | MEANING
---|---
-eq | =
-ne | !=
-gt | >
-ge | >=
-lt | <
-le | <=

> string

FLAG | MEANING
---|---
test -z string | empty string
test -n string | not empty string
test str1 = str2 | string equal
test str1 != str2 | string not equal

> logical 

FLAG | MEANING | EXAMPLE
---|---|---
-a | AND | `[ TEST1 -a TEST2 ]` <br> `[ TEST1 ] && [ TEST2 ]` 
-o | OR | `[ TEST1 -o TEST2 ]` <br> `[ TEST1 ] || [ TEST2 ]`
! | negate

## `test` and `[]`
* [SP BOOL_EXP SP]
* "VAR"
* 'CONST'

## `[ exp ]` vs `[[ exp ]]`  

[comparison]: https://blog.csdn.net/x1269778817/article/details/46535729