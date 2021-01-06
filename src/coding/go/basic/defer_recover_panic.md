# `defer`, `panic` and `recover`
The story:
1. `recover` only happens **AFTER** `panic`
>Put `recover` in `defer` is a wise decision
2. `defer` will be chained and reversely to **CURRENT** goro
3. `panic` only triggers `defer`s defined in **CURRENT** goroutine
>`panic` will not trigger `recover` from other goro
4. `panic` can be nested

## `panic`
During compile time, statement `panic` will be traspiled to `runtime.gopanic()`. What it does is:
1. create _struct_ `runtime._panic` and chain to head of panic list for **CURRENT** `goro._panic`
2. invoke **CURRENT** `goro._defer` list from head to tail by `runtime.reflectcall`
3. handle `recover` if defined
1. call `runtime.fatalpanic` to exit (`runtime.exit(2)`) unless `recover` is called

## `recover`
Compiler will transpile `recover` to `runtime.gorecover`

`recover` simply quit if **CURRENT** goro has no `panic`ed before or set `_panic.recovered = true` and let `gopanic` (triggered by `panic` as above) to deal with it.

## The whole story:
1. compiler transpile 
   * `panic` -> `runtime.gopanic`
   * `recover` -> `runtime.gorecover`
   * insert `deferreturn` right before `defer` ends
1. when `panic`(call `runtime.gopanic`), **CURRENT** goro's `_defer` list will be travelsed and executed from head (reversed to defined order)
1. in `defer`