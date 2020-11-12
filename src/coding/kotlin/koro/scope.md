# `CoroutineScope`

The `CoroutineScope` groups koros into a so-called `structured concurrency` block.

When a koro builder spwans a child koro, it inherits parent's `CoroutineContext` and sets its `Job` as a key in the _context_.

>When `CancellationException` thrown from child koro

It does so by `.cancel()` on its `Job`. And it terminates itself without cancelling parent.

>When other exception thrown from child koro

It will propagate the exception to parent and parent:
1. recursively cancel all child koros
2. call `CoroutineExceptionHandler` from the _context_ if set.

>When parent is cancelled

All child koros will be cancelled in trun

>When other Exception is thrown from parent

Parent does the same

>The scope(parent koro) will wait for all children complete


## Cancellation within _context_

When child koro is cancelled(via `Job.cancel()`), it terminates not cancelling its parent.

A `CancellationException` is thrown at a koro's suspension point when it last yield it control.

`CoroutineExceptionHandler` will by design ignore it.

## Exception within _context_

On the other hand, when a koro encounters an exception other than `CancellationException`, it will cancel its parent (propogate the exception) thus its parent will in turn cancel all children and call `CoroutineExceptionHandler` if one is  specified.

