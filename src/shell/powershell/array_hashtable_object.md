# array, hashtable and psobject

## array

## hashtable
`hashtable` is like `dict` with Python, `map` with Java and `map[string]interface{}` with GO.

```powershell
$ht=@{}
$ht=@{
    key1 = $v1
    key2 = $v2
}

$ht.add($key, $value)

# Len(map)
$ht.count

# dict.keys
$ht.keys
# dict.values
$ht.values

# test key exist
# because $false and numeric value ZERO are evaluated to false in logical expression
if ( $ht.notexits -eq $null ) { }
# DO THIS
if ( $ht.containskey($key) ) { }

# remove key
$ht.remove($key)

# clear
$ht.clear()

# ordered
[ordered]@{}

<# computed property used in 
    `select` and `format-table` 
#>
get-psdrive | where used | select name, @{
    name='totalGB'
    expression={($_.used + $_.free) / 1GB}
}
# or
select name, @{n='totalGB';e={}}

# as key word arguments
# destructure
$DHCPScope = @{
    Name        = 'TestNetwork'
    StartRange  = '10.0.0.2'
    EndRange    = '10.0.0.254'
    SubnetMask  = '255.255.255.0'
    Description = 'Network for testlab A'
    LeaseDuration = (New-TimeSpan -Days 8)
    Type = "Both"
}
Add-DhcpServerv4Scope @DHCPScope
```


### hashtable for-loop
```powershell
foreach($key in $ht.keys)
{
    write-output "key:{0}, value:{1}" -f $key, $ht[$key]
}

$ht.keys | % {
    write-output "key:{0}, value:{1}" -f $_, $ht[$_]
}

$ht.GetEnumerator() | % { write-output "key:{0}, value:{1}" -f $_.key, $ht[$_.value] }

# modify while loop
$ht.keys.clone() | % { $ht[$_] += 1 }
```

## sort
```powershell
# list.sort(key=name)
get-psdrive | sort name

# list.sort(key=lambda x: x.name)
get-psdrive | sort @{e={$_.name}}
```

## conversion
> hashtable -> pscustomobject  
`[pscustomobject]$hashtable`  

## saving from/to file
> hashtable -> pscustomobject -> csv  
`$ht | % { [pscustomobject]$_ } | export-csv $filepath`

> nested hashtable -> json -> file  
`$ht | convertto-json | set-content $filepath`  
`$ht=get-content $filepath -raw | convertfrom-json`

## pscustomobject
`[pscustomobject]@{p1="v1";p2=2}`

## pscustomobject <=> hashtable
source=`Test-NetConnection "www.google.com.tw" -Port 443`
filter=`$source | select ComputerName, RemoteAddress, TcpTestSucceeded`  
obj=`$filter.psobject.Properties | % {$ht=@{}} {$ht[$_.name]=$_.value} {$ht}`  
>$obj
Name                         Value                                                                  
----                         -----                                                                  
ComputerName                 www.google.com.tw                                                      
TcpTestSucceeded             True                                                                   
RemoteAddress                172.217.160.99


ht=`New-Object PSObject -property $ht`
>$ht
ComputerName      TcpTestSucceeded RemoteAddress 
------------      ---------------- ------------- 
www.google.com.tw             True 172.217.160.99