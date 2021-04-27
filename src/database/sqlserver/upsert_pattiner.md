# Upsert Pattern
1. legacy upsert pattern
1. better upsert pattern with concurrency
1. merge

[[Sql server hints][sqlserver-hints]]

---

The legacy pattern would be regarded as anti-pattern:  
[[if-exist-update-insert]]

Check-exists-update-else-inesrt !!!
```sql
IF EXISTS (SELECT 1 FROM dbo.t WHERE [key] = @key)
BEGIN
  UPDATE dbo.t SET val = @val WHERE [key] = @key;
END
ELSE
BEGIN
  INSERT dbo.t([key], val) VALUES(@key, @val); 
END
```

## Better Upsert Pattern
For those most-likely inserts and updates:
```sql
-- optimized for most-likely inserts
BEGIN TRANSACTION;
 
INSERT dbo.t([key], val) 
  SELECT @key, @val
  WHERE NOT EXISTS
  (
    SELECT 1 FROM dbo.t WITH (UPDLOCK, SERIALIZABLE)
      WHERE [key] = @key
  );
 
IF @@ROWCOUNT = 0
BEGIN
  UPDATE dbo.t SET val = @val WHERE [key] = @key;
END
 
COMMIT TRANSACTION;
```

```sql
-- for updates are most likely
BEGIN TRANSACTION;
 
UPDATE dbo.t WITH (UPDLOCK, SERIALIZABLE) SET val = @val WHERE [key] = @key;
 
IF @@ROWCOUNT = 0
BEGIN
  INSERT dbo.t([key], val) VALUES(@key, @val);
END
 
COMMIT TRANSACTION;
```

The `UPDLOCK` and `SERIALIZABLE` table hints are there for concurrency:

LOCKs
Name | Explain | Purpose
---|---|---
`UPDLOCK` | Taken and held until the transaction completes | protect against conversion deadlocks at the statement level (let another session wait instead of encouraging a victim to retry).
`HOLDLOCK` <br> `SERIALIZABLE` | Makes shared locks more restrictive by holding them until a transaction is completed | protect against changes to the underlying data throughout the transaction (ensure a row that doesn't exist continues to not exist).


---

## Bulk Upsert
```sql
CREATE PROCEDURE dbo.UpsertTheThings
    @tvp dbo.TableType READONLY
AS
BEGIN
  SET NOCOUNT ON;
 
  BEGIN TRANSACTION;
 
  UPDATE t WITH (UPDLOCK, SERIALIZABLE) 
    SET val = tvp.val
  FROM dbo.t AS t
  INNER JOIN @tvp AS tvp
    ON t.[key] = tvp.[key];
 
  INSERT dbo.t([key], val)
    SELECT [key], val FROM @tvp AS tvp
    WHERE NOT EXISTS (SELECT 1 FROM dbo.t WHERE [key] = tvp.[key]);
 
  COMMIT TRANSACTION;
END
```

[if-exist-update-insert]: https://sqlperformance.com/2020/09/locking/upsert-anti-pattern
[sqlserver-hints]: https://docs.microsoft.com/en-us/sql/t-sql/queries/hints-transact-sql-table?view=sql-server-ver15