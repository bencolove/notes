# Files

## Filter

```powershell
# search for files
# current folder (no -recursive)
# only *.txt
# lastwritetime later than '2021-10-25'
gci ./* -include *.txt | \
? { $_.lastwritetime -gt '2021-10-25 00:00:00' } | \
select -exp name


# current folder
# lastwritetime earlier than one day ago
gi ./* -exclude *.txt | \
? { $_.lastwritetime -lt $(get-date).adddays(-1) }

```

## Statistics
1. source
1. group
1. aggregate `select $_.group | measure`

```powershell
gci -recurse -file | 
group extension -noelement |
sort count -desc |
select -first 5 |
ft name, count
>>>
Name Count
.go  14
.log  2

gci -recurse -file |
group extension |
sort count -desc |
select -first 5 count, name, @{
    n='size';e={
        ($_.group | measure length -sum).sum
    }
}
>>> 
count name size
14    .go  42342
2     .log 12312312
```