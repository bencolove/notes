# `Get-NetTCPConnection`
`Get-NetTCPConnection`  
Filer parameters:
1. -State Listen
1. -LocalPort
1. -RemotePort
1. -OwningProcess

Select properties:
1. owningProcess(PID)

## Release PORT by killing process
```sh
# powershell
Get-NetTCPConnection -LocalPort 9090 -State listen -ErrorAction SilentlyContinue 
| Stop-Process -Force 

# linux
lsof -t -i :9090 | { read pid; ps -p $pid opid,ppid,cmd; }
lsof -ti:9090 | xargs ps opid= -p | xargs kill -9

