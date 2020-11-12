# Create Linked Server
```sql
-- create linked database
EXEC [master].[dbo].sp_addlinkedserver
	@server     = N'azure',
    @datasrc    = N'bt04server.database.windows.net',
    @srvproduct = N'Any',
    @provider   = N'SQLNCLI',
    @catalog    = N'italenttrial'
GO
-- create login
EXEC sp_addlinkedsrvlogin 
    @rmtsrvname = N'azure',   
    @useself = 'FALSE', 
    @locallogin=NULL,
    @rmtuser = 'sa',
    @rmtpassword = 'xxx'
GO
-- set [rpc][rpc-out] to true
EXEC master.dbo.sp_serveroption @server=N'azure', @optname=N'rpc', @optvalue=N'true'
GO

EXEC master.dbo.sp_serveroption @server=N'azure', @optname=N'rpc out', @optvalue=N'true'
GO

select *
from azure.db_name.dbo.table_name
```