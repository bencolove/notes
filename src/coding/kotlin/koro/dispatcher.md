# `CoroutineDispatchers`

Dispatcher | Used | Meaning | Attrs
---|---|---|---
private | `runBlocking` | dedicated event-loop thread | 
`Default` | `launch`<br>`async` | shared pool threads | max(2,#cores)

[default]: https://kotlin.github.io/kotlinx.coroutines/kotlinx-coroutines-core/kotlinx.coroutines/-dispatchers/-default.html