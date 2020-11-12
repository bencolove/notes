# Elements of `CoroutineContext`

Use with `koro_builder(ctx1 + ctx2 ...)`

_context_ is another means of grouping koros togather as parent-child relationship.

Whenever a child koro is triggered from 

---


## Common _context_ Element Keys

* `CoroutineName(name)`
* `Job`
* `CoroutineExceptionHandler {...}`
* Key<CoroutineDispatchers>

