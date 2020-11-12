# With `Webflux`

They are implemented as a bunch of extension methods:

>From `koro` to `webflux`/`reactor`/`reactive` [link][koro-reactor]

Conversion Method | Argument
---|---
`Flow.asFlux` | `CoroutineContext`
`Job.asMono` | `CoroutineContext`
`Deferred.asMono` | `CoroutineContext`
`mono` | context, `CoroutineScope.()->T?`
`flux` | context, `ProducerScope<T>.()->Unit`
`publish` | context, `ProducerScope<T>.()->Unit`


>From `webflux`/`reactor`/`reactive` to `koro` [link][reactor-koro]

Conversion | Argument
---|---
`Publisher.asFlow` |
`Publisher.awaitFirst` |
`Publisher.awaitFirstOrDefault` |
`Publisher.awaitFirstOrElse` |
`Publisher.awaitFirstOrNull` |
`Publisher.awaitLast` |
`Publisher.awaitSingle` |


[koro-reactor]: https://github.com/Kotlin/kotlinx.coroutines/blob/master/reactive/kotlinx-coroutines-reactor/README.md
[reactive-koro]: https://github.com/Kotlin/kotlinx.coroutines/tree/master/reactive/kotlinx-coroutines-reactive