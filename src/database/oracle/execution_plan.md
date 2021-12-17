# Execution Plan
1. estimated (without executing)
1. actual (executed)

Where are they [[stored][diff]]:
1. estimated  
temp table `table_plan` after _`explain plan for`_
1. actual  
_`v$`_ structures, may be _`v$cursor`_ 
---
## Estimated Plan
[[explain plan for](https://use-the-index-luke.com/sql/explain-plan/oracle/getting-an-execution-plan)] 
1) Simply preceding your SQL with `explain plan for` to generate a sessional **TEMP** table named _`PLAN_TABLE`_ with the plan saved.

2) And then format it with:

```sql
select * from table(dbms_xplan.display)
```  


3)
```sql
select plan_table_output
from table(
    dbms_xplan.display(
        NAME => 'PLAN_TABLE' if null,
        STATEMENT_ID => null,
        FORMAT => 'TYPICAL' if null
    )
);
``` 
* FORMAT: null to _`TYPICAL`_, others _`BASIC`_, _`ALL`_


---
## Actual Execution Plan
`dbms_xplan.display_cursor` will find the last executed _cursor_ from _`v$`_ structure populated with last SQL.
> `display_cursor` usage
```sql
-- 1. setup the hind for current session
alter session set statistics_level='ALL';

-- 2. SQL
select /*+ gather_plan_statistics */ cols...

-- 3. execution plan
select *
from table(
    dbms_xplan.display_cursor(
        SQL_ID => last one if null, CHILD_NUMBER => default 0,
        FORMAT => 'ALLSTATS LAST ALL +OUTLINE')
);
``` 
* SQL_ID: null to last SQL
* CHILD_NUMBER: null to 0
* FORMAT: null to _`TYPICAL`_, others _`BASIC`_, _`ALL`_ with _`ALLSTATS`_, _`LAST`_

> gather extended `A-Rows` and `A-Time`  
1. use per statement **HINT** `/*+ gather_plan_statistics */`
1. or setup **SESSION** hint `alter session set statistics_level='ALL';`

> finding the SQL_ID
```sql
select s.sql_id, s.child_number
from v$sql s
where s.sql_text like 'select xxx';

-- or in one go
select plan_table_output
from v$sql s,
table(
    dbms_xplan.display_cursor(
        s.sql_id, s.child_number, 'ALLSTATS LAST'
    )
)t
where s.sql_text like 'select /*+ gather_plan_statistics */ cols...';
```

[intro_plan]: https://blogs.oracle.com/oraclemagazine/post/how-to-read-an-execution-plan
[display_plan]: https://blogs.oracle.com/optimizer/post/how-do-i-display-and-read-the-execution-plans-for-a-sql-statement
[generate_plan]: https://blogs.oracle.com/optimizer/post/how-to-generate-a-useful-sql-execution-plan
[diff]: https://docs.oracle.com/cd/A58617_01/server.804/a58242/ch3.htm 
[trouble-shoot]: https://stackoverflow.com/questions/32823223/gather-plan-statistics-does-does-not-generate-basic-plan-statistics