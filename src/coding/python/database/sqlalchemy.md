# sqlalchemy
1. installation
1. engine
1. connection
1. transaction and isolation
1. autocommit
1. streaming result (DBAPI dependent)

Workflow:
1. create engine(dialect + pool)
1. engine.connect() to make real connection
1. sql

## installation


## create engine
```python
from sqlalchemy import create_engine

engine = create_engine(
    'postgresql://scott:tiger@localhost:5432/mydatabase'
)
```
* echo: bool
* echo_pool: str(LEVEL)
```python
import logging
logging.basicConfig()
logging.getLogger('sqlalchemy.engine').setLevel(logging.INFO)
logging.getLogger('sqlalchemy.pool').setLevel(logging.DEBUG)
```

## database urls
`dialect+driver://username:password@host:port/database`

Dialect | driver
---|---  
mssql | `pyodbc` <br> `pymssql`
mysql | `mysqldb` <br> `mysqldb`
postgresql | `psycopg2` <br> `pg8000`
oracle | `cx_oracle`
sqlite | r`sqlite:///path/to/file.db` <br> `sqlite://` (:memory)


>url_encode username and password  
```python
import urllib.parse
urllib.parse.quote_plus(password)
```

## connection
```python
from sqlalchemy import text

with engine.connect() as conn:
    result = conn.execute(text(
        "select username from users"
    ))
    for row in result:
        row['username']
```
The _result_:`CursorResult` will be closed:
1. all rows are exhausted
1. immediately upon construction in the case of returnning **NO** rows (like `update`).

## Transactions
```python
with engine.connect() as conn:
    with conn.begin():
        t1 = conn.execute(user_table.select())
        conn.execute(user_table.insert(), {'col1': 7, 'col2': 'name'})

# or concisely
with engine.begin() as conn:
    pass
```
The `Connection.begin()` returns `Transaction`:
1. `commit` when completes without exception
1. `rollback` when exception arisen

### Patterns
> transaction only  
```python
import contextlib

@contextlib.contextmanager
def transactioned(connection):
    if not connection.in_transaction():
        with connection.begin():
            yield connection
    else:
        yield connection

# usage
def do_a(connection):
    with transactioned(connection):
        do_b(connection)

def do_b(connection):
    with transactioned(connection):
        conneciton.execute(text('select 1'))
        connection.execute(table.insert(), {values})

with engine.connect() as conn:
    do_a(conn)
```

> with connection  
```python
import contextlib

def connect(engine):
    connection = None

    @contextlib.contextmanager
    def ensure_connect():
        nonlocal connection
        if connection is None:
            connection = engine.connect()
            with connection:
                with connection.begin():
                    yield connection
        else:
            yield connection
    
    return ensure_connect

# usage
def do_a(connectivity):
    with connectivity():
        do_b(connectivity)

def do_b(connectivity):
    with connectivity() as conn:
        conn.execute(text('select 1'))
        conn.execute(table.insert(), {values})

do_a(connect(engine))
```

## Isolation Level
Connection-wide:
`Connection.execution_options(isolation_level='<value>')`  
Preferred engine-wide: `create_engine(url, execution_options={"isolation_level": "REPEATABLE READ"})`  
Sub engine: `engine.execution_options(isolation_level="AUTOCOMMIT")`  
Levels:
1. "AUTOCOMMIT" (details followed)
1. "READ COMMITTED"
1. "READ UNCOMMITTED"
1. "REPEATABLE READ"
1. "SERIALIZABLE"

## Autocommit 
Implementation levels:
1. Database level (vender implementation)
1. DBAPI level (PEP-0249, always _transactional_ no `begin()`)
1. alchemy level

### Sqlalchemy autocommit
When in sqlalchemy autocommit mode:
1. it will inform(make calls to) underlying DBAPI to switch to autucommit mode per Connection
1. no `begin()` before each statement
1. issue `commit` and `rollback` after each statement


### DBAPI `autocommit` and Isolation Levels

Database | DBAPI | sqlalchemy
SQLSERVER `SET IMPLICIT_TRANSACTION`

Database | Default | Configurable
---|---|---
SQLSERVER | autocommit: <br> BEGIN and COMMIT each statement | `SET IMPLICIT_TRANSACTION` _ON_: 

## Stream result
1. Sqlalchemy always buffers all first and then fetched
1. For server-side cursor(unbuffered client-side) use DBAPI [[objects][server-side-cursor]]
```python
# one partition at a time
result = conn.execution_options(stream_results=True).execute(text('select * from table'))

for partition in result.partition(100):
    process_rows(patition)

# one row at a time with fixed buffer size

# fixed buffer
conn.execution_options(stream_results=True, max_row_buffer=100).execute(text('select * from table'))
for row in result:
    proecess_row(row)

# buffer empty then pull again
for row in result.yield_per(100):
    process_row(row)
```


[server-side-cursor]: https://docs.sqlalchemy.org/en/14/core/connections.html#dbapi-connections