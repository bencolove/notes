# Koroutine
Scope is cancellable boundaries and context associated with it will be passed down with overriding Job()

First koro:
1. GlobalScope.launch: Job
1. runBlocking: T

## Structured Scope (life boundary)

Structured Concurence means a computed boundaries for a set of coroutines within to get cancelled altogather when the boundary(scope) is cancelled.

Those parent children koros are chained by Job which inherits with CoroutineContext when launched from within a scope(bundary).

Method | suspend | return | others
---|---|---|---
runBlocking | X |
coroutineScope | O |
withContext | O | T | coroutineScope + context

<br>

## Scope
Scope | Type | Java | Android
---|---|:---:|:---:
GlobalScope | managed | O | ?
viewModelScope | managed | X | O
lifecycleScope | managed | X | O
CorotineScope | create new |  | ?
ActorScope | managed | O | X

## launcher

Launcher | suspend | return | context | start |
---|:---:|---|---|---
scope.async | X | Deferred< T > | O | O
scope.launch | X | Job | O | O

<br>

CoroutineStart | Purpose
---|---
DEFAULT | *`immediately`* schedule based on context
LAZY | schedule when `join()` or `await()`
ATOMIC | DEFAULT + non-cancellable
UNDISPATCHED | *`no suspend`*, current thread till first suspend within

<br>
---

## Cancellation
1. cancel CoroutineScope
1. cancel Coroutine
1. how cancellation is captured within a Coroutine

```java
val job1 = outerScope.launch { ... }
val job2 = outerScope.launch { ... }

// job2 is not affected
job1.cancel()

// all child coroutine (launched within) cancelled
outerScope.cancel(cause: CancellabtionException? = null)
```

>Cooperative Cancellation

`CoroutineScope.cancel()` comes with no magic to automatically cancel the currently running koro. Rather it is the koro that responsible for stopping itself to quit ASAP!!!

Ways to let cancellation honered:
1. `Coroutine.isActive`: Boolean
1. `Coroutine.ensureActive()` which throws CancellationException if **NOT** ACTIVE
1. `yield()` checks ACTIVE like ensureActive() **`before`** suspending current koro

>Cancel Job.join() and Deferred.await()
1. Job
```java
val job = scope.launch { ... }

job.cancel()
// wait until job is completed (honor cancallation due to program)
job.join()

job.join()
// nothing, the job is done
job.cancel()
```

2. Deferred
```java
val def = scope.async { ... }

def.cancel()
// suspend till job completed
def.await()

def.await()
// JobCancellationException
def.cancel()
```

> Cleanup when Cancelled
1. isActive
1. try/catch/finally

```java
while ( condition && isActive ) { ... }

// koro work is completed so clean it up
println("clean up")
```

```java
try {
    longRunningWorkMayBeCancelled()
} catch(e: CancellationException ) {
    println("cancelled")
} finally {
    // normally no suspending code anymore form here
    noSuspendCleanup()
    // unless
    withContext(NonCancellable) {
        delay(1000L)
        suspendableCleanup()
    }
}
```


<br>
---

## Context
1. Job
1. CoroutineDispatcher
1. CoroutineName
1. CoroutineExceptionHandler
1. NonCancellable

## Runtime Scheduling
**`NOT`** like `Go`, koroutines are not preemptive, which means in some cases, a koro has to `yield()` to give up current thread for other koro to kick in to resume.

In such cases, long-runing tasks should be dedicatedly put into a thread pool wrapped as an async task which will resume in koro runtime after complete from the pool.

When a koro will give up its computation ???
1. call a `suspend` function (within suspend function) and wait for its completion
1. actively call `yield()` (because `yeild` is a `suspend` function)
1. `suspend Deferred.await()` will cause `yield` **UNTIL** it returns. 


## Dispatchers
Dispatcher | JAVA | Android
---|---|---
Default | O | CPU-Bound
IO | O | IO-Bound
Unconfined | O (current thread till first suspension) | X
newSingleThreadContext <br> newFixedThreadPoolContext | O | ?
Executor | O | ?
Main | X | UI thread

## Top-level Suspend Functions
1. delay
1. yield
1. withContext
1. withTimeout (throw TimeoutException)
1. withTimeoutOrNull (no exception but null return)
1. awaitAll
1. joinAll

[cancallation]: https://medium.com/androiddevelopers/cancellation-in-coroutines-aa6b90163629