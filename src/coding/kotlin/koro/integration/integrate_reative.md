# Adapt `reactive stream` to `koro`
>Target domain: `reative stream`

>Implementation:
* Project reactor (Spring)
* RxJava

>The stack which was `RxJava + Retrofit2 + OkHttp3` may be replaced with `koro + Retrofit2 + OkHttp3` for WebAPIs

>The stack for UI reative streams?

General pattern for adapt an `reactive stream` interface to kotlin's `koro`:

```kotlin
suspend fun execute(): ResultType {
    return suspendCoroutine<ResultType> { cont ->
        ...execute the asynchronous library ending with callbacks... 

        ...call onSuccess(continuation, result) passing the cont in case of successful execution...

        ...call onError(continuation, throwable) passing the cont in case of an error during the execution...

    }
}

private fun onSuccess(cont: Continuation, result: ResultType) {
    cont.resume(result)
}

private fun onError(cont: Continuation, throwable: Throwable) {
    cont.resumeWithException(throwable)
}

```

## For `cancallable` support for `koro`
`cancellable` means when the koro hierarchy is cancelled, the underlying adapted source has a chance to clean up whatever necessary.

For example, when a `koro` applied on Java `channel` is cancelled, it has to be a way to tell the underly ing channel to close itself.

So the pattern changes as:

_Cancellable for callbacks_  
```kotlin
suspend fun execute(): ResultType {
    return suspendCancellableCoroutine<ResultType> { cont ->
        ...execute the asynchronous library ending with callbacks... {

        ...call onSuccess(continuation, result) passing the cont in case of successful execution...

        ...call onError(continuation, throwable) passing the cont in case of an error during the execution...
        }

        // install callback called when CANCEL happens 
        cont.invokeOnCancellation { cause ->
            ...clean up whatever necessary...
        }

    }
}

private fun onSuccess(cont: Continuation, result: ResultType) {
    cont.resume(result)
}

private fun onError(cont: Continuation, throwable: Throwable) {
    cont.resumeWithException(throwable)
}
```
