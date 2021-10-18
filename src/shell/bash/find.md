# `find` Command
General usage [`man find`]:  
`find <PATH> [PREDICATE] [-print] [-exec -ok COMMAND] {} \;`  
1. -print print match wherever found  
1. -exec exec COMMAND when match found  
1. -ok ask confirmation
1. PREDICATE can be logically:
    * pred1 -and pred2 
    * ! pred
    * -not pred
    * pred1 -or pred2
    * pred1, pred2

## Predicates
1. type
1. name
1. size
1. time
1. ownership
1. permission

Predicate | Usages | Meaning
---|---|---
**type** | `-type <T>` | `f` file <br> `d` directory
**name** | `-name <REGEXP>` <br> `-iname` (case-insesitive) | `-name "*.py*"` <br> `-iname \*.py`
**size** | `-size <OPTION>` | `10c` : exact `10 chars` <br> `-10M` : `<10 mega bytes` <br> `+1K` : `>1 kilo bytes` 
**time** | `-atime 2` <br> `-ctime +2` <br> `-mtime -2` <br> `-amin 2` <br> `-cmin +2` <br> `-mmin -2` <br> `-anewer FILE` <br> `-newrct '2021-09-08 16:00:00'` | 1DAY < atime < 2DAYS <br> 2DAYS < ctime <br> mtime < 2DAYS <br> 1MIN < atime < 2MINs <br> 2MINs < ctime <br> mtime < 2MINs <br> atime newer than the file <br> ctime newer than the date
**ownership** | `-user <USER>` <br> `-gruop <GROUP>` | 
**permission** | `-perm <OPTION>` | `-g=r` groups with read <br> `/u=w,g=w` user with write, groups with write <br> `664` exact 664


[example]: https://www.huaweicloud.com/articles/55ce1495e24a24113128ae23e4276f74.html