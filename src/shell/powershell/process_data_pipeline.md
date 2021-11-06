# Process Data
Basically powershell outputs results as _`object`_ or a list of _`object`_s. _`bash`_ in contrast outputs results as list of _`TEXT`_(lines)

> _`bash`_ Linux  
COMMAND -> lines

> _`powershell`_ Windows  
COMMAND -> objects
COMMAND -> objects -> `Out-String -Stream` -> lines  

`get-alias | out-string -stream | select-string 'get-com'` >>>  
Alias gcm -> Get-Command  
Alias gin -> Get-ComputerInfo  

`get-alias | select -exp displayname | sls 'get-com'` >>>
same as above

