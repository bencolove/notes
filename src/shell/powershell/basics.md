# Powershell Basics
Obeject output oriented

> Sources  
_Object_ or list of _objects_

> select, filter, map  
* `Select` _prop_list_
    * `select`, `select-object`, `select -property`
    * `select existingProps,@{'name'='extendedProp';'expression'={$_ ...}}`
* `Where` object
    * Where -Property _prop_ -GE -Value _value_ 
    * Where _prop_ -GE _value_
    * Where { $_.property -ge "value" } # `$_` refers to current object
    * Where hasProperty -eq $False
    * Where {!$_.hasProperty}
    * Where {($_.Name -Match "likethis*" -and $_.Name -notlike "that*") -and $_.hasProperty}
* `foreach ($obj in $source) { ... }`
> sort  
`Sort-Object prop`

> format and conversions 
`get-command -verb format -module microsoft.powershell.utlity`  

CMMD TYPE | NAME | EFFECT
---|---|---
Function | format-hex
cmdlet | format-custom
cmdlet | format-list | `prop_name: value` line-by-line
cmdlet | format-table | hashtab default format <br> (`-autosize`) <br> `-groupby` _prop_
cmdlet | format-wide | single property default to `name` 


To FORMAT | Cmdlet
---|---
CSV | Export-CSV _filepath_
HTML | ConvertTo-HTML -Property _prop_list_ > _filepath_
JSON | ConvertTo-JSON <br> -InputObject _object_or_piped_ <br> -Compress <br> -AsArry <br> -EnumsAsStrings

> redirection  
* `Out-file [-FilePath]` _filepath_ `[[-Encoding]]` {utf8}
* `>` _`filepath`_ 

## Cmdlet
* Get
* Set
* Start
* Stop
* Out
* New

## interpolation
`"""{0}"": {1}" -f "a", 2`
