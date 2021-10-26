# Variables and Arrays

> scopes  
* local (default)
* global ($Global:)
* environment ($env:)

> declare  
`$var = value`  

> typed declare  
`[string]$strVal1 = "strValue"`  
`$strVal1.GetType().FullName => System.String`  

> types  

TYPE | BITS |EXAMPLE
---|---|---
[string] | | System.String <br> $s="abc"
[char] | | System.Char <br> [char]$v=0x265b
[bool] | | System.Boolean <br> $true or $false
[int] | 32 |System.Int32
[long] | 64 | System.Int64
[decimal] | 128 'd' | System.Decimal <br> $v=111.222d
[double] | 8 | System.Double <br> [double]$v=2222.1111
[single] | 32 | System.Single <br> [single]$v=11111.22
[DateTime] | | System.DateTime <br> $dt=get-date
[array] | | System.Object[] <br> $arr1="first", "second", "third" <br> $arr2=@() <br> $arr3=,1 <br> [int32[]]$arr4 ($null) <br> $arr4=New-Object System.Collections.ArrayList
[hashtable] | | System.Collections.Hashtable <br> $map=@{1="v1";2="v2";3="three"} <br> $map.Add("key", "val") <br> $map.key="v" <br> $map.Remove("key")

> clear  
`Clear-Variable -Name var`

> remove  
`Remove-Variable -Name var`

## Arrays

> Length  
`$arr.Length`

> Add  
`$arr.Add("val")`

> Remove  
`$arr.Remove("val")`

> Test contains  
`$arr.Contains("val")`

> Clear (set to $null)  
`$arr=$null`

> Check  
`$arr -eq $null`

> Print  
`$arr | Out-File filepath`  
`$arr | Export-CSV > filepath`  


> Loop
```powershell
$arr = @("a", "b", "c")

foreach ($o in $arr) {
    "$o.Length=$($o.Length)" 
}

# loop with index
% {$idx=0} {statements; $idx++}

# same
foreach ($idx=0; $idx -lt $array.Length; $idx++) {...}

# same
foreach ($item in $arry) { $arry.indexof($item) }

# general usage
foreach {init} {statements} {final}
1..10 | foreach {$total=1} {$total*= $_} {$total}
```


