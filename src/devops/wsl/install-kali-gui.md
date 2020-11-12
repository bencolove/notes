# Install Kali Linux with xfce GUI

## Install WSL Version 2
`dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart`  
`dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart`  

## Install Kali

## Install GUI
`$ sudo apt install kali-desktop-xfce -y`  
`$ sudo apt install xrdp -y`  
`$ sudo service xrdp start`  

`$ ip a`  
remote to it