# Configure Newwork with `netplan`
[Example][example]

Network Type | Guest -> Internet |Host -> Guest | Guest -> Host
---|:---:|:---:|:---:
NAT | O | X | ?
Host-only | ? | O | ?

## On VirtualBox VM
On fresh VB server(20.04 used here), NAT is enabled by default.

Simplest example for enable a Host-only network that host is accessible to the guest:
1. Add an `Host-Only Network Adapter`  
    1. [VM] -> [Settings] -> [Network]
    1. [Adapter] -> select [Host-Only Adapter] -> select one
    1. If no options, create one from [Tools] -> [Create] -> [Adapter]
    ![vb-adapter-manager][vb-adapter-manager]

1. Find interface name with `ip a`

1. Enable it with `netplan`

_`/etc/netplan/config.yaml`_
```yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    # NAT
    enp0s3:
      dhcp4: true
    # Host-only
    enp0s8:
      dhcp4: true
```



[example]: https://netplan.io/examples/
[vb-adapter-manager]: ../../../devops/vm/virtualbox/img/vb-adapter-manager.png

