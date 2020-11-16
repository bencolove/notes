# Docker Network Driver

Driver | Usecase | tutorial
---|---|---
`bridge` | **User-defined bridge** for containers on same Docker host | [bridge]
`host` | share host's networking | [host]
`overlay` | for containers on different Docker hosts | [overlay]
`macvlan` | for direct attach like physical with unique MAC address | [macvlan]
`none` | isolated | 
other | third-party |

[docker-iptables]: https://docs.docker.com/network/iptables/
[bridge]: https://docs.docker.com/network/bridge/

[stand-alone]: https://docs.docker.com/network/network-tutorial-standalone/
[host]: https://docs.docker.com/network/network-tutorial-host/
[overlay]: https://docs.docker.com/network/network-tutorial-overlay/
[macvlan]: https://docs.docker.com/network/network-tutorial-macvlan/
