Scenario:
- foo.service 
 |- django service
 |- celery service

## Supervisord
-- **.conf --
```text
[program:foo_django]
command=/home/foo/bin/start_foo
user=foo
environment=LANG=en_US.UTF-8,LC_ALL=en_US.UTF-8

[program:foo_celery]
command=/home/foo/bin/start_celery
user=foo
environment=LANG=en_US.UTF-8,LC_ALL=en_US.UTF-8

[group:foo]
programs=foo_django,foo_celery
```

```sh
$ supervisorctl status foo:*
$ supervisorctl stop foo:*
$ supervisorctl start foo:*
$ supervisorctl restart foo:*
```

## `systemd`
`Unit.Wants`  if one of the `service` fails, it won't affect other services of the `target`  
`Unit.PartOf` `service.Unit.PartOf = target` means a service depend on target, when target starts/stops, its dependent also starts/stops  

-- /etc/systemd/system/foo.service --
```text
[Unit]
Description="Foo web application"
After=network.target
PartOf=foo.target

[Service]
User=foo
Group=foo
Environment=LANG=en_US.UTF-8,LC_ALL=en_US.UTF-8
ExecStart=/home/foo/bin/start_foo

[Install]
WantedBy=multi-user.target
```


-- /etc/systemd/system/foo.target --
```text
[Unit]
After=network.target
Wants=foo_django.service foo_celery.service

[Install]
WantedBy=multi-user.target
```

```sh
$ systemctl status foo
$ systemctl stop foo
$ systemctl start foo
$ systemctl restart foo
$ systemctl list-dependencies foo.target
```