# Integrate _koro_ with Java _nio_

In order to adapt Java _nio_ to be partner with _koro_, a way of converting _nio_ `Channel` events to _koro_'s `suspend` function is needed.



[how-to]: https://github.com/Kotlin/kotlinx.coroutines/blob/87eaba8a287285d4c47f84c91df7671fcb58271f/integration/kotlinx-coroutines-nio/src/Nio.kt