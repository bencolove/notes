# Handling Exceptions in koro


>Create a koro exception handler
```kotlin
val handler = CoroutineExceptionHandler {
    CoroutineContext, Throwable ->
}
```

>Use it in a child _context_

`CoroutineScope.koro_builder(handler) {...}`
