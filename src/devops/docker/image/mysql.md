# MySQL
[docker-mysql][docker-mysql]

## Image (Tags)
Tag | Image
---|---
`8.0.22` |
`5.7`

## Create Container
`docker run --name mysql8 -e MYSQL_ROOT_PASSWORD=<Passw0rd> -d -p 3306:3306 mysql:<tag> --default-authentication-plugin=mysql_native_password`

## Test Connection
`docker run -it --rm mysql8 mysql -hlocalhost -uroot -p`

## Create User/Change Password
`CREATE USER 'user'@'%' IDENTIFIED WITH mysql_native_password BY 'Passw0rd' ;`

`ALTER USER 'USER'@'%' IDENTIFIED WITH mysql_native_password BY 'Passw0rd' ; `

## Grant Permission
`GRANT ALL PRIVILEGES ON *.* to 'user'@'%' ;`

## Dump
1. `mysqldump`  
```bash
mysqldump -hlocalhost -uroot -p \
--default-character-set=utf8 \
--single-transaction \
--routines \
--events \
--trigger \
--all-databases | <database>
> 'dump.sql'
```

2. docker  
`docker exec mysql8 mysqldump > dump.sql`  or 
`sudo sh -c 'docker exec mysql8 mysqldump > dump.sql'`

## Restore
1. mysql  
`mysql -hlocalhost -uroot -p --default-character-set=utf8 --database=<target> < dump.sql`

2. docker  
`docker exec -i mysql8 mysql -uroot -p --database=<target> < dump.sql`



[docker-mysql]: https://hub.docker.com/_/mysql
