## The `flow`

`flow`s are **COLD**, when created, it returns quickly, no need to be marked as `suspend`. It will only be triggered until `collect` and can be repeated

>Where (`CoroutineContext`) it runs ???

By design, the `flow` only executes when being collected. And the thread it runs from is the the thread where the collect function is called.


>Machenism

`flow { ... emit(value) ...}`  
against `sequence`  
`sequnce { ... yield(value) ... }`

>`flow` builder

`flow{...}`  
`flowOf(...)`  
`collection|sequence .asFlow()`

>intermedia operation

The intermedia operations are COLD as expected and can call suspending functions in contrasting to the `sequence` 

1. `map`  
2. `filter`
3. `transform`  
```kotlin
flow.transform { value ->
    emit()...
    emit()...
}
```
4. `take`



>terminal

1. `toList`
2. `toSet`
3. `fist`
4. `reduce`|`fold`  
5. `collect { value -> ... }`

>cancel/stop `flow`

`throw` Exception
    
>flow _context_

By default, `flow` runs in the same thread as where it is collected.

In order to change the _context_ it should be collected(emitting value), `.flowOn(CoroutineDispatchers)` is used as:
```kotlin
val f = flow {
    .... emits ...
}.flowOn(Dispatchers.Defult) // emit in background threads

fun main() = runBlocking<Unit> {
    // collect from UI thread(Main)
    f.collect { value -> ....}
}

```

>`flow` operations are by default sequential

Koros defined in a flow's downstream operators are by default sequential. Use `bufer()` before `collect` to enable flow to run in a async manner.
```kotlin
simple().buffer().collect {
    ...value...
}
```

>skip intermedia emitted results

`.conflate()` is used to keep slow downstream collector into pace with faster upstreams.

It looks like a `buffer()` intermedia operation that only buffers the more recent **one** result and discard any in between.


Operation | Meaning
---|---
Composing mutliple flows |
`zip` | wait for both results to emit the pair
`combine` | whenever either one updates, emit the pair
Flatten flows |
`flattenConcat` <br> `flatMapConcat` | emit inner flow results sequentially
`flattenMerge` <br> `flatMapMerge` | emit inner flow results concurrently
`flatMapLatest` | restart collector when a new result is emit from inner flows

---

## _flow_ Execution

About where a flow is launched(which thread essentially), two are involved:
1. `.flowOn(Dispatchers)` means
    emission and all intermedia operations happen in the `Dispatchers` while collection happens in the caller thread.
1. `.lauchIn(CoroutineScope)`
    the whole flow(emission, intermedia and terminal opertions) will happen in the scope

## flow Exception

Any exception thrown in terminal phrase or intermedia phrase can be caught from the outmost(caller of the flow) collect operation and all afterwards koro will be cancelled.

Two ways to handel excptions:
1. outmost `try...catch` arround `collect`
1. `.catch{ exception }` to transform exceptions

>`.catch{ exception }` only catches upstream exceptions and no way for downstream

In order to use `.catch` to handle `collect` downstream operator as well, `onEach {}.catch{}.collect()` recepit can be used.

## flow Completion in Action

>enable flow cancellation

Any flows that is cancellable must be somehow when emitting values it checks cancellation status.

while `emit(value)` has something behind to do the check, flow from collections like `IntRange.asFlow()`,`(1..5).asFlow()` has no this ability. Thus their downstream operation should include something like `.onEach { currentCoroutineContext().ensureActive() }` to throw a cancellation exception to trigger the process. `.cancellable()` is essentially same as above.

The recipt comes in `Collection|Sequence.asFlow().cancellable()`

## Thoughts of _flow_ Completion
When upstream emits completes, a `StopException` may be thrown from the source and passed downstream all the way to:
1. `.onCompletion { cause }`
    it only sees exceptions, when it'd be `StopException`, a successful completion is implied and cause will be null;
    when other exceptions including cancellation and failure, cause will be set accordingly, a failure completion is then indicated
1. `.catch { exception }`
    it only sees exceptions, excluding `StopException`
1. `.collect` terminal operations will know successful completion is trigger and exit normally. Otherwise it rethrow it to outter scope.


