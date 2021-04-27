# Compare *Coroutines*
1. *`Golang`*
1. *`Kotlin`*
1. *`Python`*

>Extra  

*`javascript` ES6* `async/await` syntantic sugar behind the scene: Promise and callback?


## Golang vs Kotlin
1. threading models
1. schedulers
1. scope/cancellation

Langurage | Implementation | Threadings | Schdulers | Scope/Cancellation | Sync Style
---|---|---|---|---|---
Golang | PC + SP | G-M-P | block + preempt | no, cooperative | no
Kotlin | Statemachine + CPS | ThreadPool | schedule onto threads | yes, exception | yes

> Threadings  

Golang allows goros to call systemcall then trap into kernel and blocked. But before the `G(goro)` does that, the runtime will de-copule the current `M(thread)` (soon be blocked) and find free `M(thread)` to execute rest goros from current `P(processor)` and let the `G(goro)` blocks `M(thread)`.

Kotlin does not have the privileges to play around `PC` and `SP` to reserve execution context. It can only `syntantic sugar` the pragdim that from current `thread` schedule `suspend fun` onto another (if not `join/await`ed) or `suspend` (by `Object.wait`) to block it to await for the task to be done like old-school manner.

> Sync Style for Async

Kolin codes sync style for async flows while go still has to deal with syncing.