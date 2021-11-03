# grammars
* boolean expression
* control flow
* debug script

## debug script
* -n grammar check
* -v `cat` before execute
* -x run with script content (useful)

## the Boolean condition
**`NO`** Boolean values !!!
What can do is to compare:
* 0 or 1 `[ $int -eq 0 ]`
* strings ("false", "true") `[ $str = 'False' ]`

## boolean exp for test

## flow controll

## conditional flow

### if-else
```sh
if [ BOOLEXP ]; then
  CMD;CMD;
elif [ BOOLEXP ]
then
  CMD;CMD;
else
fi
```

### case (switch)
```sh
case $var in
  "v1")
    CMD
    CMD
    ;;
  "v2")
    CMD;CMD;
    ;;
  "v3")
    CMD;CMD;
    ;;
  *)
    DEFAULT_CMD
    ;;
esac
```

### while-do-done
```shell
while [ BOOLEXP ]
do
  ...
done

# infinite-while-loop
while :
do
  ...
done

# line-by-line
while -r read line
do
  ...
done < FILE

# one-liner
while BOOLEXP; do CMD;CMD; done
```

### for-do-done
```sh
for $var in v1 v2 v3
do
  CMD
done

for (( i=1; i<$sum; i=i+1 ))
do
  CMD
done
```

[tricks]: https://codertw.com/%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80/544464/