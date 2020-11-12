# `perf`
[tutorial][tutorial]
* events
* counting events with `perf stat`
* sampling with `perf record`
* live analysis with `perf top`

## Common Recipes

>`top -b -p $(pidof PROG)`

>`ps --format pid,pcpu,cputime,etime,size,vsz,cmd -p`

> `stat/record` a process by `PID` at rate of `1000 samples/sec` for `5s`

`perf record -p PID -F 1000 sleep 5`

## 1. Countable Events
`perf list`  

Events | Meaning | Common Events
---|---|---
Hardware event | Performance Monitoring Unit (**PMU**) on processor | `cpu-cycles` <br> `instructions` <br> `cache-misses` <br> `branch-misses`
Software event | kernel counters | `context-switches`(`cs`) <br> `cpu-clock`
Hardware cache event | `perf_events` predefined | `L1-dcache-loads` 
Tracepint event | tracing | `sched:sched_stat_runtime` <br> `syscalls:sys_enter_socket`


Target | Meaning | CMD use as Default value 
---|---|---
-a | all cpus, `cpu-wide` mode | `top`
-C | `per-cpu` mode |
-p | `per-process` mode | `stat`
-t | `per-thread` mode | `record`


---

## 2. Counting with `stat`
`perf stat -B dd if=/dev/zero of=/dev/null count=1000000`  

`stat` options | Meaning
---|---
-e | event symbolic names by `,` no space
-B | thousand separator `,`
event:_modifier_ | `u` priv level3,2,1 _user_ <br> `k` priv level 0 _kernel_ <br> `h` hypervisor events <br> `H` host machine <br> `G` guest machine
-a | all _cpus_
-C _n_ | selective cpu by `n`
-p PID | attach to process by PID
-i | no-inherit, not counting for child processes
-t TID | attach to thread by TID
-r _loop_ | repeat then stats as mean and standard deviation
-x _delimiter_ | set delimiter to output fields, **NOT** compatible with `-B`

### `stat` Mode
* `perf stat COMMAND`
* `perf stat -p PID -t TID [COMMAND]`

Either counting for a _COMMAND_ during its execution or **attaching** to a process/thread during following command's execution like `sleep 2` for 2 seconds

PID can be found by `ps ax | grep COMMAND`  
TID can be found by `ps -L ax | grep COMMAND | head -5`

---

## 3. Sampling `record`
* sampleing with `record`
* analysis with `report`
* souce level analysis with `annotate`

Sampling is based on **event occurence** not wall-clock timer tick. The **event** here is by default the `cycle` events. When the counter for `cycle` overflows, a sampling is triggered to collect inform about the target. 

### Custom Sampling Rate
* by frequence -F
* by num of events -c

Means | Option | Meaning
---|---|---
frequency | -F _persec_ | num of samples per second
period | -c _numOfEvts_ | one sample every number of events



>First step `record`  to _perf.data_

>Analysis with `report`

## 4. Live Analysis `top`
`perf top` by default monitors all cpus. Use `-C` to monitor specified cpu.

When topping, press `s` to disassambly (drill down) into specified function. 


[tutorial]: https://perf.wiki.kernel.org/index.php/Tutorial