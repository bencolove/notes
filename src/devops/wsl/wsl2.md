# WSL 2
* Upgrade/enable it on Win10
* Export distro
* Import distro
* `wsl`
* Location of distros
* [Resize distros][resize-distro]

---

## Install/Upgrade/Enable WSL2


---

## Export/Import Linux Distro
```powershell
# export
PS> wsl --export ubuntu-20.04 E:\vm\wsl\distro\ubuntu2004.tar
PS> wsl --export kali-linux E:\vm\wsl\distro\kali.tar

# import
PS> wsl --import ubu-go E:\vm\wsl\distro\ubuntu2004.tar
# set wsl2
PS> wsl --set-version ubu-go 2
# or in one go
PS> wsl --import ubu-go E:\vm\wsl\distro\ubuntu2004.tar --version 2
```

## `wsl`
```powershell
# list
PS> wsl -l -v
# stop a distro
PS> wsl -t <distro>
# delete a disto
PS> wsl --unregister <distro>

# stop all distors
PS> wsl shutdown

# run a distro
PS> wsl -d <distro> -u <user>
# set default distro
PS> wsl --set-default|-s <distro>

# set a distro WSL version
PS> wsl --set-version <distro> <version>
```

## Locations
```powershell
PS> Get-AppxPackage -Name "*<distro>*" | Select PackageFamilyName
PackageFamilyName
-----------------
CanonicalGroupLimited.Ubuntu18.04onWindows_79rhkp1fndgsc
CanonicalGroupLimited.Ubuntu20.04onWindows_79rhkp1fndgsc

# disto location
PS> explorer.exe %LOCALAPPDATA%\Packages\<PackageFamilyName>\LocalState\<disk>.vhdx
```

[resize-distro]: https://docs.microsoft.com/en-us/windows/wsl/compare-versions#expanding-the-size-of-your-wsl-2-virtual-hardware-disk