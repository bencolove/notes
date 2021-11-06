# Ways to Run Executables
1. Invoke-Command
2. Start-Process
3. Start-Job

## Start-Process
Run the executable in a new process opening a new powershell window to in the same window(
with _`-NoNewWindow`_) and output to its residing window's stdout.

> Example  
`Start-Process powershell { $pid } -NoNewWindow`

## Start-Job
Run a _`powershell`_ block `{ script }` in a background process which can be managed by:
1. get-job `gjb`
1. wait-job `wjb`
1. receive-job `rcjb`
1. stop-job `stjb`
1. suspend-job `sujb`
1. resume-job `rujb`
1. remove-jb `rjb`

Read job's stdout,stderr by `recieve-job`


