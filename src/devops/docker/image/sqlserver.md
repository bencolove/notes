# Docker Images of Databases

## SqlServer
[microsfot-sqlserver][mcr-mssql]
[docker-mssql][docker-mssql]

### 1.Create container
```bash
$ docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=<Str0ngPassw0rd>' -p 1433:1433 --name sqlserver -h hostname -d mcr.microsoft.com/mssql/server:2019-latest
```
### 2.Images (Tags)
Tag | Cmd
---|---
`2017-latest` | docker pull mcr.microsoft.com/mssql/server:2017-latest
`2019-latest` | docker pull mcr.microsoft.com/mssql/server:2019-latest
`2017-CU21-ubuntu-16.04` | docker pull mcr.microsoft.com/mssql/server:<TAG>
`2019-CU8-ubuntu-16.04` | Same as above

### 3.Test connection
`docker exec -it <container_id|container_name> /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P <your_password>`

### 4.Change password
Otherwise `ps -eax` will show $SA_PASSWORD.
```bash
$ docker exec -it sqlserver /opt/mssql-tools/bin/sqlcmd \
   -S localhost -U SA -P "<YourStrong@Passw0rd>" \
   -Q 'ALTER LOGIN SA WITH PASSWORD="<YourNewStrong@Passw0rd>"'
```

### 5.`sqlcmd`
> from within container
1. `docker exec -it sqlserver bash`  
`/opt/mssql-tools/bin/sqlcmd -S localhost,port -Usa -P<password>` or
1. `docker exec -it sqlserver /opt/mssql-tools/bin/sqlcmd -S localhost,port -Usa -P<password>`

> from outside container  
`sqlcmd -S localhost,port -Usa -P<password>`

> Query  
`select Name from sys.databases `  
`go`

[mcr-mssql]: https://docs.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker?view=sql-server-ver15&pivots=cs1-bash
[docker-mssql]: https://hub.docker.com/_/microsoft-mssql-server
