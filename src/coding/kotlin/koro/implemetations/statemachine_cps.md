# How *kotlin* Implements *corotines*

*`kotlin`* implements *koroutines* by 
* **State machine**
* **CPS** (Continuation Passing Styles, **`callback`** in short)

For a single coroutine block (like `launch`, `async`, `runBlocking`), the compiler transforms (devides the entire block at *suspension* points, calling `suspendCoroutin` or *suspend* functions into a state machine with *`numOfSuspension`*` + 1` states.

When calling `Deferred.await()` (returned by `async` builder) with `Continuation` object, it passes along with a *callback* which upon success or failure will *resume* the state machine.

[]: https://www.bennyhuo.com/2019/04/01/basic-coroutines/
[]: https://medium.com/google-developer-experts/coroutines-suspending-state-machines-36b189f8aa60
[]: https://stackoverflow.com/questions/53526556/how-do-kotlin-coroutines-work-internally
[]: https://stackoverflow.com/questions/47871868/what-does-the-suspend-function-mean-in-a-kotlin-coroutine

