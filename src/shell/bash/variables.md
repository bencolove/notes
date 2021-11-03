# Variables
* delcare variable
* set/unset
* `${}`

## list all defiend shell variables
`declare -p` list all variables  
`declare -xp` == `env` list all environmental variables
>OUTPUT:  
declare -ir PPID="1123"
declare -i BASHPID
declare -x PATH="paths"
declare -i RANDOM
declare -a PIPESTATUS
declare -- CUSTOMVAR

* -i integer value
* -r readonly
* -x export
* -a array
* -- custom variable

## Special Variables
VAR | MEANING
---|---
`$$` | current PID
`cat /proc/self/stat | awk '{ print $4 }'` | current PID
`$!` | last background PID
`$?` | last command return code
`$*` | "$1 $2 $3"
`$@` | "$1" "$2" "$3"
`$#` | num args
`$0` | script name
`BASHPID` | current PID
`PPID` | parent PID
`PWD` | current Working Directory
`PATH` |
`USER` |
`UID` |
`EUID` | effective UID, `[[ -O $file ]]` to test file's owner is current effective USER
`LOGNAME` |

## `${}`

> default value  
`var=${var:-default}`

> string variable substring and replace  

`str=12345` | `${}`
---|---  
${var:start} | ${str:1} 2345 <br> ${str:1:3} 234 <br> ${str: -2} 45 <br>
${var/pattern/replace} | ${str/2/a} 1a345 
${var//pattern/replace} | replace all pattern