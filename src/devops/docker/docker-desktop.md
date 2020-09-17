# About `docker-desktop`

### Q: How to move docker-desktop data to another location.  
A: [step-by-step][thread]

>Background:  
`wsl -l -v` should list two docker-desktop related running wsl-vms:
* docker-desktop
* docker-desktop-data

They should be by default located : `%USERPROFILE%\AppData\Local\Docker\wsl\data`

>What are the steps:
1. Shutdown wsl-vms
2. Export `docker-desktop-data`
3. Unregister from `wsl`
4. Import again at desired location


>How-to Details
1. Shutdown wsl-vms
	1.1 on [docker-desktop icon] -> [quit docker desktop]
	1.2 `wsl -l -v` to ensure they are `Stopped`

2. Export `docker-desktop-data` only
	`wsl --export docker-desktop-data "E:\vm\wsl\distro\docker-desktop-data.tar"`
	
3. Unregister from `wsl`
	`wsl --unregister docker-desktop-data`
	make sure by `wsl -l -v`
	
4. Import into desired location `E:\docker-desktop-data`
	`wsl --import docker-desktop-data "E:\docker-desktop-data" "E:\vm\wsl\distro\docker-desktop-data.tar" --version 2`




[thread]: https://stackoverflow.com/questions/62441307/how-can-i-change-the-location-of-docker-images-when-using-wsl2-with-windows-10-h