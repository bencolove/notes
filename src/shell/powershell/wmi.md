# _`WMI`_ and _`CIM`_



---

_`WMI`_ cmdlets
1. Get-WmiObject `gwmi`
1. Invoke-WmiMethod `iwmi`
1. Register-WmiEvent 
1. Remove-WmiObject `rwmi`
1. set-WmiInstance `swmi`

> General Usage  
`gwmi -class WMI_CLASS -Filter 'PROP=VAL'`  
`gwmi -Query 'select * from WMI_CLASS where PROP=VAL'`  

## Classes
CLASS | RESOURCES
---|---
Win32_ComputerSystem | Get-Host
Win32_LogicalDisk | Get-PSDrive
Win32_Process | Get-Process
Win32_NetworkAdapterConfiguration | TCP/IP configuration
Win32_NetworkAdapter | hardware

### List Classes
`Get-WmiObject -List filter`

### List Class Method  and Parameters
> List class methods  
`([wmiclass]"win32_volumne").methods.name`

> List method parameters
`([wmiclass]"win32_volume").GetMethodParameters('Format')` 

## Get, Set and Invoke
> Get-set modify  
`gwmi -class win32_volume -filter "name='d:\\'" | swmi -arguments @{label="new_label"}`  

> get-call invoke  
`gwmi -class win32_process -filter "name='calc.exe'" | iwmi -name terminate`  
`iwmi -class win32_process -name create -argumentlist 'calc.exe', 'c:\script'`  

## Events
`Get-WmiObject -List *event*`  

---

## CIM

_`CIM`_ cmdlets

Group | Cmdlets | Alias |WMI Cmdlets
---|---|---|---
-- session --
New-CimSessionOption | `ncso` 
New-CimSession | `ncms`
Get-CimSession | `gcms`
Remove-CimSession | `rcms`
--  class --
Get-CimClass | `gcls`
-- objects --
Get-CimAssociatedInstance | `gcai` | Get-WmiObject
New-CimInstance | `ncim` | Get-WmiObject
Get-CimInstance | `gcim` | Get-WmiObject
Set-CimInstance | `scim` | Set-WmiInstance
Remove-CimInstance | `rcim` | Remove-WmiInstance
-- invoke --
Invoke-CimMethod | `icim` | Invoke-WmiMethod
-- events --
Register-CimIndicationEvent | `rcie` | Register-WmiEvent

### Insight _`class`_
> List _class_  
`Get-CimClass [-Class Win32_ComputerSystem]`  
* CimClassName
* CimClassMethods
* CimClassProperties

> Insight _methods_  
`$class.CimClassMethods["Rename"].Parameters`  

### List _`instance`_
> List filter  
`get-ciminstance -class win32_process -filter 'name="powershell.exe"'`  

> List query  
`get-ciminstance -query "select * from win32_process where name='powershell.exe'"`  

> Refresh properties  
1. `$p = get-ciminstance -class win32_process -filter 'name=$pid'`
1. `$p | ft kernelmodetime`
1. `$p | get-ciminstance | ft kernelmodetime`
It works like: `$p` caches the state(results) when executed and of type `ciminstance`. Afterwards, `$p` serves like an ID to the resource when `get-ciminstance` again.

> Nested classes (joined case)  
1. `$nic = get-ciminstance -class win32_networkadapter -filter "netenabled=$true AND netconnection like '%wi-fi%'"`
1. `get-ciminstance -ciminstance $nic -resultclassname win32_networkadapterconfiguration | select ipaddress`

> Remove  
`get-ciminstance -class win32_process -filter "name='calc.exe'" | remove-cimsintance`

> Modify  
`get-ciminstance -class win32_volumne -filter "name='c:\\'" | set-ciminstance -property @{label="testlabel"}`

### Invoke _`instance`_'s Method
Since `Get-CimInstance` returns a standard instance of cim _`class`_ with which you may not invoke directly its method with `.` notation like _`WMI`_ way, `Invoke-CimMethod` comes into play:  
`Invoke-CimMethod -class win32_process -methodname create -arguments @{commandline="calc.exe";currentdirectory="c:\scripts"}`  


## CIM Events
Workflow:
1. find event class
1. register
1. fetch events
1. unregister

### Find Event Class
`get-cimclass cim_inst*` or  
`get-cimclass *event | select -exp cimclassname`  

### Register Event
System clock events
1. `$q = "select * from cim_instmodification where targetinstance isa 'win32_localtime'"`
1. `register-cimindicationevent -query $q`

### Fetch Events
`Get-Event | select timegenerated`

### Unregister
`Unregister-Event *`  
`Get-Event | Remove-Event`

## CIM Session
Remote access with protocol `WS-MAN` by default unless specified `DCOM`.
1. -ComputerName for one-time operation
1. -CIMsession for a series of actions

> Create CIM session  
Remote `WinRm` service needed:
1. `$cim = New-CimSession -ComputerName webr201`
1. `$cim`
1. `Get-CimInstance -class win32_OperatingSystem -CimSession $cim`

> Test and Drop-back to DCOM  
1. `Test-WSMAN -ComputerName webr201`
1. `$cimopts = New-CimSessionOption -Protocol DCOM`
1. `$cimDcom = New-CimSession -ComputerName server02 -SessionOption $cimopts`
