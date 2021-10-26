# Powershell.exe

## help
`get-help -name cmdlet -online`

## list-member
`object_unknown | get-member`

## invoke powershell
`powershell.exe -c '...'`  
`powershell.exe -f file.ps1`  

## performance
Performance consideration:
1. objects first
    `(get-host).version.tostring()`  
    `(0..10).foreach({$_})`  
    `format-table dayofweek -inputobject (get-date)`
2. pipe
    `get-host | . { proces { $_.version.tostring() }}`  
    `0..10 | foreach { $_ }`  
    `get-date | format-table dayofweek`

> General way  
Cmdlet1 | Cmdlet2  
to:  
Cmdlet2 -inputobject (Cmdlet1)

## predifined variables
VAR | MEANING
---|---
`$?` | `bool` indicatin last comand successfully or not
`$args` | passed-in arguments as `array`

## named parameter
-_script.ps1_-
```powershell
param (
    [Parameter(Mandatory)]
    [string]$strValue = "default",

    [int]$intValue
)
```
> Call it with
`powershell.exe -f script.ps1 -intValue 1 -strValue or_not_given`