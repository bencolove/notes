# Concurrency on `webflux`

Concurrency Platform | Stack | Threading
---|---|---
Java NIO.1 | `FileChannel`(`ByteChannel`) + `select`
Java NIO.2 | `AysnchrousFileChannel` +  _callback_| dedicated thread pool (5)
Netty | event loop | 
`webflux`(`reactor`+ `Netty`) | `DataBufferUtils.readByteChannel`


[tutorial]: https://www.baeldung.com/spring-webflux-concurrency