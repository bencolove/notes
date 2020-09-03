# Use `Multipass` for Development and Testing
Senarios:
* Create
* List
* Info
* Quick REPL
* Shell in
* Share Data
* Delete

## Create Fresh VM
```bash
$ multipass launch -n <vm-name>

$ multipass launch -c 1 -m 2G -d 5G -n <vm-name> 20.04

```
* -n vm-name
* -c cpu
* -m memory
* -d disk
* 20.04 version

For versions, they can be found by  
`$ multipass find`
Often used ones:
* 16.04
* 18.04
* 20.04
* core
* core18

---
## List VMs
`$ multipass list`  

---
## Info of VMs
`$ multipass info <vm-name>`  

---
## Quick REPL
`$ multipass exec <vm-name> <command>`

---

## Shell in VMs
`$ multipass shell <vm-name>`

---
## Share Data
Data can be shared between host and guest by
* share directories
* transfer files

```bash
$ multipass mount <dir-to-share> <vm-name>
$ multipass info <vm-name>
$ multipass shell <vm-name>
$ multipass unmount <vm-name>

$ multipass transfer <vm-name>:source <vm-name>:destination
```

---
## Delete VMs
```bash
$ multipass stop <vm>
$ multipass delete <vm>
$ multipass purge
```
>NOTE:  
>  Delete an VM will not remove it instead only mark it as 'DELETED'.
> The 'DELETED' vm will be removed completed by **purge**.

