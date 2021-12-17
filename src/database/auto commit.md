# Autocommit
1. SQLServer (default _`enabled`_)
1. Mysql (default _`disabled`_)

## SQLServer
Autocommit _default_ _`enabled`_
1. with `BEGIN TRANS`  
1. No `BEGIN TRANS`
1. SET IMPLICIT_TRANSACTION ON|OFF

>1. With `BEGIN TRANS`  
A program managed transaction is created and `COMMIT` or `ROLLBACK` has to be explicit issue

>2. No `BEGIN TRANS`  
When **SET IMPLICIT_TRANSACTION**:
1. `ON`: a `BEGIN TRANS` will be issued automatically before next statement, and therefore `COMMIT` or `ROLLBACK` has to be followed.
1. `OFF`: each statement is wrapped in a transaction. It _`begins`_ before and _`commit`_ afterwards automatically
1. default `ON`

## Mysql
Autocommit _default_ _`disabled`_
Each statement is wrapped in a transaction. It _`begins`_ before and _`commit`_ afterwards automatically.