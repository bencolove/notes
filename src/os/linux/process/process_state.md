# Process State
* RUNNING (EXECUTING, WAIT FOR EXECUTION)
* INTERRUPT (SLEEP, BLOCKED)
* UNINTERRUPT (until `INT` happens)
* HALT (terminated, `fd` existed before `wait4()`)
* STOP (after `SIGSTOP`,`SIGTSTP`,`SIGTTOU`)