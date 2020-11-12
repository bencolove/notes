# The **_channel_**
[official][channel]

The **_channel_** looks very much similar to the golang's version.

Channel | golang | kotlin
---|---|---| 
* bidirect channel <br> * receive channel <br> * send channel | make(chan <-) | `Channel<T>` <br> `RecieveChannel<T>`, `produce {...}` <br> `SendChannel<T>`
send | ch <- value | chan.send(value)
recv | value := <- ch | chan.receive()
close | close(ch) | chan.close()
loop | `for v := range ch {...}` | `for (v in chan) {...}` <br> `chan.consumeEach {...}`

>`ReceiveChannel<T>.consumeEach{...}` **WILL** `close` its associated _channel_ when complete successfully or fail.

## **_ticker_**

Ticker | golang | kotlin
---|---|---
periodic notification | `time.NewTicker(time.Duration(n) * time.Millisecond)` | `val tickerChannel = ticker(delayMillis=n, initialDelayMIllis=n,mode=TickerMode.FIXED_DELAY)`
stop ticker | `ticker.Stop()` | `tickerChannel.cancel()`
receive tick | `select { case <- ticker.C }` | `tickerChannel.receive()`


[channel]: https://kotlinlang.org/docs/reference/coroutines/channels.html