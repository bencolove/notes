# Mysql Error Handling

## 1. Error Handler
[document][offical]
> **DECLARE** handler_action **HANDLER**  
> **FOR** condition_value [, condition_value ]...  
> statement
>
> handler_action := **CONTINUE** | **EXIT**  
>
> condition_value := mysql_error_code  
> | **SQLSTATE VALUE**     
> | condition_name  
> | **SQLWARNING**  
> | **NOT FOUND**  
> | **SQLEXCEPTION**


### 1.1 handler_action
value | meaning
---|---
CONTINUE|carry on execution in the current block
EXIT|exit current block

> The current block refers to the **BEGIN**...**END** block in which the handler is declared


### 1.2 condition_value

#### 1.2.1 mysql_error_code
An integer literal such as **1051** specify "unknown table"
```sql
DECLARE CONTINUE HANDLER FOR 1051
    BEGIN
        -- staments
    END;
```
> error code **0** means success

#### 1.2.2 SQLSTATE
A 5-character string literal equivalent to error code such as '42S01' specify "unknow table"
```sql
DECLARE CONTINUE HANDLER FOR SQLSTATE '42S02'
    BEGIN
        -- staments
    END;
```
> SQLSTATE starting with **00** means success

Special classes of SQLSTATE include 
1. **SQLWARNING** that begin with **'01'**
1. **NOT FOUND** that begin with **'02'**
1. **SQLEXCEPTION** that do not begin with **'00','01','02'** 

#### 1.2.3 conditino_name
User previously specified associating with either a error_code or SQLSTATE
```sql
DECLARE condition_name CONDITION FOR 
    <error_code> | SQLSTATE <value>
```
Some common codes [mysql error code][mysql_error_code_list]
error_code | SQLSTATE | meaning | message template
---|---|---|---
1050 | 42S01 | ER_TABLE_EXITS_ERROR | table %s already exists
1051 | 42S02 | ER_BAD_TABLE_ERROR | unknow table %s
1146 | 42S02 | ER_NO_SUCH_TABLE | table %s not exists
1022 | 23000 | ER_DUP_KEY | duplicate key in table %s
1048 | 23000 | ER_BAD_NULL_ERROR | Column %s cannot be null
1052 | 23000 | ER_NON_UNIQ_ERROR | Column %s in %s in ambiguous
1062 | 23000 | ER_DUP_ENTRY | Duplicate enry %s for key %s
1069 | 23000 | ER_DUP_UNIQUE | unique constraint to table %s
1216 | 23000 | ER_NO_REFERENCERD_ROW | cannot foreign key constraint fails
1217 | 23000 | ER_ROW_IS_REFERENCED | cannot delete or update foreign key constraint fails 



## Handler Resolve
In the case when no handler is defined in the block, action will be taken based on the condition class
* **SQLEXCEPTION** terminates current stored program like an **EXIT** handler, and handles control to outer/caller program if any
* **SQLWARNING** continue execution like an **CONTINUE** handler
* **NOT FOUND** **CONTINUE** for normal cases and **EXIT** for SIGNAL/RESIGNAL 

---

## 2. DIAGNOSTICS
[GET DIAGNOSTICS statment][get_diagnostics]  
Diagnostic areas are like error trace stack in Java.


### 2.1 Diagnostic Area
There are two diagnostic areas predefined: 
1. **CURRENT** (default) 
2. **STACK**

Depending on whether in the scope of an error handler defined within the stored program:
1. when **NOT** in a handler,  only **CURRENT** area is prepared for last executed statment;
2. when **IN** a handler,
    * **CURRENT** diagnostic area is populated for the last executed statement;
    * **STATCK** area is popluated for the outter/parent/caller context right before entering the handler

### 2.2 Diagnostic Area Structure
A Diagnostic Area is structured like:
```
Statement information:
  row count
  ... other statement information items ...
Condition area list:
  Condition area 1:
    error code for condition 1
    error message for condition 1
    ... other condition information items ...
  Condition area 2:
    error code for condition 2:
    error message for condition 2
    ... other condition information items ...
  Condition area 3:
    error code for condition 3
    error message for condition 3
    ... other condition information items ...
```
*statement information* includes
1. **ROW_COUNT** number of conditions occured
1. **NUMBER** affected row counts 

*condition information* includes
1. **MYSQL_ERRNO** related error_code
1. **RETURNED_SQLSTATE** related SQLSTATE
1. **MESSAGE_TEXT** related error message

### 2.3 Retrieve Diagnostic Area Data
For access to statement information
```sql
GET [CURRENT | STACKED] DIAGNOSTICS 
@p1= NUMBER, 
@p2=ROW_COUNT;
```

For retrieving condition information
```sql
GET [CURRENT | STACKED] DIAGNOSTICS CONDITION 1
@p3=MYSQL_ERRNO, 
@p4=RETURNED_SQLSTATE, 
@p5=MESSAGE_TEXT;
```

---

## 3. Signal Statement
[signal statment][signal_statement]  
SIGNAL is the way to return/throw/raise an error.

1. emit with **SQLSTATE**
```sql
-- inside a procedure
    DECLARE divide_by_zero CONDITION FOR SQLSTATE '22012';
    IF @case = 'with named condition' THEN
        SIGNAL divide_by_zero;
    ELSEIF @case = 'without error code' THEN 
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'An error occured';
    ELSEIF @case = 'with error code' THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'An error occured',
                MYSQL_ERRNO = 1001;
    END IF;
    
```

### Signal Condition Information Items
The following condition information items can be set when SINGAL or RESIGNAL:

Item Name | Difinition
---|---
CLASS_ORIGIN          | VARCHAR(64)
SUBCLASS_ORIGIN       | VARCHAR(64)
CONSTRAINT_CATALOG    | VARCHAR(64)
CONSTRAINT_SCHEMA     | VARCHAR(64)
CONSTRAINT_NAME       | VARCHAR(64)
CATALOG_NAME          | VARCHAR(64)
SCHEMA_NAME           | VARCHAR(64)
TABLE_NAME            | VARCHAR(64)
COLUMN_NAME           | VARCHAR(64)
CURSOR_NAME           | VARCHAR(64)
MESSAGE_TEXT          | VARCHAR(128)
MYSQL_ERRNO           | SMALLINT UNSIGNED


SQLSTATE Class Pattern | Meaning
---|---
'00xxx' | success
'01xxx' | waning
'02xxx' | not found
!= '02' | exception
'40xxx' | user-defined exception






[offical]: https://dev.mysql.com/doc/refman/8.0/en/declare-handler.html
[mysql_error_code_list]: https://dev.mysql.com/doc/refman/8.0/en/server-error-reference.html
[get_diagnostics]: https://dev.mysql.com/doc/refman/8.0/en/get-diagnostics.html
[signal_statement]: https://dev.mysql.com/doc/refman/8.0/en/signal.html



