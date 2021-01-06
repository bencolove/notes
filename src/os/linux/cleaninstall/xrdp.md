# _xrdp_

## install
```bash
apt update

apt install xrdp

apt-add-repository universe

apt install gnome-tweak-tool

apt search gnome-shell-extension

apt install gnome-shell-extension-ubuntu-dock

gnome-tweaks
```
## Issues
>Popups
 of 'Authentication required for color managed device' hangs  

```bash
cat << EOF > /etc/polkit-1/localauthority/50-local.d/45-allow-colord.pkla

[Allow Colord all Users]
Identity=unix-user:*
Action=org.freedesktop.color-manager.create-device;org.freedesktop.color-manager.create-profile;org.freedesktop.color-manager.delete-device;org.freedesktop.color-manager.delete-profile;org.freedesktop.color-manager.modify-device;org.freedesktop.color-manager.modify-profile
ResultAny=no
ResultInactive=no
ResultActive=yes

[Allow Package Management all Users]
Identity=unix-user:*
Action=org.debian.apt.*;io.snapcraft.*;org.freedesktop.packagekit.*;com.ubuntu.update-notifier.*
ResultAny=no
ResultInactive=no
ResultActive=yes
EOF
```
