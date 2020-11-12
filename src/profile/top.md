# `top`

By default `top` run on screen with interactive mode. Commands like `f` to config what `field/column` to be displayed.

To some extend, `top` can run in `b`(batch) mode to periodically sending off outputs:
* -p PID
* -c full COMMAND path
* -d _N_ interval in seconds
* -n _N_ iteration/loops
* -b batch mode

Combied with `grep '^ '` or `tail -1` to grab the wanted process info and with `awk { printf("col5: %10s, col6: %-10s", $5, $6); }` to prettify output of **DEFAULT** columns. (YES, columns are not customizable here)

## Customize `top` Output
It is trivial as

1. run `top` interactively and select which column and sort column you want by `f` key
2. `W` key to save the config file somewhere to `$HOME/.toprc` and next time when `top` is executed, this file get read. Simply restore to default settings by removing this file.
or `$HOME/config/procps/toprc`
