# Channel

_--runtime.hchan--_
```go
type hchan struct {
    qcount   uint
	dataqsiz uint
	buf      unsafe.Pointer
	elemsize uint16 // element data size
	closed   uint32
	elemtype *_type // element data type
	sendx    uint  
	recvx    uint
	recvq    waitq // blocked goro list
	sendq    waitq // blocked goro list

	lock mutex
}

type waitq struct {
    first *runtime.sudog
    last  *runtime.sudog
}
```
A buffered `chan` has an underlying ring structure, which is implemented by array buffer:
Field | Meaning
---|---
`qcount` | element count
`dataqsiz` | ring buffer length
`buf` | pointer
`sendx` | position for send 
`recvx` | position for recv

## Sending to _channel_
* when existing blocked goro on _recv_, set _data_ and make blocked goro next to run
* if no blocked goro and capacity of buffer, write _data_ to ring buffer
* otherwise block itself and yield gosched
* _`panic`_ when sending to closed _channel_

## Recieving from _channel_
`i <- ch` or `i, ok <- ch`  
and finally called to `runtime.chanrecv`  

* blocked when reading from empty _channel_
* reading from closed _channel_, `(_, false)` is returned
*
*

## Close __chanel__
* `__panic__` when closing closed _channel_
* blocked recv goros will be waken and `(zero-value, false)` is returned
* blocked send goros will `__panic__`

## Patterns
Restricts on _channel_:
1. It is advised that only channel senders to close the channel
1. Sending to closed _channel_ will _`panic`_
1. Reading from closed _channel_ will succeed immediately with _ok_=false

There is no way to tell whether a _channel_ is closed in a case when multiple senders trying to send data to a same _channel_ and they will _panic_.

>Better way  
Try to distinguish two use cases of _channel_: `control` and `data`.
1. Use `control` _channel_ singal reading goros to stop by closing it with `context.WithCancel(context).closeFunc`.
1. Use `data` _channel_ to pass data and use _ok_ to check whether the channel is closed

A _channel_ without buffer cost only around 12 words and _struct{}_ will not cost space. So it is best to use `chan struct{}` as control channel.
