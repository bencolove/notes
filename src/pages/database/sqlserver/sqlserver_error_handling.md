# SqlServer Error Handing
[Try..Catch statment][SqlServer_try_catch]
## Syntax
```sql
BEGIN TRY
    /* T-SQL Statments */
    /* Throw Error */
END TRY
BEGIN CATCH
    /* Print Error */
    /* Rollback Transaction */
END CATCH
```
---
## Types of Error
1. System defined
1. User defined

### User Defined Error
```sql
Throw 60000, N'Number is invalid', 5
```
Synposis  
`Throw <Error-Number>,<Message>, <State>` 
`RAISERROR(<Message>, <serverity-level>, <state>)`  

## Info about Error
Func or Global Var | datatype | meaning
---|---|---
@@ERROR | int |error number of last executed statement or 0 if success
ERROR_NUMBER() | int | 
ERROR_STATE() | int |
ERROR_MESSAGE() | nvarchar(4000) |
ERROR_SEVERITY() | int |
ERROR_LINE() | int |
ERROR_PROCEDURE() | nvarchar(128) | name of sp or trigger which causes the error

Actions can be used within `TRY-CATCH` block:
function | meaning
---|---
RAISEERROR() |
GOTO() |

### Error Serverity
Serverity Level | Meaning
---|---
13 | transaction deadlock errors
14 | security errors like permission denied
15 | syntax errors
16 | general errors can be fixed by users

## Error Handiling in Transaction Management
```sql
Begin Transaction Trans  
Begin Try  
    Delete From Employee Where Employee.Emp_IID<3  
    Update Employee Set Employee.First_Name='Pankaj kumar' Where Employee.Emp_IID='6th' /* Error Will Occur Here */  
    If @@TranCount>0  
    begin Commit Transaction Trans  
    End  
End Try  
Begin Catch  
    if @@TranCount>0  
        Print 'Error Is Occur in Transaction'  
        begin Rollback Transaction Trans  
    End  
End Catch
```

## Sophisticated Error Handling with XACT
```sql
create procedure [usp_my_procedure_name]
as
BEGIN
    -- abort any partially completed state transaction
    SET XACT_ABORT ON;
    -- not sending back rowcount meta data
    SET NOCOUNT ON;

    -- ensure unique save point name if you care
    DECLARE @mark CHAR(32) = replace(newid(), '-', '');
    DECLARE @trancount INT = @@TRANCOUNT;
    BEGIN TRY
        -- @@TRANCOUNT=0 means no outter transaction
        -- So begin one which can be rollbacked
        IF @trancount = 0
            BEGIN TRANSACTION @mark;
        ELSE
            -- in case of nested transaction
            -- only rollback to this save point
            SAVE TRANSACTION @mark; -- ( or use usp_my_procedure_name)

        -- Do the actual work here

lbexit:
        -- no outer transaction, simply commit
        IF @trancount = 0   
            COMMIT TRANSACTION @mark;
    END TRY
    BEGIN CATCH
        DECLARE @error INT, @message NVARCHAR(4000), @xstate INT;
        select @error = ERROR_NUMBER(), 
            @message = ERROR_MESSAGE(), 
            @xstate = XACT_STATE();
        -- uncommitable state, should definitely rollback
        IF @xstate = -1
            ROLLBACK;
        -- still rollback whole batch for no nested transaction
        IF @xstate = 1 and @trancount = 0
            ROLLBACK
        -- rollback to save point in nested transaction
        IF @xstate = 1 and @trancount > 0
            ROLLBACK TRANSACTION @mark -- usp_my_procedure_name;
        -- pass controll back to outter namespace
        RAISERROR('usp_my_procedure_name: %d: %s', 16, 1, @error, @message) ;
        -- OR THROW
    END CATCH   
END
go
```
This can be used as a template:
```sql
create procedure [usp_my_procedure_name]
as
BEGIN
    -- abort any partially completed state transaction
    SET XACT_ABORT ON;
    -- not sending back rowcount meta data
    SET NOCOUNT ON;

    DECLARE @mark CHAR(32) = replace(newid(), '-', '');
    DECLARE @trans INT = @@TRANCOUNT;

    IF @trans = 0
        BEGIN TRANSACTION @mark;
    ELSE
        SAVE TRANSACTION @mark;

    BEGIN TRY
        -- do work here

        IF @trans = 0
            COMMIT TRANSACTION @mark;
    END TRY
    BEGIN CATCH
        IF xact_state() = 1 OR (@trans = 0 AND xact_state() <> 0) 
            ROLLBACK TRANSACTION @mark;
        THROW;
    END CATCH
END
GO
```



## XACT_ABORT

## XACT_STATE()
Test XACT_STATE:  
1. If 1, the transaction is committable.  
1. If -1, the transaction is uncommittable and should be rolled back.  
1. If XACT_STATE = 0 means that there is no transaction and 
a commit or rollback operation would generate an error.

Utilize `@@TRANCOUNT` to test whether a nested transaction is currently applied. In a nested transaction context, only current amount of operations should be rollbacked (using **SAVE_POINT**). The whole would be rollbacked if otherwise.


[SqlServer_try_catch]: https://docs.microsoft.com/en-us/sql/t-sql/language-elements/try-catch-transact-sql?view=sql-server-ver15