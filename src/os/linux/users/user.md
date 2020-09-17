# Manage `user`

## 1. Create user
```bash
# create a user with all defaults
$ sudo adduser <user>
# add user to sudo group
$ sudo usermod -aG sudo <user>
# test 
```

## 3. Delete user
`$ sudo deluser --remove-home <user>`

## 4. Config `ssh`
--_`/etc/ssh/sshd_config`_--
```text
...
PasswordAuthentication no -> yes
...
```
>Restart `ssh`
```bash
$ sudo systemctl restart ssh
```