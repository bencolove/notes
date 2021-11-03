# `ps`
The `ps`(_process state_) at time being. Use `top` to keep tailing/watching.

> States
1. D uninterruptible like IO
1. R runnable(on run queue)
1. S sleeping
1. T traced or stopped
1. Z defaunct

## filter
* -p PID
* -u UID
* -ef | grep
* aux | grep

## select columns  
1. pid
1. ppid
1. uid
1. user
1. comm only file name
1. cmd file name and args



