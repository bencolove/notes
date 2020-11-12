# Basic of Kotlin coroutine
* koro builder
* `CoroutineScope`
* `CoroutineContext`
* context (dispathers)

---

##  _`coroutine builder`_

```kotlin
fun  <T> runBlocking(
    context: CoroutineContext = EmptyCoroutineContext,
    block: suspend CoroutineScope.()->T
): T

fun CoroutineScope.launch(
    context: CoroutineContext = EmptyCoroutineContext,
    start: CoroutineStart.DEFAULT,
    block: suspend CoroutineScope.()->Unit
): Job

fun CoroutineScope.async(
    context: CoroutineContext = EmptyCoroutineContext,
    start: CoroutineStart.DEFAULT,
    block: suspend CoroutineScope.()->T
): Deferred<T>

suspend fun <T> withContext(
    context: CoroutineContext,
    block: suspend CoroutineScope.() -> T
): T

suspend fun <T> withTimeout(
    timeMillis: Long,
    block: suspend CoroutineScope.() -> T
): T

suspend fun <T> withTimeoutOrNull(
    timeMillis: Long,
    block: suspend CoroutineScope.() -> T
): T?
```

Koro Builder | Will `suspend` | Is `CoroutineScope` extension | Start Rightaway | Return
---|:---:|:---:|---|---
`runBlocking` | N | N | Y | T
`launch` | N | Y | `CoroutineStart` |`Job`
`async` | N | Y | `CoroutineStart` | `Deferred<T>`
`withContext` | Y | N | Y |T
`withTimeout` | Y | N | Y | T
`withTimeoutOrNull` | Y | N | Y | T?


>`CoroutineStart.LAZY` + `Job.start()` or `suspend Job.join()` or `suspend Deferred.await()`


---

## `CorotineScope` 
[properties and methods][scope]


Properties:  
1. _coroutineContext_: `CoroutineContext`
1. _isActive_: `Boolean`

Extension functions:
1. `launch`
1. `async`
1. `cancel`


>_`scope builder`_

Scope Builder | Diff | ???
---|---|---
`GlobalScope` | predefined global one | `GlobalScope.async{...}`
`CoroutineScope()` | general purpose | `CoroutineScope.async{...}`
`MainScope()` | for UI with Dispatchers.Default |   
scoping function|
coroutineScope |
withContext |

>Structured Concurrency

Try to group related `koro`s togather into one `coroutineScope.async<Type>` for:  
1. when parent koro is cancelled, all children koros are cancelled
1. any exception thrown in the `async` block, other sleeping koros will receive `JobCancellationException` to terminate the block as whole
2. will it wait until all child koro ends ???

>Child koro

When a child koro is launched from the `CoroutineScope` of parent koro, then:
* child's _context_ inherit parent's
* child's `Job` becomes a child of parent's `Job`

>the `GlobalScope`
It has no parent in the first place, therefore any children koros of `GlobalScope` will not be tied to any _context_ and execute **indepently**.

## `CoroutineContext`

>`CoroutineContext` is passed down in hierachy as a koro builder inherits (called without context paramter) its parent's context. By default, the `GlobalScope` use `Dispatchers.Default`.

Context elements:
* name via `CoroutineName(String)`
* `ThreadLocal` via `threadLocal.asContextElement(value = VALUE)`
* [Job][context.Job]
* [CoroutineDispacher][context.CoroutineDispatchers]

>Combine *context* elements by:  
`launch(Dispatchers.Default + CoroutineName("test"))`

### context.`CoroutineDispatchers`

Dispatcher Type | Meaning | Usecases
---|---|---
`Dispatchers.Default` | pool of background threads | CPU-bound jobs
`Dispatchers.IO` | pool of on-demand create threads | I/O-bound
`Dispatchers.Unconfined` | start in current thread and in another used by the suspended function
private | create | `newSingleThreadContext` <br> `newFixedThreadPoolContext`
`Executor` | convert | `asCoroutineDispatcher`

<br>
<br>


### context.`Job`
job := `coroutineContext[Job]`  
isActive := `coroutineContext[Job]?.isActive==true`


### Adapt `ThreadLocal` into _context_
It is not THREAD local anymore, it is now KORO local

>Create

`ThreadLocal.withInitial { expr }`  
`ThreadLocal<Type>()`

>Modify

`threadLocal.get()`  
`threadLocal.set(VALUE)` is not working as it used to, instead use `threadLocal.asContextElement(value=VALUE)` to pass down through to child koro


[context.Job]: https://kotlin.github.io/kotlinx.coroutines/kotlinx-coroutines-core/kotlinx.coroutines/-job/index.html

[context.CoroutineDispatchers]: https://kotlin.github.io/kotlinx.coroutines/kotlinx-coroutines-core/kotlinx.coroutines/-coroutine-dispatcher/index.html

[scope]: https://kotlin.github.io/kotlinx.coroutines/kotlinx-coroutines-core/kotlinx.coroutines/-coroutine-scope/index.html

[ThreadLocalContextElement]: https://kotlin.github.io/kotlinx.coroutines/kotlinx-coroutines-core/kotlinx.coroutines/-thread-context-element/index.html
[ThreadLocal.asContextElement]: https://kotlin.github.io/kotlinx.coroutines/kotlinx-coroutines-core/kotlinx.coroutines/java.lang.-thread-local/as-context-element.html